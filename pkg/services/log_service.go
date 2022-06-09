package services

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Log(message string, typeMsg string, args ...interface{}) {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	switch typeMsg {
	case "Info":
		log.Debug().
			Str("message", message).Send()
	case "Error":
		log.Error().
			Str("message", message).Send()
	default:
		log.Debug().
			Str("message", message).Send()

	}

}
