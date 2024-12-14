# Gator CLI

Gator is a command-line tool for aggregating and managing RSS feeds. With Gator, you can fetch posts from your favorite RSS feeds, follow and unfollow feeds, and view aggregated posts – all from the comfort of your terminal.

## Prerequisites

To run Gator, you’ll need the following installed on your system:
1.	PostgreSQL: Installation instructions can be found [here](https://www.postgresql.org/download/).
2.	Go (version 1.18 or higher): Installation instructions can be found here [here](https://golang.org/doc/install).

___

## Installation

You can install the Gator CLI directly using Go:

```go install github.com/yourusername/gator@latest```

This will download, build, and install the Gator binary to your $GOPATH/bin. 
Ensure $GOPATH/bin is in your system’s PATH so you can run the gator command globally.

___

## Configuration

Before running Gator, you need to set up a configuration file. This file specifies the PostgreSQL database connection details and the current logged-in user as well.
1.	Create a .gatorconfig.json file in your home directory:
```touch ~/.gatorconfig.json```
2. Populate the file with your database connection URL:
```
{
  "db_url": "postgres://username:password@localhost:5432/dbname?sslmode=disable"
}
```

Replace:
•	username with your PostgreSQL username.
•	password with your PostgreSQL password.
•	dbname with the name of your database.

3. Ensure your database is set up and migrations have been applied:

```
goose up
```

___

## Running the Program

You can now run the Gator CLI. Here are some common commands:

Login

Log in as a user:

```gator login <username>```

If the user does not exist, you can create it by registering:

```gator register <username>```

___

## Manage Feeds

1. Add a Feed

Add a new RSS feed:

```gator addfeed "Feed Name" "https://example.com/rss"```

2. Follow a Feed

Follow a feed to receive its posts:

```gator follow "https://example.com/rss"```

3. Unfollow a feed

```gator unfollow "https://example.com/rss"```

4. View feeds you are following

```gator following```

___

## Aggregating posts

The agg command fetches posts from your followed feeds at regular intervals:

```gator agg 1m```

• Replace 1m with the desired interval (e.g., 1s, 1h, etc.).

• The command will run indefinitely, fetching posts and printing them to the console.

___

## View posts

View recent posts from your followed feeds:

```gator getposts <limit>```

• Replace <limit> with the number of posts you’d like to retrieve.

___

## Development and Contribution

Setting Up the Project

1.	Clone the repository:
```
git clone https://github.com/pedroomedicina/gator.git

cd gator
```
2. Install dependencies:

```
go mod tidy
```

3. Apply database migrations:
```
cd sql/schema

goose up
```

4. Run the CLI locally

```go run .```

___

## Testing

Test your changes manually by running Gator commands or write tests for critical functionality.

___

## License

This project is licensed under the MIT License. See the LICENSE file for details.