package commander

import (
	"context"
)

type Command struct {
	Name        string
	Description string
	Handler     func(ctx context.Context, args []string) error
}

func (c *Command) Run(ctx context.Context, args []string) error {
	return c.Handler(ctx, args)
}
