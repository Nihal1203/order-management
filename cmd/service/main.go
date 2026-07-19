package main

import (
	"context"
	"log"

	app "github.com/Nihal1203/order-management-system/internal/app"
)

func main() {
	ctx := context.Background()
	err := app.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
