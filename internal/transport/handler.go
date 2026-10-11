package transport

import (
	"github.com/ycodeclub/Nex/module"
	nexgrpc "github.com/ycodeclub/Nex/transport/grpc"
	nexhttp "github.com/ycodeclub/Nex/transport/http"
	"github.com/ycodeclub/nex-demo/internal/service"
	v1 "github.com/ycodeclub/nex-demo/proto/v1"
)

// Handler 提供 Auth 传输能力（HTTP / gRPC 共用 AuthServer）。
type Handler struct {
	v1.UnimplementedAuthServer
	svc service.AuthService
}

var (
	_ module.WebModule  = (*Handler)(nil)
	_ module.GRPCModule = (*Handler)(nil)
	_ v1.AuthServer     = (*Handler)(nil)
)

// NewHandler 创建传输层处理器。
func NewHandler(svc service.AuthService) *Handler {
	return &Handler{svc: svc}
}

// RegisterPublicRoutes 挂载 proto 生成的公共路由。
func (h *Handler) RegisterPublicRoutes(routes *nexhttp.RouteGroup) {
	v1.RegisterAuthHTTPPublic(routes, h)
}

// RegisterPrivateRoutes 挂载 proto 生成的私有路由（JWT）。
func (h *Handler) RegisterPrivateRoutes(routes *nexhttp.RouteGroup) {
	v1.RegisterAuthHTTPPrivate(routes, h)
}

// RegisterGRPC 向 Nex gRPC Server 注册 Auth 服务。
func (h *Handler) RegisterGRPC(server *nexgrpc.Server) {
	v1.RegisterAuthServer(server.GRPCServer(), h)
}
