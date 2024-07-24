package media

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
)

func Download(c *gin.Context) {
	id := c.Param("id")
	session := sessions.Default(c)
	clientEmail := session.Get("user_email")
	result, err := handlers.MediaTable.CmdRead("id", id)
	if err != nil {
		handlers.Log.Infof("Media not found: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Media not found"})
		return
	}

	media := result[0]
	if media["locked"] == true {
		handlers.Log.Info("Media is not unlocked!")
		c.JSON(http.StatusForbidden, gin.H{"error": "Media is not unlocked!"})
		return
	}
	if media["client_email"] != clientEmail {
		handlers.Log.Info("You are not allowed to access this resource!")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to access this resource!"})
		return
	}

	originalPath := media["original_path"].(string)
	c.File(originalPath)
}
