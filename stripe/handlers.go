package stripe

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v79"
)

func CreatePaymentIntentHandler(c *gin.Context) {
	var params stripe.PaymentIntentParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := NewStripeClient()
	paymentIntent, err := client.CreatePaymentIntent(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paymentIntent)
}

func RetrievePaymentIntentHandler(c *gin.Context) {
	id := c.Param("id")

	client := NewStripeClient()
	paymentIntent, err := client.RetrievePaymentIntent(id, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paymentIntent)
}

func ConfirmPaymentIntentHandler(c *gin.Context) {
	id := c.Param("id")
	var params stripe.PaymentIntentConfirmParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := NewStripeClient()
	paymentIntent, err := client.ConfirmPaymentIntent(id, &params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paymentIntent)
}

func CapturePaymentIntentHandler(c *gin.Context) {
	id := c.Param("id")
	var params stripe.PaymentIntentCaptureParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := NewStripeClient()
	paymentIntent, err := client.CapturePaymentIntent(id, &params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paymentIntent)
}

func CancelPaymentIntentHandler(c *gin.Context) {
	id := c.Param("id")
	var params stripe.PaymentIntentCancelParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := NewStripeClient()
	paymentIntent, err := client.CancelPaymentIntent(id, &params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paymentIntent)
}
