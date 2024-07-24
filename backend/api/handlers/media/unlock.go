package media

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
	"strconv"
)

func Unlock(c *gin.Context) {
	mediaId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handlers.Log.Errorf("Fail to parse media id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse media id!"})
		return
	}
	mediaData, err := handlers.MediaTable.CmdRead("id", mediaId)
	if err != nil {
		handlers.Log.Errorf("Failed to read media data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read media data!"})
		return
	}
	c.JSON(http.StatusOK, mediaData[0])
}
