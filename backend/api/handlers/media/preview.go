package media

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
	"sade-backend/pkg/utility"
)

func PreviewMedia(c *gin.Context) {
	previewPath := c.Query("preview_path")
	if previewPath == "" {
		handlers.Log.Error("preview_path is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "No preview path provided"})
		return
	}

	media, err := handlers.MediaTable.CmdRead("preview_path", previewPath)
	if err != nil || len(media) == 0 {
		handlers.Log.Errorf("Media not found: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Media not found!"})
		return
	}

	media[0]["type"] = utility.DetermineMediaType(media[0]["preview_path"].(string))

	c.JSON(http.StatusOK, gin.H{"id": media[0]["id"], "preview_path": media[0]["preview_path"], "price": media[0]["price"], "type": media[0]["type"]})
}
