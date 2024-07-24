package routes

import (
	"github.com/gin-gonic/gin"
	"sade-backend/api/handlers/transaction"
	"sade-backend/api/middleware"
)

func PaymentRoute(r *gin.Engine) {
	r.POST("/process", middleware.AuthMiddleware(), middleware.AuthorizationMiddleware("client"), transaction.Process)
	r.GET("/transaction", middleware.AuthMiddleware(), middleware.AuthorizationMiddleware("admin"), transaction.GetTransactions)
}
