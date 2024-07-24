package media

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"sade-backend/api/handlers"
	"sade-backend/api/models"
	"sade-backend/pkg/utility"
	"strconv"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("media")
	if err != nil {
		handlers.Log.Errorf("Error uploading file: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error uploading file!"})
		return
	}
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		handlers.Log.Info("User not logged in")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	email := c.PostForm("email")
	priceValue := c.PostForm("price")
	price, err := strconv.ParseFloat(priceValue, 64)
	link, err := utility.GenerateLink()
	if err != nil {
		handlers.Log.Errorf("Error generating link: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error generating link!"})
		return
	}

	fileType := utility.CheckFileType(file.Filename)
	mediaDir := "../media/"

	originalDir := filepath.Join(mediaDir, "original")
	previewsDir := filepath.Join(mediaDir, "previews")
	err = os.MkdirAll(originalDir, 0775)
	if err != nil {
		handlers.Log.Errorf("Error creating original media directory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create original media directory!"})
		return
	}
	err = os.MkdirAll(previewsDir, 0775)
	if err != nil {
		handlers.Log.Errorf("Error creating previews media directory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create previews media directory!"})
		return
	}
	originalPath := filepath.Join(originalDir, file.Filename)
	previewsPath := filepath.Join(previewsDir, "preview_"+file.Filename)
	err = c.SaveUploadedFile(file, originalPath)
	if err != nil {
		handlers.Log.Errorf("Error uploading file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot upload file!"})
		return
	}
	err = handlers.Fw.ApplyWatermark(fileType, originalPath, previewsPath)
	if err != nil {
		handlers.Log.Errorf("Error apply watermark: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot apply watermark!"})
		return
	}
	media := models.Media{
		UploadedBy:   userID.(int64),
		ClientEmail:  email,
		PreviewPath:  previewsPath,
		OriginalPath: originalPath,
		Price:        price,
		Locked:       true,
	}
	err = handlers.A.RegisterUser(email, link, "client", "", "")
	if err != nil {
		_ = handlers.A.LoginUser(email, link)
	}

	data := map[string]interface{}{
		"uploaded_by":   media.UploadedBy,
		"client_email":  media.ClientEmail,
		"preview_path":  media.PreviewPath,
		"original_path": media.OriginalPath,
		"price":         media.Price,
		"locked":        media.Locked,
	}
	err = handlers.MediaTable.CmdInsert(data)
	if err != nil {
		handlers.Log.Errorf("Error inserting media: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot insert media!"})
		return
	}

	previewLink := fmt.Sprintf("http://localhost:8080/validate/%s?preview_path=%s", link, previewsPath)
	err = handlers.N.SendLink(email, previewLink, "Media Preview", "../static/template/media_preview.html")
	if err != nil {
		handlers.Log.Errorf("Error sending email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot send email!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully uploaded!"})
}
