package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis/product"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerErpProductSpecificsRouter)
}

// registerErpProductSpecificsRouter
func registerErpProductSpecificsRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &product.ErpProductSpecifics{}
	r := v1.Group("/product/specifics").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetErpProductSpecificsList)
		r.GET("/:id", api.GetErpProductSpecifics)
		r.POST("", api.InsertErpProductSpecifics)
		r.PUT("/:id", api.UpdateErpProductSpecifics)
		r.DELETE("", api.DeleteErpProductSpecifics)
	}
}
