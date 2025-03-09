//go:build !linux

package commands

import (
	"os"
	"github.com/rs/zerolog"
)

func WithNewNamespace() CmdOptions {
	logger := zerolog.New(os.Stdout)
	return func(c *Command) {
		logger.Info().Msg("Non-linux OS detected, not using new namespace")
	}
}