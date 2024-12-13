package handlers

import (
	"blog_aggregator/command"
	"blog_aggregator/internal/config"
	"blog_aggregator/internal/database"
	"context"
	"errors"
	"fmt"
)

func UnfollowFeed(s *config.State, cmd command.Command, user database.User) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("login command expects an argument: feed_url")
	}

	url := cmd.Arguments[0]
	fmt.Printf("Following URL: %s\n", url)

	err := s.Db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("error while unfollowing feed url: %s", err)
	}

	fmt.Printf("Unfollowed feed url: %s\n", url)
	return nil
}
