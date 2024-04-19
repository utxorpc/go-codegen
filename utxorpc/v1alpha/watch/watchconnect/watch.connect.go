// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: utxorpc/v1alpha/watch/watch.proto

package watchconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	watch "github.com/utxorpc/go-codegen/utxorpc/v1alpha/watch"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// WatchServiceName is the fully-qualified name of the WatchService service.
	WatchServiceName = "utxorpc.v1alpha.watch.WatchService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// WatchServiceWatchTxProcedure is the fully-qualified name of the WatchService's WatchTx RPC.
	WatchServiceWatchTxProcedure = "/utxorpc.v1alpha.watch.WatchService/WatchTx"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	watchServiceServiceDescriptor       = watch.File_utxorpc_v1alpha_watch_watch_proto.Services().ByName("WatchService")
	watchServiceWatchTxMethodDescriptor = watchServiceServiceDescriptor.Methods().ByName("WatchTx")
)

// WatchServiceClient is a client for the utxorpc.v1alpha.watch.WatchService service.
type WatchServiceClient interface {
	WatchTx(context.Context, *connect.Request[watch.WatchTxRequest]) (*connect.ServerStreamForClient[watch.WatchTxResponse], error)
}

// NewWatchServiceClient constructs a client for the utxorpc.v1alpha.watch.WatchService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWatchServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) WatchServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &watchServiceClient{
		watchTx: connect.NewClient[watch.WatchTxRequest, watch.WatchTxResponse](
			httpClient,
			baseURL+WatchServiceWatchTxProcedure,
			connect.WithSchema(watchServiceWatchTxMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// watchServiceClient implements WatchServiceClient.
type watchServiceClient struct {
	watchTx *connect.Client[watch.WatchTxRequest, watch.WatchTxResponse]
}

// WatchTx calls utxorpc.v1alpha.watch.WatchService.WatchTx.
func (c *watchServiceClient) WatchTx(ctx context.Context, req *connect.Request[watch.WatchTxRequest]) (*connect.ServerStreamForClient[watch.WatchTxResponse], error) {
	return c.watchTx.CallServerStream(ctx, req)
}

// WatchServiceHandler is an implementation of the utxorpc.v1alpha.watch.WatchService service.
type WatchServiceHandler interface {
	WatchTx(context.Context, *connect.Request[watch.WatchTxRequest], *connect.ServerStream[watch.WatchTxResponse]) error
}

// NewWatchServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWatchServiceHandler(svc WatchServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	watchServiceWatchTxHandler := connect.NewServerStreamHandler(
		WatchServiceWatchTxProcedure,
		svc.WatchTx,
		connect.WithSchema(watchServiceWatchTxMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/utxorpc.v1alpha.watch.WatchService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WatchServiceWatchTxProcedure:
			watchServiceWatchTxHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWatchServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWatchServiceHandler struct{}

func (UnimplementedWatchServiceHandler) WatchTx(context.Context, *connect.Request[watch.WatchTxRequest], *connect.ServerStream[watch.WatchTxResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("utxorpc.v1alpha.watch.WatchService.WatchTx is not implemented"))
}