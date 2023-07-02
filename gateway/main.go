package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/exp/slog"
)

var GATEWAY_ID = 123123
var PORT = ":8000"

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	logger.Info("GATEWAY SERVICE")

	host := os.Getenv("BACKEND_HOST")
	port := os.Getenv("BACKEND_PORT")
	address := host + ":" + port

	http.HandleFunc("/", LoggerMiddleware(logger, address))

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		logger.Error("failed to run server", slog.String("error", err.Error()))
	}
}

func LoggerMiddleware(logger *slog.Logger, address string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" {
			return
		}

		logger.Info("incoming request")

		resp, err := http.Get(fmt.Sprintf("http://%s/", address))
		if err != nil {
			logger.Error("failed to make request", slog.String("error", err.Error()))
		}
		defer resp.Body.Close()

		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error("failed to decode body", slog.String("error", err.Error()))
		}

		res := map[string]interface{}{
			"response":   string(bytes),
			"gateway_id": GATEWAY_ID,
		}

		json.NewEncoder(w).Encode(res)
	}
}
