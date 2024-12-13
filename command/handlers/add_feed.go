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

func AddFeed(s *config.State, cmd command.Command, user database.User) error {
	if len(cmd.Arguments) < 2 {
		return errors.New("add feed command expects two arguments: name, url")
	}

	name := cmd.Arguments[0]
	url := cmd.Arguments[1]

	newUuid := uuid.New()
	now := time.Now()
	feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        newUuid,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	feedFollow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	fmt.Printf("Feed created successfully: %+v\n", feed)
	fmt.Printf("Feed Follow created successfully: %+v\n", feedFollow)
	return nil
}
