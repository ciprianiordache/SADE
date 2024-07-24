package routes

import (
	"github.com/gin-gonic/gin"
	"sade-backend/api/handlers/media"
	"sade-backend/api/middleware"
)

func MediaRoute(r *gin.Engine) {
	r.POST("/upload", middleware.AuthMiddleware(), middleware.AuthorizationMiddleware("admin"), media.Upload)
	r.GET("/validate/:link", media.Validate)
	r.GET("/unlock/:id", middleware.AuthMiddleware(), middleware.AuthorizationMiddleware("client"), media.Unlock)
	r.GET("/media", middleware.AuthMiddleware(), middleware.AuthorizationMiddleware("admin"), media.GetMedia)
	r.GET("/preview", middleware.AuthMiddleware(), middleware.AuthorizationMiddleware("client"), media.PreviewMedia)
	r.GET("/original", middleware.AuthMiddleware(), middleware.AuthorizationMiddleware("client"), media.OriginalMedia)
	r.POST("/download/:id", middleware.AuthMiddleware(), middleware.AuthorizationMiddleware("client"), media.Download)
	r.POST("/resent", middleware.AuthMiddleware(), middleware.AuthorizationMiddleware("admin"), media.Resent)
}
