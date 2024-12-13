package middleware

import (
	"blog_aggregator/command"
	"blog_aggregator/internal/config"
	"blog_aggregator/internal/database"
	"context"
	"errors"
)

func LoggedIn(handler func(s *config.State, cmd command.Command, user database.User) error) func(state *config.State, command command.Command) error {
	return func(state *config.State, cmd command.Command) error {
		if state.Config.CurrentUserName == "" {
			return errors.New("no user is currently logged in")
		}

		user, err := state.Db.GetUser(context.Background(), state.Config.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(state, cmd, user)
	}
}
