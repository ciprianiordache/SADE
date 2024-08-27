package media

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
	"sade-backend/pkg/utility"
)

func Resent(c *gin.Context) {
	email := c.PostForm("email")
	mediaID := c.PostForm("media_id")
	link, err := utility.GenerateLink()
	if err != nil {
		handlers.Log.Errorf("Unable to generate link: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate link!"})
		return
	}
	mediaData, err := handlers.MediaTable.CmdRead("id", mediaID)
	if err != nil {
		handlers.Log.Errorf("Unable to read media data: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read media data!"})
		return
	}
	media := mediaData[0]
	err = handlers.A.LoginUser(email, "", link)
	if err != nil {
		handlers.Log.Errorf("Unable to login: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to login!"})
		return
	}
	previewLink := fmt.Sprintf("http://localhost:8080/validate/%s?preview_path=%s", link, media["preview_path"])
	err = handlers.N.SendLink(email, previewLink, "Media Preview", "../static/template/media_preview.html")
	if err != nil {
		handlers.Log.Errorf("Unable to send preview: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to send preview!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully resent preview"})
}
