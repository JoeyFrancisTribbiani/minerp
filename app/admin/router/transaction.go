package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis/finance"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerErpFinanceTransactionRouter)
}

// registerErpFinanceTransactionRouter
func registerErpFinanceTransactionRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &finance.ErpFinanceTransaction{}
	r := v1.Group("/finance/transaction").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetErpFinanceTransactionList)
		r.GET("/:id", api.GetErpFinanceTransaction)
		r.POST("", api.InsertErpFinanceTransaction)
		r.PUT("/:id", api.UpdateErpFinanceTransaction)
		r.DELETE("", api.DeleteErpFinanceTransaction)
	}
}
