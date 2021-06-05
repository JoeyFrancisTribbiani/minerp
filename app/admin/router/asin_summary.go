package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis/finance"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerErpFinanceAsinSummaryRouter)
}

// registerErpFinanceAsinSummaryRouter
func registerErpFinanceAsinSummaryRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &finance.ErpFinanceAsinSummary{}
	r := v1.Group("/finance/asin_summary").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetErpFinanceAsinSummaryList)
		r.GET("/:id", api.GetErpFinanceAsinSummary)
		r.POST("", api.InsertErpFinanceAsinSummary)
		r.PUT("/:id", api.UpdateErpFinanceAsinSummary)
		r.DELETE("", api.DeleteErpFinanceAsinSummary)
	}
}
