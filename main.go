package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/charanpy/todoapi/database"
	"github.com/charanpy/todoapi/router"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Todo API")

	godotenv.Load()

	database.Handler()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := database.Client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer database.Client.Disconnect(ctx)




	app := router.SetRoutes()

	app.Run(":3000")

}
