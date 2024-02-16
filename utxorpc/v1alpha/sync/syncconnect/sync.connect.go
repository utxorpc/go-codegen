// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: utxorpc/v1alpha/sync/sync.proto

package syncconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	sync "github.com/utxorpc/go-codegen/utxorpc/v1alpha/sync"
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
	// ChainSyncServiceName is the fully-qualified name of the ChainSyncService service.
	ChainSyncServiceName = "utxorpc.v1alpha.sync.ChainSyncService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ChainSyncServiceFetchBlockProcedure is the fully-qualified name of the ChainSyncService's
	// FetchBlock RPC.
	ChainSyncServiceFetchBlockProcedure = "/utxorpc.v1alpha.sync.ChainSyncService/FetchBlock"
	// ChainSyncServiceDumpHistoryProcedure is the fully-qualified name of the ChainSyncService's
	// DumpHistory RPC.
	ChainSyncServiceDumpHistoryProcedure = "/utxorpc.v1alpha.sync.ChainSyncService/DumpHistory"
	// ChainSyncServiceFollowTipProcedure is the fully-qualified name of the ChainSyncService's
	// FollowTip RPC.
	ChainSyncServiceFollowTipProcedure = "/utxorpc.v1alpha.sync.ChainSyncService/FollowTip"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	chainSyncServiceServiceDescriptor           = sync.File_utxorpc_v1alpha_sync_sync_proto.Services().ByName("ChainSyncService")
	chainSyncServiceFetchBlockMethodDescriptor  = chainSyncServiceServiceDescriptor.Methods().ByName("FetchBlock")
	chainSyncServiceDumpHistoryMethodDescriptor = chainSyncServiceServiceDescriptor.Methods().ByName("DumpHistory")
	chainSyncServiceFollowTipMethodDescriptor   = chainSyncServiceServiceDescriptor.Methods().ByName("FollowTip")
)

// ChainSyncServiceClient is a client for the utxorpc.v1alpha.sync.ChainSyncService service.
type ChainSyncServiceClient interface {
	FetchBlock(context.Context, *connect.Request[sync.FetchBlockRequest]) (*connect.Response[sync.FetchBlockResponse], error)
	DumpHistory(context.Context, *connect.Request[sync.DumpHistoryRequest]) (*connect.Response[sync.DumpHistoryResponse], error)
	FollowTip(context.Context, *connect.Request[sync.FollowTipRequest]) (*connect.ServerStreamForClient[sync.FollowTipResponse], error)
}

// NewChainSyncServiceClient constructs a client for the utxorpc.v1alpha.sync.ChainSyncService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewChainSyncServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ChainSyncServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &chainSyncServiceClient{
		fetchBlock: connect.NewClient[sync.FetchBlockRequest, sync.FetchBlockResponse](
			httpClient,
			baseURL+ChainSyncServiceFetchBlockProcedure,
			connect.WithSchema(chainSyncServiceFetchBlockMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		dumpHistory: connect.NewClient[sync.DumpHistoryRequest, sync.DumpHistoryResponse](
			httpClient,
			baseURL+ChainSyncServiceDumpHistoryProcedure,
			connect.WithSchema(chainSyncServiceDumpHistoryMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		followTip: connect.NewClient[sync.FollowTipRequest, sync.FollowTipResponse](
			httpClient,
			baseURL+ChainSyncServiceFollowTipProcedure,
			connect.WithSchema(chainSyncServiceFollowTipMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// chainSyncServiceClient implements ChainSyncServiceClient.
type chainSyncServiceClient struct {
	fetchBlock  *connect.Client[sync.FetchBlockRequest, sync.FetchBlockResponse]
	dumpHistory *connect.Client[sync.DumpHistoryRequest, sync.DumpHistoryResponse]
	followTip   *connect.Client[sync.FollowTipRequest, sync.FollowTipResponse]
}

// FetchBlock calls utxorpc.v1alpha.sync.ChainSyncService.FetchBlock.
func (c *chainSyncServiceClient) FetchBlock(ctx context.Context, req *connect.Request[sync.FetchBlockRequest]) (*connect.Response[sync.FetchBlockResponse], error) {
	return c.fetchBlock.CallUnary(ctx, req)
}

// DumpHistory calls utxorpc.v1alpha.sync.ChainSyncService.DumpHistory.
func (c *chainSyncServiceClient) DumpHistory(ctx context.Context, req *connect.Request[sync.DumpHistoryRequest]) (*connect.Response[sync.DumpHistoryResponse], error) {
	return c.dumpHistory.CallUnary(ctx, req)
}

// FollowTip calls utxorpc.v1alpha.sync.ChainSyncService.FollowTip.
func (c *chainSyncServiceClient) FollowTip(ctx context.Context, req *connect.Request[sync.FollowTipRequest]) (*connect.ServerStreamForClient[sync.FollowTipResponse], error) {
	return c.followTip.CallServerStream(ctx, req)
}

// ChainSyncServiceHandler is an implementation of the utxorpc.v1alpha.sync.ChainSyncService
// service.
type ChainSyncServiceHandler interface {
	FetchBlock(context.Context, *connect.Request[sync.FetchBlockRequest]) (*connect.Response[sync.FetchBlockResponse], error)
	DumpHistory(context.Context, *connect.Request[sync.DumpHistoryRequest]) (*connect.Response[sync.DumpHistoryResponse], error)
	FollowTip(context.Context, *connect.Request[sync.FollowTipRequest], *connect.ServerStream[sync.FollowTipResponse]) error
}

// NewChainSyncServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewChainSyncServiceHandler(svc ChainSyncServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	chainSyncServiceFetchBlockHandler := connect.NewUnaryHandler(
		ChainSyncServiceFetchBlockProcedure,
		svc.FetchBlock,
		connect.WithSchema(chainSyncServiceFetchBlockMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	chainSyncServiceDumpHistoryHandler := connect.NewUnaryHandler(
		ChainSyncServiceDumpHistoryProcedure,
		svc.DumpHistory,
		connect.WithSchema(chainSyncServiceDumpHistoryMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	chainSyncServiceFollowTipHandler := connect.NewServerStreamHandler(
		ChainSyncServiceFollowTipProcedure,
		svc.FollowTip,
		connect.WithSchema(chainSyncServiceFollowTipMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/utxorpc.v1alpha.sync.ChainSyncService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ChainSyncServiceFetchBlockProcedure:
			chainSyncServiceFetchBlockHandler.ServeHTTP(w, r)
		case ChainSyncServiceDumpHistoryProcedure:
			chainSyncServiceDumpHistoryHandler.ServeHTTP(w, r)
		case ChainSyncServiceFollowTipProcedure:
			chainSyncServiceFollowTipHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedChainSyncServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedChainSyncServiceHandler struct{}

func (UnimplementedChainSyncServiceHandler) FetchBlock(context.Context, *connect.Request[sync.FetchBlockRequest]) (*connect.Response[sync.FetchBlockResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("utxorpc.v1alpha.sync.ChainSyncService.FetchBlock is not implemented"))
}

func (UnimplementedChainSyncServiceHandler) DumpHistory(context.Context, *connect.Request[sync.DumpHistoryRequest]) (*connect.Response[sync.DumpHistoryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("utxorpc.v1alpha.sync.ChainSyncService.DumpHistory is not implemented"))
}

func (UnimplementedChainSyncServiceHandler) FollowTip(context.Context, *connect.Request[sync.FollowTipRequest], *connect.ServerStream[sync.FollowTipResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("utxorpc.v1alpha.sync.ChainSyncService.FollowTip is not implemented"))
}
