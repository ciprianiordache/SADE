package authentication

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
	"sade-backend/pkg/utility"
)

func Login(c *gin.Context) {
	email := c.PostForm("email")
	link, err := utility.GenerateLink()
	if err != nil {
		handlers.Log.Errorf("Unable to generate link: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to generate link."})
		return
	}
	err = handlers.A.LoginUser(email, link)
	if err != nil {
		handlers.Log.Errorf("Unable to log in: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Unable to log in: %v", err)})
		return
	}

	err = handlers.N.SendLink(email, link, "Log In to SADE", "../static/template/login_link.html")
	if err != nil {
		handlers.Log.Errorf("error sending magic link: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending magic link."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link sent successfully, please check your email!"})
}
