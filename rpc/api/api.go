package api

import (
	"context"

	"github.com/OnenessLabs/genie/common"
	"github.com/OnenessLabs/genie/protobuf/gen/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type server struct {
	api.UnimplementedPublicServer
}

func (s *server) Version(ctx context.Context, _ *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return wrapperspb.String(common.GetAppVersion().WithCommit()), nil
}

func Register(s grpc.ServiceRegistrar) {
	api.RegisterPublicServer(s, &server{})
}

func RegisterHandler(ctx context.Context, mux *runtime.ServeMux) {
	api.RegisterPublicHandlerServer(ctx, mux, &server{})
}
