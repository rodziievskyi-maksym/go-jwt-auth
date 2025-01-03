package main

import (
	"github.com/rodziievskyi-maksym/go-jwt-auth/api"
	"log/slog"
	"os"
)

func main() {
	//what it oi.Writer?

	//add logs folder to env file
	//add log level to env file
	//add http address to env file

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	address := "localhost:8090"
	httpServer := api.NewHTTPServer(address)
	if err := httpServer.Start(); err != nil {
		slog.Error("failed to start HTTP server", err)
	}

}
