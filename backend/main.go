package main

import (
	"github.com/gin-contrib/cors"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/config"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/db"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/router"
)

func main() {
	db.Connect()
	db.Migrate()

	router := router.SetupRoter()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.FRONTEND_URL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	router.Run(":" + config.PORT)
}
