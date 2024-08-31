package authentication

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
	"sade-backend/api/models"
)

func SaveSession(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handlers.Log.Errorf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("first_name", user.FirstName)
	session.Set("last_name", user.LastName)
	session.Set("user_email", user.Email)
	session.Set("role", user.Role)
	session.Set("verified", user.Verified)
	if err := session.Save(); err != nil {
		handlers.Log.Errorf("Error saving session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving session."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
