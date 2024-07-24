package authentication

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"sade-backend/api/handlers"
)

func Verify(c *gin.Context) {
	link := c.Param("link")
	user, valid, err := handlers.A.ValidateLink(link)

	if !valid {
		handlers.Log.Info("Invalid link:", link)
		if err != nil {
			handlers.Log.Errorf("Error: %v", err)
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid link"})
		return
	}
	v := url.Values{}
	v.Set("id", fmt.Sprintf("%d", user.ID))
	v.Set("first_name", user.FirstName)
	v.Set("last_name", user.LastName)
	v.Set("email", user.Email)
	v.Set("role", user.Role)
	v.Set("verified", fmt.Sprintf("%t", user.Verified))

	c.Redirect(http.StatusFound, fmt.Sprintf("%s/confirm?", handlers.OriginPort)+v.Encode())
}
