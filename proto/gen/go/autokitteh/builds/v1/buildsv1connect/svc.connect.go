// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: autokitteh/builds/v1/svc.proto

package buildsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "go.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/builds/v1"
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
	// BuildsServiceName is the fully-qualified name of the BuildsService service.
	BuildsServiceName = "autokitteh.builds.v1.BuildsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BuildsServiceGetProcedure is the fully-qualified name of the BuildsService's Get RPC.
	BuildsServiceGetProcedure = "/autokitteh.builds.v1.BuildsService/Get"
	// BuildsServiceListProcedure is the fully-qualified name of the BuildsService's List RPC.
	BuildsServiceListProcedure = "/autokitteh.builds.v1.BuildsService/List"
	// BuildsServiceSaveProcedure is the fully-qualified name of the BuildsService's Save RPC.
	BuildsServiceSaveProcedure = "/autokitteh.builds.v1.BuildsService/Save"
	// BuildsServiceDeleteProcedure is the fully-qualified name of the BuildsService's Delete RPC.
	BuildsServiceDeleteProcedure = "/autokitteh.builds.v1.BuildsService/Delete"
	// BuildsServiceDownloadProcedure is the fully-qualified name of the BuildsService's Download RPC.
	BuildsServiceDownloadProcedure = "/autokitteh.builds.v1.BuildsService/Download"
)

// BuildsServiceClient is a client for the autokitteh.builds.v1.BuildsService service.
type BuildsServiceClient interface {
	Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error)
	List(context.Context, *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error)
	Save(context.Context, *connect.Request[v1.SaveRequest]) (*connect.Response[v1.SaveResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
	Download(context.Context, *connect.Request[v1.DownloadRequest]) (*connect.Response[v1.DownloadResponse], error)
}

// NewBuildsServiceClient constructs a client for the autokitteh.builds.v1.BuildsService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBuildsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BuildsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &buildsServiceClient{
		get: connect.NewClient[v1.GetRequest, v1.GetResponse](
			httpClient,
			baseURL+BuildsServiceGetProcedure,
			opts...,
		),
		list: connect.NewClient[v1.ListRequest, v1.ListResponse](
			httpClient,
			baseURL+BuildsServiceListProcedure,
			opts...,
		),
		save: connect.NewClient[v1.SaveRequest, v1.SaveResponse](
			httpClient,
			baseURL+BuildsServiceSaveProcedure,
			opts...,
		),
		delete: connect.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+BuildsServiceDeleteProcedure,
			opts...,
		),
		download: connect.NewClient[v1.DownloadRequest, v1.DownloadResponse](
			httpClient,
			baseURL+BuildsServiceDownloadProcedure,
			opts...,
		),
	}
}

// buildsServiceClient implements BuildsServiceClient.
type buildsServiceClient struct {
	get      *connect.Client[v1.GetRequest, v1.GetResponse]
	list     *connect.Client[v1.ListRequest, v1.ListResponse]
	save     *connect.Client[v1.SaveRequest, v1.SaveResponse]
	delete   *connect.Client[v1.DeleteRequest, v1.DeleteResponse]
	download *connect.Client[v1.DownloadRequest, v1.DownloadResponse]
}

// Get calls autokitteh.builds.v1.BuildsService.Get.
func (c *buildsServiceClient) Get(ctx context.Context, req *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// List calls autokitteh.builds.v1.BuildsService.List.
func (c *buildsServiceClient) List(ctx context.Context, req *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Save calls autokitteh.builds.v1.BuildsService.Save.
func (c *buildsServiceClient) Save(ctx context.Context, req *connect.Request[v1.SaveRequest]) (*connect.Response[v1.SaveResponse], error) {
	return c.save.CallUnary(ctx, req)
}

// Delete calls autokitteh.builds.v1.BuildsService.Delete.
func (c *buildsServiceClient) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// Download calls autokitteh.builds.v1.BuildsService.Download.
func (c *buildsServiceClient) Download(ctx context.Context, req *connect.Request[v1.DownloadRequest]) (*connect.Response[v1.DownloadResponse], error) {
	return c.download.CallUnary(ctx, req)
}

// BuildsServiceHandler is an implementation of the autokitteh.builds.v1.BuildsService service.
type BuildsServiceHandler interface {
	Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error)
	List(context.Context, *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error)
	Save(context.Context, *connect.Request[v1.SaveRequest]) (*connect.Response[v1.SaveResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
	Download(context.Context, *connect.Request[v1.DownloadRequest]) (*connect.Response[v1.DownloadResponse], error)
}

// NewBuildsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBuildsServiceHandler(svc BuildsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	buildsServiceGetHandler := connect.NewUnaryHandler(
		BuildsServiceGetProcedure,
		svc.Get,
		opts...,
	)
	buildsServiceListHandler := connect.NewUnaryHandler(
		BuildsServiceListProcedure,
		svc.List,
		opts...,
	)
	buildsServiceSaveHandler := connect.NewUnaryHandler(
		BuildsServiceSaveProcedure,
		svc.Save,
		opts...,
	)
	buildsServiceDeleteHandler := connect.NewUnaryHandler(
		BuildsServiceDeleteProcedure,
		svc.Delete,
		opts...,
	)
	buildsServiceDownloadHandler := connect.NewUnaryHandler(
		BuildsServiceDownloadProcedure,
		svc.Download,
		opts...,
	)
	return "/autokitteh.builds.v1.BuildsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BuildsServiceGetProcedure:
			buildsServiceGetHandler.ServeHTTP(w, r)
		case BuildsServiceListProcedure:
			buildsServiceListHandler.ServeHTTP(w, r)
		case BuildsServiceSaveProcedure:
			buildsServiceSaveHandler.ServeHTTP(w, r)
		case BuildsServiceDeleteProcedure:
			buildsServiceDeleteHandler.ServeHTTP(w, r)
		case BuildsServiceDownloadProcedure:
			buildsServiceDownloadHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBuildsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBuildsServiceHandler struct{}

func (UnimplementedBuildsServiceHandler) Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autokitteh.builds.v1.BuildsService.Get is not implemented"))
}

func (UnimplementedBuildsServiceHandler) List(context.Context, *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autokitteh.builds.v1.BuildsService.List is not implemented"))
}

func (UnimplementedBuildsServiceHandler) Save(context.Context, *connect.Request[v1.SaveRequest]) (*connect.Response[v1.SaveResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autokitteh.builds.v1.BuildsService.Save is not implemented"))
}

func (UnimplementedBuildsServiceHandler) Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autokitteh.builds.v1.BuildsService.Delete is not implemented"))
}

func (UnimplementedBuildsServiceHandler) Download(context.Context, *connect.Request[v1.DownloadRequest]) (*connect.Response[v1.DownloadResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("autokitteh.builds.v1.BuildsService.Download is not implemented"))
}
