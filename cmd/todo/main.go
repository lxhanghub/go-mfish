package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	// "github.com/lxhanghub/newb/pkg/cache"
	// "github.com/lxhanghub/newb/pkg/database"
	"github.com/lxhanghub/newb/internal/todo/grpcapi/hello"
	"github.com/lxhanghub/newb/pkg/host"
	"go.uber.org/zap"
	//_ "newb-sample/api/todo/docs" // swagger 一定要有这行,指向你的文档地址
)

func main() {

	// 创建服务主机构建器
	builder := host.NewWebHostBuilder()

	// 配置应用配置,内置环境变量读取和命令行参数读取
	builder.ConfigureAppConfiguration(func(build host.ConfigBuilder) {
		build.AddYamlFile("../../configs/config.yaml")
	})

	// 配置依赖注入
	// builder.ConfigureServices(database.PostgresModule())

	// builder.ConfigureServices(cache.RedisModule())

	//构建应用
	app, err := builder.Build()

	if err != nil {
		fmt.Printf("Failed to build application: %v\n", err)
		return
	}

	//配置请求中间件,支持跳过
	//app.UseMiddleware(middleware.NewAuthorizationMiddleware([]string{"/hello"}))

	//app.UseSwagger()

	// 配置路由
	app.MapRoutes(func(router *echo.Echo) {
		router.GET("/ping", func(c echo.Context) error {
			return c.JSON(200, map[string]string{
				"message": "hello world",
			})
		})
	})

	// Grpc 服务注册
	app.MapGrpcServices(hello.NewHelloService)

	// 运行应用
	if err := app.Run(); err != nil {
		app.Logger().Error("Error running application", zap.Error(err))
	}
}
