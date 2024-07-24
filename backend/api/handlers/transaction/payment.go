package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
	"net/http"
	"sade-backend/api/handlers"
	"sade-backend/api/models"
)

func Process(c *gin.Context) {
	var req models.Transaction
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	media, err := handlers.MediaTable.CmdRead("id", req.MediaID)
	if err != nil || len(media) == 0 {
		handlers.Log.Errorf("Error while reading media: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Media not found!"})
		return
	}

	mediaData := media[0]
	price, ok := mediaData["price"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid price"})
		return
	}
	amount := int64(price * 100)
	stripe.Key = handlers.StripeKey
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(amount),
		Currency:    stripe.String(req.Currency),
		Description: stripe.String(req.Description),
		Source:      &stripe.SourceParams{Token: stripe.String(req.Token)},
	}
	ch, err := charge.New(params)
	if err != nil {
		handlers.Log.Errorf("Error while charging: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment failed!"})
		return
	}

	data := map[string]interface{}{
		"locked": false,
	}
	err = handlers.MediaTable.CmdUpdate(data, int(req.MediaID))
	if err != nil {
		handlers.Log.Errorf("Error while updating media: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlock media!"})
		return
	}
	transactionData := map[string]interface{}{
		"admin_id":       mediaData["uploaded_by"],
		"media_id":       req.MediaID,
		"client_email":   mediaData["client_email"],
		"amount":         amount,
		"currency":       req.Currency,
		"description":    req.Description,
		"transaction_id": ch.ID,
	}
	err = handlers.TransactionTable.CmdInsert(transactionData)
	if err != nil {
		handlers.Log.Errorf("Error while inserting transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log transaction!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment successful!"})
}
