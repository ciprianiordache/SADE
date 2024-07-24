package media

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
	"sade-backend/pkg/utility"
)

func GetMedia(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	media, err := handlers.MediaTable.CmdRead("uploaded_by", userID)
	if err != nil {
		handlers.Log.Errorf("Error while reading data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while reading data!"})
		return
	}
	for i := range media {
		media[i]["type"] = utility.DetermineMediaType(media[i]["preview_path"].(string))
	}
	c.JSON(http.StatusOK, gin.H{"media": media})
}
