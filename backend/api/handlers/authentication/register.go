package authentication

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
	"sade-backend/pkg/utility"
)

func Register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	link, err := utility.GenerateLink()
	if err != nil {
		handlers.Log.Errorf("Unable to generate link: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to generate link."})
		return
	}
	err = handlers.A.RegisterUser(email, link, "admin", firstName, lastName, password)
	if err != nil {
		handlers.Log.Errorf("Unable to register user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to register user."})
		return
	}

	err = handlers.N.SendLink(email, link, "Register to SADE", "../static/template/register_link.html")
	if err != nil {
		handlers.Log.Errorf("Error sending register link: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending link."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Please check your email!"})
}
