package handlers

import (
	"blog_aggregator/command"
	"blog_aggregator/internal/config"
	"blog_aggregator/internal/database"
	"context"
	"fmt"
	"strconv"
)

func Browse(s *config.State, cmd command.Command, user database.User) error {
	var limit int32
	if len(cmd.Arguments) > 0 {
		rawLimit, err := strconv.Atoi(cmd.Arguments[0])
		if err != nil {
			return err
		}

		limit = int32(rawLimit)
	} else {
		limit = 2
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("failed to get posts for user: %w", err)
	}

	fmt.Printf("showing the latest %d posts for user: %s \n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Url: %s\n", post.Url)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("Published Date: %s\n", post.PublishedAt.Time)
		fmt.Println("-")
	}

	return nil
}
