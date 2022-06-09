package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Log(message string, args ...interface{}) {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print(message)
}
