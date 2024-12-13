package handlers

import (
	"blog_aggregator/command"
	"blog_aggregator/internal/config"
	"blog_aggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"time"
)

func FollowFeed(s *config.State, cmd command.Command, user database.User) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("login command expects a single argument: feed_url")
	}

	url := cmd.Arguments[0]
	fmt.Printf("Following URL: %s\n", url)

	feed, err := s.Db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}

	now := time.Now()
	feedFollow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %v", err)
	}

	fmt.Printf("Created feed follow successfully:\nFeed name: %s\n User: %s\n", feedFollow.FeedName, user.Name)

	return nil
}
