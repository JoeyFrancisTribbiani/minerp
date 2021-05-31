package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis/finance"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerErpFinanceSummaryRouter)
}

// registerErpFinanceSummaryRouter
func registerErpFinanceSummaryRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &finance.ErpFinanceSummary{}
	r := v1.Group("/finance/summary").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetErpFinanceSummaryList)
		r.GET("/:id", api.GetErpFinanceSummary)
		r.POST("", api.InsertErpFinanceSummary)
		r.PUT("/:id", api.UpdateErpFinanceSummary)
		r.DELETE("", api.DeleteErpFinanceSummary)
	}
}
