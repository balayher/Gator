package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {

	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("could not fetch feed: %w", err)
	}

	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
