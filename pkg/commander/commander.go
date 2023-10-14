package commander

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	flag "github.com/spf13/pflag"
)

type Commander struct {
	args      []string
	Commands  []Command
	OutWriter io.Writer
	flags     *flag.FlagSet
}

func (cmder *Commander) Args() []string {
	if cmder.args == nil {
		cmder.args = os.Args
	}

	return cmder.args
}

func (cmder *Commander) Run(ctx context.Context) error {
	if len(cmder.Args()) < 1 {
		return ErrBadArgs
	}
	cmder.Flags().Parse(cmder.Args())

	if cmder.Args()[1] == "help" {
		cmder.DisplayHelp(ctx)
		return nil
	}

	for _, command := range cmder.Commands {
		if command.Name == cmder.Args()[1] {
			return command.Run(ctx, cmder.Args())
		}
	}

	return ErrCmdNotFound
}

func (cmder *Commander) AddCommand(ctx context.Context, command Command) {
	cmder.Commands = append(cmder.Commands, command)
}

func (cmder *Commander) DisplayHelp(ctx context.Context) {
	for index, command := range cmder.Commands {
		fmt.Fprintf(cmder.getOutWriter(), "[%v] \nCommand: %v \nDescription: %v\n", index+1, command.Name, command.Description)
	}
}

func (cmder *Commander) getOutWriter() io.Writer {
	if cmder.OutWriter == nil {
		log.Printf("[WARN] outWriter empty")
		return os.Stdout
	}

	return cmder.OutWriter
}

func (cmder *Commander) Flags() *flag.FlagSet {
	if cmder.flags == nil {
		cmder.flags = flag.NewFlagSet("Commander", flag.ContinueOnError)
	}

	return cmder.flags
}
