package main

import (
	"context"
	"github.com/Arh0rn/gotgaibot1/internal/app"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}
	if err := a.Run(); err != nil {
		panic(err)
	}
}
