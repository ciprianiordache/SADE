package media

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
)

func Validate(c *gin.Context) {
	link := c.Param("link")
	previewPath := c.Query("preview_path")
	client, ok, err := handlers.A.ValidateLink(link)
	if err != nil {
		handlers.Log.Errorf("Error validating link: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error validating link"})
		return
	}
	if !ok {
		handlers.Log.Infof("Invalid link: %v", link)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Link invalid"})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", client.ID)
	session.Set("user_email", client.Email)
	session.Set("role", client.Role)
	session.Set("verified", client.Verified)
	err = session.Save()
	if err != nil {
		handlers.Log.Errorf("Error saving session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving session"})
		return
	}

	redirectURL := fmt.Sprintf("%s/preview?preview_path=%s", handlers.OriginPort, previewPath)
	c.Redirect(http.StatusFound, redirectURL)
}
