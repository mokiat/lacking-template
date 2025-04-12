package main

import (
	"log/slog"
	"os"

	"github.com/mokiat/lacking-studio/studio"
)

func main() {
	if err := studio.Run(); err != nil {
		slog.Error("Error",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}
}
