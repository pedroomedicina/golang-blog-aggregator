package handlers

import (
	"blog_aggregator/command"
	"blog_aggregator/internal/config"
	"context"
	"fmt"
)

func ListUsers(s *config.State, _ command.Command) error {
	users, err := s.Db.GetUserNames(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list user names %v\n", err)
	}

	for _, user := range users {
		currentUserPlaceHolder := ""
		if user == s.Config.CurrentUserName {
			currentUserPlaceHolder = "(current)"
		}

		fmt.Printf("* %s %s\n", user, currentUserPlaceHolder)
	}

	return nil
}
