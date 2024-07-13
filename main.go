package main

import (
	"log"
	"stripe-api/stripe"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	router := gin.Default()

	// Registrar rutas de Stripe
	stripe.RegisterRoutes(router)

	// Iniciar el servidor
	router.Run(":8080")
}
