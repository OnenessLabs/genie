package rpc

import (
	"context"
	"fmt"
	"github.com/OnenessLabs/genie/config"
	"github.com/OnenessLabs/genie/rpc/api"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net"
)

func NewGRPCServer(lc fx.Lifecycle, cfg *config.Config) *grpc.Server {
	s := grpc.NewServer()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", cfg.AppConfig.GrpcAddr)
			if err != nil {
				return err
			}
			fmt.Println("Starting GRPC server at " + cfg.AppConfig.GrpcAddr)
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

//func NewHTTPServer(lc fx.Lifecycle, cfg *config.Config) *http.Server {
//	s := &http.Server{Addr: cfg.AppConfig.HttpAddr}
//	lc.Append(fx.Hook{
//		OnStart: func(ctx context.Context) error {
//			ln, err := net.Listen("tcp", s.Addr)
//			if err != nil {
//				return err
//			}
//			fmt.Println("Starting HTTP server at :8080")
//			mux := runtime.NewServeMux()
//			api.RegisterHandler(ctx, mux)
//			s.Handler = mux
//			go s.Serve(ln)
//			return nil
//		},
//		OnStop: func(ctx context.Context) error {
//			return s.Shutdown(ctx)
//		},
//	})
//	return s
//}
