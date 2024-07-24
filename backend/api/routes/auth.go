package routes

import (
	"github.com/gin-gonic/gin"
	"sade-backend/api/handlers/authentication"
)

func AuthRoute(r *gin.Engine) {
	r.POST("/register", authentication.Register)
	r.POST("/login", authentication.Login)
	r.GET("/verify/:link", authentication.Verify)
	r.POST("/save", authentication.SaveSession)
	r.GET("/check", authentication.Check)
	r.POST("/logout", authentication.Logout)
}
