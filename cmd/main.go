// go:build linux
package main

import (
	"github.com/goRunner/internals/commands"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	cmd := os.Args[1]
	args := os.Args[2:]

	logger := zerolog.New(os.Stdout)
	logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	logger = logger.With().Timestamp().Logger()

	// command := commands.NewCommand(&logger, "ls", commands.WithArgs("-la"), commands.WithStdOut(), commands.WithNewNamespace())
	logger.Info().Msg("Starting command " + cmd)
	command := commands.NewCommand(&logger, cmd, commands.WithArgs(args...), commands.WithStdOut(), commands.WithNewNamespace())
	command.Run()
}