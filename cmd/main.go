// go:build linux
package main

import (
	"github.com/goRunner/internals/commands"
	"github.com/rs/zerolog"
	"os"
)

func main() {

	logger := zerolog.New(os.Stdout)
	logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	logger = logger.With().Timestamp().Logger()
	
	command := commands.NewCommand(&logger, "ls", commands.WithArgs("-la"), commands.WithStdOut(), commands.WithNewNamespace())

	command.Run()
}