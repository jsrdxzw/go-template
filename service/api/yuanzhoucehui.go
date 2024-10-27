package api

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"yuanzhoucehui.com/be/service/api/internal/handler"
	"yuanzhoucehui.com/be/service/api/internal/middleware"
	"yuanzhoucehui.com/be/service/api/internal/svc"
	"yuanzhoucehui.com/be/service/app"
)

func CreateServer(env string) (*rest.Server, *svc.ServiceContext) {
	envConfig := app.LoadConfig(env)
	ctx := svc.NewServiceContext(envConfig, env)
	server := rest.MustNewServer(envConfig.RestConf,
		rest.WithNotAllowedHandler(middleware.NewCorsMiddleware().Handler()))
	defer server.Stop()

	err := app.InitExtensions(envConfig, env)

	if err != nil {
		app.Log.Panic(fmt.Errorf("init extensions failed: %w", err))
	}

	handler.RegisterHandlers(server, ctx)

	return server, ctx
}

func StartHTTP(env string) {
	server, ctx := CreateServer(env)
	app.Log.Info(fmt.Sprintf("Starting server at %s:%d...", ctx.Config.Host, ctx.Config.Port))
	server.Start()
}
