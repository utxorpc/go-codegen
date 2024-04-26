/// A consistent view of the state of the ledger

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: utxorpc/v1alpha/query/query.proto

package queryconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	query "github.com/utxorpc/go-codegen/utxorpc/v1alpha/query"
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
	// QueryServiceName is the fully-qualified name of the QueryService service.
	QueryServiceName = "utxorpc.v1alpha.query.QueryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// QueryServiceReadParamsProcedure is the fully-qualified name of the QueryService's ReadParams RPC.
	QueryServiceReadParamsProcedure = "/utxorpc.v1alpha.query.QueryService/ReadParams"
	// QueryServiceReadUtxosProcedure is the fully-qualified name of the QueryService's ReadUtxos RPC.
	QueryServiceReadUtxosProcedure = "/utxorpc.v1alpha.query.QueryService/ReadUtxos"
	// QueryServiceSearchUtxosProcedure is the fully-qualified name of the QueryService's SearchUtxos
	// RPC.
	QueryServiceSearchUtxosProcedure = "/utxorpc.v1alpha.query.QueryService/SearchUtxos"
	// QueryServiceStreamUtxosProcedure is the fully-qualified name of the QueryService's StreamUtxos
	// RPC.
	QueryServiceStreamUtxosProcedure = "/utxorpc.v1alpha.query.QueryService/StreamUtxos"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	queryServiceServiceDescriptor           = query.File_utxorpc_v1alpha_query_query_proto.Services().ByName("QueryService")
	queryServiceReadParamsMethodDescriptor  = queryServiceServiceDescriptor.Methods().ByName("ReadParams")
	queryServiceReadUtxosMethodDescriptor   = queryServiceServiceDescriptor.Methods().ByName("ReadUtxos")
	queryServiceSearchUtxosMethodDescriptor = queryServiceServiceDescriptor.Methods().ByName("SearchUtxos")
	queryServiceStreamUtxosMethodDescriptor = queryServiceServiceDescriptor.Methods().ByName("StreamUtxos")
)

// QueryServiceClient is a client for the utxorpc.v1alpha.query.QueryService service.
type QueryServiceClient interface {
	ReadParams(context.Context, *connect.Request[query.ReadParamsRequest]) (*connect.Response[query.ReadParamsResponse], error)
	ReadUtxos(context.Context, *connect.Request[query.ReadUtxosRequest]) (*connect.Response[query.ReadUtxosResponse], error)
	SearchUtxos(context.Context, *connect.Request[query.SearchUtxosRequest]) (*connect.Response[query.SearchUtxosResponse], error)
	StreamUtxos(context.Context, *connect.Request[query.ReadUtxosRequest]) (*connect.ServerStreamForClient[query.ReadUtxosResponse], error)
}

// NewQueryServiceClient constructs a client for the utxorpc.v1alpha.query.QueryService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewQueryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) QueryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &queryServiceClient{
		readParams: connect.NewClient[query.ReadParamsRequest, query.ReadParamsResponse](
			httpClient,
			baseURL+QueryServiceReadParamsProcedure,
			connect.WithSchema(queryServiceReadParamsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		readUtxos: connect.NewClient[query.ReadUtxosRequest, query.ReadUtxosResponse](
			httpClient,
			baseURL+QueryServiceReadUtxosProcedure,
			connect.WithSchema(queryServiceReadUtxosMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		searchUtxos: connect.NewClient[query.SearchUtxosRequest, query.SearchUtxosResponse](
			httpClient,
			baseURL+QueryServiceSearchUtxosProcedure,
			connect.WithSchema(queryServiceSearchUtxosMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		streamUtxos: connect.NewClient[query.ReadUtxosRequest, query.ReadUtxosResponse](
			httpClient,
			baseURL+QueryServiceStreamUtxosProcedure,
			connect.WithSchema(queryServiceStreamUtxosMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// queryServiceClient implements QueryServiceClient.
type queryServiceClient struct {
	readParams  *connect.Client[query.ReadParamsRequest, query.ReadParamsResponse]
	readUtxos   *connect.Client[query.ReadUtxosRequest, query.ReadUtxosResponse]
	searchUtxos *connect.Client[query.SearchUtxosRequest, query.SearchUtxosResponse]
	streamUtxos *connect.Client[query.ReadUtxosRequest, query.ReadUtxosResponse]
}

// ReadParams calls utxorpc.v1alpha.query.QueryService.ReadParams.
func (c *queryServiceClient) ReadParams(ctx context.Context, req *connect.Request[query.ReadParamsRequest]) (*connect.Response[query.ReadParamsResponse], error) {
	return c.readParams.CallUnary(ctx, req)
}

// ReadUtxos calls utxorpc.v1alpha.query.QueryService.ReadUtxos.
func (c *queryServiceClient) ReadUtxos(ctx context.Context, req *connect.Request[query.ReadUtxosRequest]) (*connect.Response[query.ReadUtxosResponse], error) {
	return c.readUtxos.CallUnary(ctx, req)
}

// SearchUtxos calls utxorpc.v1alpha.query.QueryService.SearchUtxos.
func (c *queryServiceClient) SearchUtxos(ctx context.Context, req *connect.Request[query.SearchUtxosRequest]) (*connect.Response[query.SearchUtxosResponse], error) {
	return c.searchUtxos.CallUnary(ctx, req)
}

// StreamUtxos calls utxorpc.v1alpha.query.QueryService.StreamUtxos.
func (c *queryServiceClient) StreamUtxos(ctx context.Context, req *connect.Request[query.ReadUtxosRequest]) (*connect.ServerStreamForClient[query.ReadUtxosResponse], error) {
	return c.streamUtxos.CallServerStream(ctx, req)
}

// QueryServiceHandler is an implementation of the utxorpc.v1alpha.query.QueryService service.
type QueryServiceHandler interface {
	ReadParams(context.Context, *connect.Request[query.ReadParamsRequest]) (*connect.Response[query.ReadParamsResponse], error)
	ReadUtxos(context.Context, *connect.Request[query.ReadUtxosRequest]) (*connect.Response[query.ReadUtxosResponse], error)
	SearchUtxos(context.Context, *connect.Request[query.SearchUtxosRequest]) (*connect.Response[query.SearchUtxosResponse], error)
	StreamUtxos(context.Context, *connect.Request[query.ReadUtxosRequest], *connect.ServerStream[query.ReadUtxosResponse]) error
}

// NewQueryServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewQueryServiceHandler(svc QueryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	queryServiceReadParamsHandler := connect.NewUnaryHandler(
		QueryServiceReadParamsProcedure,
		svc.ReadParams,
		connect.WithSchema(queryServiceReadParamsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceReadUtxosHandler := connect.NewUnaryHandler(
		QueryServiceReadUtxosProcedure,
		svc.ReadUtxos,
		connect.WithSchema(queryServiceReadUtxosMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceSearchUtxosHandler := connect.NewUnaryHandler(
		QueryServiceSearchUtxosProcedure,
		svc.SearchUtxos,
		connect.WithSchema(queryServiceSearchUtxosMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceStreamUtxosHandler := connect.NewServerStreamHandler(
		QueryServiceStreamUtxosProcedure,
		svc.StreamUtxos,
		connect.WithSchema(queryServiceStreamUtxosMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/utxorpc.v1alpha.query.QueryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case QueryServiceReadParamsProcedure:
			queryServiceReadParamsHandler.ServeHTTP(w, r)
		case QueryServiceReadUtxosProcedure:
			queryServiceReadUtxosHandler.ServeHTTP(w, r)
		case QueryServiceSearchUtxosProcedure:
			queryServiceSearchUtxosHandler.ServeHTTP(w, r)
		case QueryServiceStreamUtxosProcedure:
			queryServiceStreamUtxosHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedQueryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedQueryServiceHandler struct{}

func (UnimplementedQueryServiceHandler) ReadParams(context.Context, *connect.Request[query.ReadParamsRequest]) (*connect.Response[query.ReadParamsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("utxorpc.v1alpha.query.QueryService.ReadParams is not implemented"))
}

func (UnimplementedQueryServiceHandler) ReadUtxos(context.Context, *connect.Request[query.ReadUtxosRequest]) (*connect.Response[query.ReadUtxosResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("utxorpc.v1alpha.query.QueryService.ReadUtxos is not implemented"))
}

func (UnimplementedQueryServiceHandler) SearchUtxos(context.Context, *connect.Request[query.SearchUtxosRequest]) (*connect.Response[query.SearchUtxosResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("utxorpc.v1alpha.query.QueryService.SearchUtxos is not implemented"))
}

func (UnimplementedQueryServiceHandler) StreamUtxos(context.Context, *connect.Request[query.ReadUtxosRequest], *connect.ServerStream[query.ReadUtxosResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("utxorpc.v1alpha.query.QueryService.StreamUtxos is not implemented"))
}
