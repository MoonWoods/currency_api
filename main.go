package main

import (
	"log/slog"

	"github.com/MoonWoods/currency_api/pkg/api"
)

func main() {
	slog.Info("Starting up Currency API")
	api.SetupRouter()
	api.StartRouter()
}
