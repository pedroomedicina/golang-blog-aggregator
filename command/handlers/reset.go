package handlers

import (
	"blog_aggregator/command"
	"blog_aggregator/internal/config"
	"context"
	"fmt"
)

func Reset(s *config.State, _ command.Command) error {
	err := s.Db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete users: %v", err)
	}

	fmt.Println("Users have been deleted successfully")
	return nil
}
