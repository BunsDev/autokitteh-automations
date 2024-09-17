// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: autokitteh/events/v1/svc.proto

package eventsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "go.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/events/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// EventsServiceName is the fully-qualified name of the EventsService service.
	EventsServiceName = "autokitteh.events.v1.EventsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// EventsServiceSaveProcedure is the fully-qualified name of the EventsService's Save RPC.
	EventsServiceSaveProcedure = "/autokitteh.events.v1.EventsService/Save"
	// EventsServiceGetProcedure is the fully-qualified name of the EventsService's Get RPC.
	EventsServiceGetProcedure = "/autokitteh.events.v1.EventsService/Get"
	// EventsServiceListProcedure is the fully-qualified name of the EventsService's List RPC.
	EventsServiceListProcedure = "/autokitteh.events.v1.EventsService/List"
)

// EventsServiceClient is a client for the autokitteh.events.v1.EventsService service.
type EventsServiceClient interface {
	Save(context.Context, *connect.Request[v1.SaveRequest]) (*connect.Response[v1.SaveResponse], error)
	Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error)
	// List returns events without their data.
	List(context.Context, *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error)
}

// NewEventsServiceClient constructs a client for the autokitteh.events.v1.EventsService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEventsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) EventsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &eventsServiceClient{
		save: connect.NewClient[v1.SaveRequest, v1.SaveResponse](
			httpClient,
			baseURL+EventsServiceSaveProcedure,
			opts...,
		),
		get: connect.NewClient[v1.GetRequest, v1.GetResponse](
			httpClient,
			baseURL+EventsServiceGetProcedure,
			opts...,
		),
		list: connect.NewClient[v1.ListRequest, v1.ListResponse](
			httpClient,
			baseURL+EventsServiceListProcedure,
			opts...,
		),
	}
}

// eventsServiceClient implements EventsServiceClient.
type eventsServiceClient struct {
	save *connect.Client[v1.SaveRequest, v1.SaveResponse]
	get  *connect.Client[v1.GetRequest, v1.GetResponse]
	list *connect.Client[v1.ListRequest, v1.ListResponse]
}

// Save calls autokitteh.events.v1.EventsService.Save.
func (c *eventsServiceClient) Save(ctx context.Context, req *connect.Request[v1.SaveRequest]) (*connect.Response[v1.SaveResponse], error) {
	return c.save.CallUnary(ctx, req)
}

// Get calls autokitteh.events.v1.EventsService.Get.
func (c *eventsServiceClient) Get(ctx context.Context, req *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// List calls autokitteh.events.v1.EventsService.List.
func (c *eventsServiceClient) List(ctx context.Context, req *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// EventsServiceHandler is an implementation of the autokitteh.events.v1.EventsService service.
type EventsServiceHandler interface {
	Save(context.Context, *connect.Request[v1.SaveRequest]) (*connect.Response[v1.SaveResponse], error)
	Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error)
	// List returns events without their data.
	List(context.Context, *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error)
}

// NewEventsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEventsServiceHandler(svc EventsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	eventsServiceSaveHandler := connect.NewUnaryHandler(
		EventsServiceSaveProcedure,
		svc.Save,
		opts...,
	)
	eventsServiceGetHandler := connect.NewUnaryHandler(
		EventsServiceGetProcedure,
		svc.Get,
		opts...,
	)
	eventsServiceListHandler := connect.NewUnaryHandler(
		EventsServiceListProcedure,
		svc.List,
		opts...,
	)
	return "/autokitteh.events.v1.EventsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case EventsServiceSaveProcedure:
			eventsServiceSaveHandler.ServeHTTP(w, r)
		case EventsServiceGetProcedure:
			eventsServiceGetHandler.ServeHTTP(w, r)
		case EventsServiceListProcedure:
			eventsServiceListHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedEventsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEventsServiceHandler struct{}

func (UnimplementedEventsServiceHandler) Save(context.Context, *connect.Request[v1.SaveRequest]) (*connect.Response[v1.SaveResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autokitteh.events.v1.EventsService.Save is not implemented"))
}

func (UnimplementedEventsServiceHandler) Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autokitteh.events.v1.EventsService.Get is not implemented"))
}

func (UnimplementedEventsServiceHandler) List(context.Context, *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autokitteh.events.v1.EventsService.List is not implemented"))
}
