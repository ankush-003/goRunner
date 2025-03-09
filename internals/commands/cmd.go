package commands

import (
	"os"
	"os/exec"

	"github.com/rs/zerolog"
)

type Command struct {
	*exec.Cmd
}

type CmdOptions func(*Command)

func WithArgs(args ...string) CmdOptions {
	return func(c *Command) {
		c.Args = args
	}
}

func WithStdOut() CmdOptions {
	return func(c *Command) {
		c.Stdout = os.Stdout
		c.Stdin = os.Stdin
		c.Stderr = os.Stderr
	}
}

func NewCommand(logger *zerolog.Logger, command string, opts ...CmdOptions) *Command {
	sublogger := logger.With().Str("command", command).Logger()	
	sublogger.Info().
		Str("command", command).
		Msg("Creating new command")
	cmd := &Command{
		Cmd: exec.Command(command),
	}

	for _, opt := range opts {
		opt(cmd)
	}
	
	return cmd
}