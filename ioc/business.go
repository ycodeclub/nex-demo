package ioc

import (
	"fmt"

	"github.com/ycodeclub/Nex/app"
	"github.com/ycodeclub/Nex/module"
	"github.com/ycodeclub/Nex/stack"
	nexjwt "github.com/ycodeclub/nex-contrib/jwt"
	appmod "github.com/ycodeclub/nex-demo/internal"
	"github.com/ycodeclub/nex-demo/internal/repository"
	"github.com/ycodeclub/nex-demo/internal/service"
	"github.com/ycodeclub/nex-demo/internal/transport"
)

// provideModules 是 Composition Root：构造函数注入业务依赖。
func provideModules(settings stack.Settings) (app.HTTPModule, module.GRPCModule, error) {
	secret := settings.Auth.JWT.Secret
	if secret == "" {
		return nil, nil, fmt.Errorf("auth.jwt.secret 未配置")
	}
	var jwtOpts []nexjwt.Option
	if iss := settings.Auth.JWT.Issuer; iss != "" {
		jwtOpts = append(jwtOpts, nexjwt.WithIssuer(iss))
	}
	tokens := nexjwt.NewManager(secret, jwtOpts...)

	repo := repository.NewRepository()
	svc := service.NewAuthService(repo, tokens)
	handler := transport.NewHandler(svc)
	return appmod.NewModule(handler), appmod.NewGRPCModule(handler), nil
}
