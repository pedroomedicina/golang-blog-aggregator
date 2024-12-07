package handlers

import (
	"blog_aggregator/command"
	"blog_aggregator/internal/config"
	"blog_aggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func Register(s *config.State, cmd command.Command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("register command expects a single argument: username")
	}

	username := cmd.Arguments[0]
	_, err := s.Db.GetUser(context.Background(), username)
	if err == nil {
		fmt.Printf("User with name '%s' already exists.\n", username)
		return fmt.Errorf("user with name '%s' already exists", username)
	}

	newUuid := uuid.New()
	now := time.Now()
	user, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        newUuid,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("failed to set current user: %w", err)
	}

	fmt.Printf("User created successfully: %+v\n", user)
	return nil
}
