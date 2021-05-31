package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis/finance"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTransactionDetailRouter)
}

// registerTransactionDetailRouter
func registerTransactionDetailRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &finance.TransactionDetail{}
	r := v1.Group("/finance/transaction_detail").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetTransactionDetailList)
		r.GET("/:id", api.GetTransactionDetail)
		r.POST("", api.InsertTransactionDetail)
		r.PUT("/:id", api.UpdateTransactionDetail)
		r.DELETE("", api.DeleteTransactionDetail)
	}
}
