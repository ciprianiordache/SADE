package transaction

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sade-backend/api/handlers"
)

func GetTransactions(c *gin.Context) {
	session := sessions.Default(c)
	adminID := session.Get("user_id")

	transactions, err := handlers.TransactionTable.CmdRead("admin_id", adminID)
	if err != nil {
		handlers.Log.Errorf("Error while reading transactions: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read transactions"})
		return
	}

	for i := range transactions {
		if adminID != transactions[i]["admin_id"] {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to read this transaction"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions})
}
