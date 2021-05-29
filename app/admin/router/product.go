package router

import (
    "github.com/gin-gonic/gin"
    jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

    "go-admin/app/admin/models"
    "go-admin/app/admin/service/dto"
    "go-admin/common/actions"
    "go-admin/common/middleware"
)

func init()  {
	routerCheckRole = append(routerCheckRole, registerProductRouter)
}

// 需认证的路由代码
func registerProductRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
    r := v1.Group("/erp").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
    {
        model := &models.Product{}
        r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.ProductSearch), func() interface{} {
            list := make([]models.Product, 0)
            return &list
        }))
        r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.ProductById), nil))
        r.POST("", actions.CreateAction(new(dto.ProductControl)))
        r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.ProductControl)))
        r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.ProductById)))
    }
}
