package authentication

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
)

func Check(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	firstName := session.Get("first_name")
	lastName := session.Get("last_name")
	userEmail := session.Get("user_email")
	userRole := session.Get("role")
	verified := session.Get("verified")

	if userID == nil {
		handlers.Log.Info("No active session")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No active session!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "You are now connected!", "user_id": userID, "first_name": firstName, "last_name": lastName, "user_email": userEmail, "role": userRole, "verified": verified})

}
