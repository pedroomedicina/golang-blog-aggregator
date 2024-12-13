package handlers

import (
	"blog_aggregator/command"
	"blog_aggregator/internal/config"
	"blog_aggregator/internal/database"
	"context"
	"fmt"
)

func ListFollowing(s *config.State, _ command.Command, user database.User) error {
	feeds, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed fetching feeds followed by current user: %v", err)
	}

	fmt.Printf("Feeds followed by current user:\n")
	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.Name)
	}
	return nil
}
