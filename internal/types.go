package app

import (
	"github.com/ycodeclub/Nex/app"
	"github.com/ycodeclub/nex-demo/internal/transport"
)

// Module 表示 HTTP 模块对外暴露能力。
type Module struct {
	*transport.Handler
}

var _ app.HTTPModule = (*Module)(nil)
