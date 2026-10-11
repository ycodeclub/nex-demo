package app

import (
	"github.com/ycodeclub/Nex/module"
	"github.com/ycodeclub/nex-demo/internal/transport"
)

// NewModule 组装 HTTP 模块。
func NewModule(handler *transport.Handler) *Module {
	return &Module{Handler: handler}
}

// NewGRPCModule 组装 gRPC 模块（与 HTTP 共用 Handler / AuthServer）。
func NewGRPCModule(handler *transport.Handler) module.GRPCModule {
	return handler
}
