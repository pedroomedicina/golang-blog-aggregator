package command

import (
	"blog_aggregator/internal/config"
	"errors"
)

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Handlers map[string]func(*config.State, Command) error
}

func NewCommands() *Commands {
	return &Commands{
		Handlers: make(map[string]func(*config.State, Command) error),
	}
}

func (c *Commands) Register(name string, handler func(*config.State, Command) error) {
	c.Handlers[name] = handler
}

func (c *Commands) Run(s *config.State, cmd Command) error {
	handler, exists := c.Handlers[cmd.Name]
	if !exists {
		return errors.New("command not found: " + cmd.Name)
	}
	return handler(s, cmd)
}
