package initialize

import (
	"net/http"
	"witcier/go-api/global"
	"witcier/go-api/middleware"
	"witcier/go-api/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	apiRouter := router.RouterGroupApp

	Router.StaticFS(global.Config.Local.Path, http.Dir(global.Config.Local.StorePath))

	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.Log.Info("use middleware cors")

	Router.Use(middleware.DefaultLimit())
	global.Log.Info("use middleware ip limit")

	Router.Use(middleware.Permission())
	global.Log.Info("use middleware permission")

	// 公共路由
	PublicRouter := Router.Group("")
	{
		PublicRouter.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
		apiRouter.InitLoginRouter(PublicRouter)
	}

	// 私有路由
	PrivateRouter := Router.Group("api")
	PrivateRouter.Use(middleware.Auth())
	{
		apiRouter.InitUserRouter(PrivateRouter)
		apiRouter.InitDepartmentRouter(PrivateRouter)
		apiRouter.InitPositionRouter(PrivateRouter)
		apiRouter.InitMenuRouter(PrivateRouter)
		apiRouter.InitPersonRouter(PrivateRouter)
		apiRouter.InitRoleRouter(PrivateRouter)
		apiRouter.InitPermissionRouter(PrivateRouter)
	}

	global.Log.Info("router register success")

	return Router
}
