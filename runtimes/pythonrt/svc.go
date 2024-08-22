package pythonrt

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"go.autokitteh.dev/autokitteh/runtimes/pythonrt/pb"
	"go.autokitteh.dev/autokitteh/sdk/sdkservices"
	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type remoteSvc struct {
	pb.UnimplementedWorkerServer

	log       *zap.Logger
	cbs       *sdkservices.RunCallbacks
	runID     sdktypes.RunID
	syscallFn sdktypes.Value
	lis       net.Listener
	srv       *grpc.Server
	port      int
	runner    pb.RunnerClient

	// One of these will signal end of execution
	result chan []byte
	error  chan string
}

func newRemoteSvc(log *zap.Logger, cbs *sdkservices.RunCallbacks, runner pb.RunnerClient, runID sdktypes.RunID, syscallFn sdktypes.Value) *remoteSvc {
	svc := remoteSvc{
		log:    log,
		cbs:    cbs,
		runner: runner,
		runID:  runID,

		result: make(chan []byte, 1),
		error:  make(chan string, 1),
	}

	if syscallFn.IsValid() {
		svc.syscallFn = syscallFn
	}

	return &svc
}

func (s *remoteSvc) Health(context.Context, *pb.HealthRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{}, nil
}

func pyLevelToZap(level string) zapcore.Level {
	switch level {
	case "DEBUG":
		return zap.DebugLevel
	case "INFO":
		return zap.InfoLevel
	case "WARNING":
		return zap.WarnLevel
	case "ERROR":
		return zap.ErrorLevel
	}

	return zap.InfoLevel
}

func (s *remoteSvc) Log(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	if req.Level == "" {
		return nil, status.Error(codes.InvalidArgument, "empty level")
	}
	level := pyLevelToZap(req.Level)
	s.log.Log(level, req.Message, zap.String("source", "python"))
	return &pb.LogResponse{}, nil
}

func (s *remoteSvc) Print(ctx context.Context, req *pb.PrintRequest) (*pb.PrintResponse, error) {
	s.cbs.Print(ctx, s.runID, req.Message)
	return &pb.PrintResponse{}, nil
}

// ak functions

func (s *remoteSvc) Sleep(ctx context.Context, req *pb.SleepRequest) (*pb.SleepResponse, error) {
	if req.DurationMs < 0 {
		return nil, status.Error(codes.InvalidArgument, "negative time")
	}

	secs := float64(req.DurationMs) / 1000.0
	args := []sdktypes.Value{sdktypes.NewFloatValue(secs)}
	_, err := s.cbs.Call(ctx, s.runID, s.syscallFn, args, nil)
	var resp pb.SleepResponse
	if err != nil {
		resp.Error = err.Error()
		err = status.Errorf(codes.Internal, "sleep(%f) -> %s", secs, err)
	}

	return &resp, err
}

func (s *remoteSvc) Subscribe(ctx context.Context, req *pb.SubscribeRequest) (*pb.SubscribeResponse, error) {
	if req.Connection == "" || req.Filter == "" {
		return nil, status.Error(codes.InvalidArgument, "missing connection name or filter")
	}

	args := []sdktypes.Value{
		sdktypes.NewStringValue(req.Connection),
		sdktypes.NewStringValue(req.Filter),
	}
	out, err := s.cbs.Call(ctx, s.runID, s.syscallFn, args, nil)
	if err != nil {
		err = status.Errorf(codes.Internal, "subscribe(%s, %s) -> %s", req.Connection, req.Filter, err)
		return &pb.SubscribeResponse{Error: err.Error()}, err
	}

	signalID := out.GetString().Value()
	resp := pb.SubscribeResponse{SignalId: signalID}
	return &resp, nil
}

func (s *remoteSvc) NextEvent(ctx context.Context, req *pb.NextEventRequest) (*pb.NextEventResponse, error) {
	if len(req.SignalIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "at least one signal ID required")
	}
	if req.TimeoutMs < 0 {
		return nil, status.Error(codes.InvalidArgument, "timeout < 0")
	}

	args := make([]sdktypes.Value, len(req.SignalIds))
	for i, id := range req.SignalIds {
		args[i] = sdktypes.NewStringValue(id)
	}
	// timeout is kw only
	kw := make(map[string]sdktypes.Value)
	if req.TimeoutMs > 0 {
		kw["timeout"] = sdktypes.NewFloatValue(float64(req.TimeoutMs) / 1000.0)
	}

	val, err := s.cbs.Call(ctx, s.runID, s.syscallFn, args, kw)
	if err != nil {
		err = status.Errorf(codes.Internal, "next_event(%s, %s) -> %s", req.SignalIds, req.TimeoutMs, err)
		return &pb.NextEventResponse{Error: err.Error()}, err
	}

	out, err := val.Unwrap()
	if err != nil {
		err = status.Errorf(codes.Internal, "can't unwrap %v - %s", val, err)
		return &pb.NextEventResponse{Error: err.Error()}, err
	}

	data, err := json.Marshal(out)
	if err != nil {
		err = status.Errorf(codes.Internal, "can't json.Marshal %v - %s", out, err)
		return &pb.NextEventResponse{Error: err.Error()}, err
	}

	resp := pb.NextEventResponse{
		Event: &pb.Event{
			Data: data,
		},
	}
	return &resp, nil
}

func (s *remoteSvc) Unsubscribe(ctx context.Context, req *pb.UnsubscribeRequest) (*pb.UnsubscribeResponse, error) {
	args := []sdktypes.Value{
		sdktypes.NewStringValue(req.SignalId),
	}
	_, err := s.cbs.Call(ctx, s.runID, s.syscallFn, args, nil)
	if err != nil {
		err = status.Errorf(codes.Internal, "subscribe(%s) -> %s", req.SignalId, err)
		return &pb.UnsubscribeResponse{Error: err.Error()}, err
	}

	return &pb.UnsubscribeResponse{}, nil
}

func (s *remoteSvc) call(ctx context.Context, callID string, fn sdktypes.Value) {
	out, err := s.cbs.Call(ctx, s.runID, fn, nil, nil)
	req := pb.ActivityReplyRequest{
		CallId: callID,
	}

	switch {
	case err != nil:
		req.Error = fmt.Sprintf("call_id: %s - %s", callID, err)
		s.log.Error("activity reply error", zap.Error(err))
	case !out.IsBytes():
		req.Error = fmt.Sprintf("call output not bytes: %#v", out)
		s.log.Error("activity reply error", zap.String("error", req.Error))
	default:
		data := out.GetBytes().Value()
		req.Result = data
	}

	reply, err := s.runner.ActivityReply(ctx, &req)
	switch {
	case err != nil:
		s.log.Error("activity reply error", zap.Error(err))
	case reply.Error != "":
		s.log.Error("activity reply error", zap.String("error", reply.Error))
	}
}

// Runner starting activity
func (s *remoteSvc) Activity(ctx context.Context, req *pb.ActivityRequest) (*pb.ActivityResponse, error) {
	xid := sdktypes.NewExecutorID(s.runID)
	name := fmt.Sprintf("fn_%s", req.CallId)
	fn, err := sdktypes.NewFunctionValue(xid, name, nil, nil, pyModuleFunc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "new function value: %s", err)
	}

	go s.call(ctx, req.CallId, fn)

	return &pb.ActivityResponse{}, nil
}

func (s *remoteSvc) Done(ctx context.Context, req *pb.DoneRequest) (*pb.DoneResponse, error) {
	if req.Error != "" {
		s.error <- req.Error
	} else {
		s.result <- req.Result
	}

	return &pb.DoneResponse{}, nil
}

func freePort() (int, error) {
	conn, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	conn.Close()
	return conn.Addr().(*net.TCPAddr).Port, nil
}

// Start starts the server on a free port in a new goroutine.
// It returns the port the server listens on.
func (s *remoteSvc) Start() error {
	port, err := freePort()
	if err != nil {
		return err
	}
	s.port = port

	addr := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.lis = lis

	srv := grpc.NewServer(grpc.UnaryInterceptor(newInterceptor(s.log)))
	pb.RegisterWorkerServer(srv, s)
	reflection.Register(srv)

	s.log.Info("server starting", zap.String("address", addr))

	go func() {
		if err := srv.Serve(lis); err != nil {
			s.log.Error("serve gRPC", zap.Error(err))
		}
	}()

	return nil
}

func (s *remoteSvc) Stop() {
	s.srv.Stop()
	if err := s.lis.Close(); err != nil {
		s.log.Error("close listener", zap.Error(err))
	}
}

func newInterceptor(log *zap.Logger) func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	fn := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		log.Info("call", zap.String("method", info.FullMethod))

		return handler(ctx, req)
	}

	return fn
}
