package main

import (
	"context"

	"github.com/defany/auth-service/app/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a := app.NewApp()

	if err := a.Run(ctx); err != nil {
		panic(err)
	}
}
