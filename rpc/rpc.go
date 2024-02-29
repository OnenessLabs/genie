package rpc

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/OnenessLabs/genie/rpc/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func NewGRPCServer(lc fx.Lifecycle) *grpc.Server {
	s := grpc.NewServer()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", ":9090")
			if err != nil {
				return err
			}
			fmt.Println("Starting GRPC server at :9090")
			api.Register(s)
			go s.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.GracefulStop()
			return nil
		},
	})
	return s
}

func NewHTTPServer(lc fx.Lifecycle) *http.Server {
	s := &http.Server{Addr: ":8080"}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", s.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at :8080")
			mux := runtime.NewServeMux()
			api.RegisterHandler(ctx, mux)
			s.Handler = mux
			go s.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return s.Shutdown(ctx)
		},
	})
	return s
}
