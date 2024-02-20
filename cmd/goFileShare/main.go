package main

import (
	"goFileShare/internal/app"
)

func main() {
	r := app.SetupRouter()
	r.Run(":8080")
}
