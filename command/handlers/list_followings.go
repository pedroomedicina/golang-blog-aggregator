package handlers

import (
	"blog_aggregator/command"
	"blog_aggregator/internal/config"
	"context"
	"fmt"
)

func ListFollowing(s *config.State, _ command.Command) error {
	username := s.Config.CurrentUserName
	currentUser, err := s.Db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("error with current user %v", err)
	}

	feeds, err := s.Db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("failed fetching feeds followed by current user: %v", err)
	}

	fmt.Printf("Feeds followed by current user:\n")
	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.Name)
	}
	return nil
}
