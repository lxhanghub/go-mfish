package webapi

import (
	"github.com/labstack/echo/v4"
	"github.com/lxhanghub/go-workit/pkg/api"
	"go.uber.org/zap"
)

func Hello(
	router *echo.Echo, //echo
	log *zap.Logger, // 日志
) {

	// 创建路由组
	group := router.Group("/hello")

	// 创建路由

	group.GET("", HelloNewb(log))
}

// HelloNewb godoc
// @Summary hello Newb
// @Description 返回 "hello newb"
// @Tags Hello
// @Accept json
// @Produce json
// @Success 200 {object} api.Response[string]
// @Router /hello [get]
func HelloNewb(log *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := api.Success("hello newb")
		return c.JSON(200, data)
	}
}
