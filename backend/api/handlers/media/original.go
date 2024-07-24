package media

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
	"sade-backend/pkg/utility"
)

func OriginalMedia(c *gin.Context) {
	mediaID := c.Query("media_id")
	fmt.Println(mediaID)
	session := sessions.Default(c)
	clientEmail := session.Get("user_email")

	media, err := handlers.MediaTable.CmdRead("id", mediaID)
	if err != nil {
		handlers.Log.Errorf("Unable to read media id: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read media id"})
		return
	}
	if media[0]["locked"] == true {
		handlers.Log.Info("Media is not unlocked!")
		c.JSON(http.StatusForbidden, gin.H{"error": "Media is not unlocked!"})
		return
	}
	if media[0]["client_email"] != clientEmail {
		handlers.Log.Info("Unauthorized user!")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to access this resource!"})
		return
	}
	media[0]["type"] = utility.DetermineMediaType(media[0]["preview_path"].(string))

	c.JSON(http.StatusOK, gin.H{"id": media[0]["id"], "original_path": media[0]["original_path"], "type": media[0]["type"]})
}
