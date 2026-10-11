package ioc

import (
	"context"
	"strings"

	"github.com/ycodeclub/Nex/log"
	"github.com/ycodeclub/Nex/middleware"
	"github.com/ycodeclub/Nex/stack"
	nexjwt "github.com/ycodeclub/nex-contrib/jwt"
	"github.com/ycodeclub/nex-contrib/middleware/logging"
	"github.com/ycodeclub/nex-contrib/middleware/recovery"
	reqidmw "github.com/ycodeclub/nex-contrib/middleware/requestid"
)

// Run 加载配置、装配传输层并启动服务（stack → app.New，无 Wire）。
func Run() error {
	cfg, settings, err := stack.Load("")
	if err != nil {
		return err
	}
	httpMod, grpcMod, err := provideModules(settings)
	if err != nil {
		return err
	}
	logger := log.NewDefaultLogger()
	opts := []stack.Option{
		stack.WithConfig(cfg),
		stack.WithSettings(settings),
		stack.WithMiddlewares(
			recovery.Middleware(recovery.WithLogger(logger)),
			reqidmw.Middleware(),
			logging.Middleware(logging.WithLogger(logger)),
		),
		stack.WithHTTPModule(httpMod),
		stack.WithGRPCModule(grpcMod),
	}
	if secret := settings.Auth.JWT.Secret; secret != "" {
		var jwtOpts []nexjwt.Option
		if iss := settings.Auth.JWT.Issuer; iss != "" {
			jwtOpts = append(jwtOpts, nexjwt.WithIssuer(iss))
		}
		// HTTP 私有路由 + gRPC；Login RPC 跳过验签（见 jwtExceptLogin）。
		opts = append(opts, stack.WithPrivateMiddlewares(jwtExceptLogin(secret, jwtOpts...)))
	}
	a, err := stack.New(opts...)
	if err != nil {
		return err
	}
	return a.Run()
}

// jwtExceptLogin 在 gRPC Login 上跳过 JWT（PrivateMw 会挂到全部 gRPC 方法）。
func jwtExceptLogin(secret string, opts ...nexjwt.Option) middleware.Middleware {
	inner := nexjwt.Server(secret, opts...)
	return func(next middleware.Handler) middleware.Handler {
		authed := inner(next)
		return func(ctx context.Context, req any) (any, error) {
			op := middleware.Operation(ctx)
			if strings.HasSuffix(op, "/Login") {
				return next(ctx, req)
			}
			return authed(ctx, req)
		}
	}
}
