package authentication

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
)

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		handlers.Log.Info("Failed to logout!")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
