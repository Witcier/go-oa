package router

import (
	v1 "witcier/go-api/api/v1"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (r *MenuRouter) InitRoleRouter(Router *gin.RouterGroup) {
	router := Router.Group("")
	roleApi := v1.ApiGroup.RoleApi
	{
		router.GET("/roles", roleApi.List)
		router.POST("/roles", roleApi.Store)
		router.PATCH("/roles/:id", roleApi.Update)
		router.DELETE("/roles/:id", roleApi.Delete)
		router.GET("/roles/menu/:id", roleApi.GetMenu)
		router.PATCH("/roles/menu/:id", roleApi.Menu)
		router.GET("/roles/permission/:id", roleApi.GetPermission)
		router.PATCH("/roles/permission/:id", roleApi.Permission)
	}
}
