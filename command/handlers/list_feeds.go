package handlers

import (
	"blog_aggregator/command"
	"blog_aggregator/internal/config"
	"context"
	"fmt"
)

func ListFeeds(s *config.State, _ command.Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list feeds %v\n", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed:\n Name: %s\n Url: %s\n User Name: %s\n\n", feed.Name, feed.Url, feed.UserName)
	}

	return nil
}
