package handlers

import (
	"blog_aggregator/command"
	"blog_aggregator/internal/config"
	"errors"
	"fmt"
)

func Login(s *config.State, cmd command.Command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("login command expects a single argument: username")
	}

	username := cmd.Arguments[0]
	err := s.Config.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user %w", err)
	}

	fmt.Printf("User has been set to %s\n", username)
	return nil
}
