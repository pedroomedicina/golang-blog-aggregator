package handlers

import (
	"blog_aggregator/command"
	"blog_aggregator/internal/config"
	"blog_aggregator/internal/database"
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "gator")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch feed: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var feed RSSFeed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, fmt.Errorf("failed to unmarshal feed: %w", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return &feed, nil
}

func scrapeFeeds(s *config.State) {
	nextFeedToFetch, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("failed to fetch the next feed: %v\n", err)
	}

	fmt.Printf("Fetching feed: %s (%s)\n", nextFeedToFetch.Name, nextFeedToFetch.Url)
	err = s.Db.MarkFeedFetched(context.Background(), nextFeedToFetch.ID)
	if err != nil {
		fmt.Printf("failed to mark feed as fetched: %v\n", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeedToFetch.Url)
	if err != nil {
		fmt.Printf("failed to fetch feed: %v\n", err)
	}

	for _, item := range feed.Channel.Item {
		uuid := uuid.New()
		now := time.Now()
		_, err := s.Db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid,
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: html.UnescapeString(item.Description), Valid: item.Description != ""},
			PublishedAt: sql.NullTime{Time: *rssParsedDate(item.PubDate), Valid: rssParsedDate(item.PubDate) != nil},
			FeedID:      nextFeedToFetch.ID,
		})

		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				fmt.Printf("Post already exists: %s", item.Link)
				continue
			}
			fmt.Printf("Error saving post: %v", err)
		}
	}
}

func Aggregate(s *config.State, cmd command.Command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("agg command expects a single argument: time_between_reqs")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %v", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func rssParsedDate(date string) *time.Time {
	t, err := time.Parse(time.RFC1123Z, date)
	if err != nil {
		return nil
	}
	return &t
}
