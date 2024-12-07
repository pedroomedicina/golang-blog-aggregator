package main

import (
	"blog_aggregator/command"
	"blog_aggregator/command/handlers"
	"blog_aggregator/internal/config"
	"fmt"
	"log"
	"os"
)

func main() {
	configuration, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	state := config.NewState(&configuration)
	commands := command.NewCommands()
	commands.Register("login", handlers.Login)

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

	fmt.Println("Command executed successfully")
}
