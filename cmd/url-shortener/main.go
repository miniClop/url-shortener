package main

import (
	"url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server:
}