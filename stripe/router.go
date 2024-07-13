package stripe

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	stripeGroup := router.Group("/payment_intents")
	{
		stripeGroup.POST("/", CreatePaymentIntentHandler)
		stripeGroup.GET("/:id", RetrievePaymentIntentHandler)
		stripeGroup.POST("/:id/confirm", ConfirmPaymentIntentHandler)
		stripeGroup.POST("/:id/capture", CapturePaymentIntentHandler)
		stripeGroup.POST("/:id/cancel", CancelPaymentIntentHandler)
	}
}
