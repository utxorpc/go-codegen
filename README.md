# go-codegen

Generated Go code for the [UTxO RPC](https://utxorpc.org) spec, produced from [`utxorpc/spec`](https://github.com/utxorpc/spec).

Contains the protobuf message types and [`connect-go`](https://connectrpc.com) service definitions (clients and handlers) for the UTxO RPC services.

## Looking to build a client?

If you're calling a UTxO RPC service from a Go app, use [`github.com/utxorpc/go-sdk`](https://github.com/utxorpc/go-sdk) instead. It wraps this package with a nicer client API and some examples to get you going. Start there.

## Server-side usage

This package is what you want if you're implementing a UTxO RPC service. Each service has a `*ServiceHandler` interface, along with an `Unimplemented*ServiceHandler` you can embed, in its `*connect` subpackage. Implement the interface, mount the handler on an `http.ServeMux`, and serve it over HTTP/2.

### Install

```sh
go get github.com/utxorpc/go-codegen
```

### Example: implementing `SyncService`

```go
package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/sync"
	"github.com/utxorpc/go-codegen/utxorpc/v1alpha/sync/syncconnect"
)

type syncServer struct {
	syncconnect.UnimplementedSyncServiceHandler
}

func (s *syncServer) ReadTip(
	ctx context.Context,
	req *connect.Request[sync.ReadTipRequest],
) (*connect.Response[sync.ReadTipResponse], error) {
	return connect.NewResponse(&sync.ReadTipResponse{
		// populate response fields from your backend
	}), nil
}

func main() {
	mux := http.NewServeMux()
	mux.Handle(syncconnect.NewSyncServiceHandler(&syncServer{}))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
	log.Fatal(srv.ListenAndServe())
}
```

Embedding `Unimplemented*ServiceHandler` means you only have to write the RPCs you actually implement; the rest return `CodeUnimplemented`.

Handlers from `New*ServiceHandler` support Connect, gRPC, and gRPC-Web, with Protobuf and JSON codecs and gzip.

### Available services

Under `github.com/utxorpc/go-codegen/utxorpc/v1alpha`:

- `sync` / `syncconnect` — chain sync (follow tip, dump history, fetch blocks)
- `query` / `queryconnect` — chain and ledger state queries
- `submit` / `submitconnect` — transaction submission and monitoring
- `watch` / `watchconnect` — tx-level event streaming
- `cardano` — Cardano-specific message types
- `bitcoin` — Bitcoin-specific message types

## License

See [LICENSE](LICENSE).
