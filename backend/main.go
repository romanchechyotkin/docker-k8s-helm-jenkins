package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"

	"golang.org/x/exp/slog"
)

var count uint64
var PORT = ":8001"

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/favicon.ico" {
		atomic.AddUint64(&count, 1)
	}

	fmt.Fprintf(w, "requests amount: %d\n", count)
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	logger.Info("BACKEND SERVICE")

	http.HandleFunc("/", index)

	logger.Info("service running on port", slog.String("port", PORT))

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		logger.Error("failed to run server", slog.String("error", err.Error()))
	}
}
