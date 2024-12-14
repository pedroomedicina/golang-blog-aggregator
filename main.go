package main

import (
	"blog_aggregator/command"
	"blog_aggregator/command/handlers"
	"blog_aggregator/internal/config"
	"blog_aggregator/internal/database"
	"blog_aggregator/middleware"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	configuration, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	state := config.NewState(&configuration)
	db, err := sql.Open("postgres", state.Config.DBURL)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	dbQueries := database.New(db)
	state.SetDb(dbQueries)

	commands := command.NewCommands()
	commands.Register("login", handlers.Login)
	commands.Register("register", handlers.Register)
	commands.Register("reset", handlers.Reset)
	commands.Register("users", handlers.ListUsers)
	commands.Register("agg", handlers.Aggregate)
	commands.Register("addfeed", middleware.LoggedIn(handlers.AddFeed))
	commands.Register("feeds", handlers.ListFeeds)
	commands.Register("follow", middleware.LoggedIn(handlers.FollowFeed))
	commands.Register("following", middleware.LoggedIn(handlers.ListFollowing))
	commands.Register("unfollow", middleware.LoggedIn(handlers.UnfollowFeed))
	commands.Register("browse", middleware.LoggedIn(handlers.Browse))

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("No enough arguments were provided. Usage: <command> [arguments]")
	}

	commandName := args[0]
	commandArgs := args[1:]
	userCommand := command.Command{
		Name:      commandName,
		Arguments: commandArgs,
	}

	err = commands.Run(state, userCommand)
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
