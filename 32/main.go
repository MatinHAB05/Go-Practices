package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	// Create a new logger instance
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Log messages
	logger.Info().Str("key", "value").Msg("This is an info message")
	logger.Error().Err(fmt.Errorf("An error occurred")).Msg("This is an error message")




}
