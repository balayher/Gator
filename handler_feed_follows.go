package main

import (
	"context"
	"fmt"
	"time"

	"github.com/balayher/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	currentFeed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't retrieve feed: %w", err)
	}

	newFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    currentFeed.ID,
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), newFollow)
	if err != nil {
		return fmt.Errorf("couldn't create follow: %w", err)
	}

	fmt.Println("Follow created successfully:")
	printFeedFollow(follow.UserName, follow.FeedName)
	fmt.Println("=====================================")

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:     %s\n", username)
	fmt.Printf("* Feed:     %s\n", feedname)
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't retrieve feed follows: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("No feed follows found.")
		return nil
	}

	fmt.Printf("User %s is following:", user.Name)
	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}
	fmt.Println("=====================================")

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	currentFeed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't retrieve feed: %w", err)
	}

	unfollow := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: currentFeed.ID,
	}

	err = s.db.DeleteFeedFollow((context.Background()), unfollow)
	if err != nil {
		return fmt.Errorf("feed could not be unfollowed: %w", err)
	}

	fmt.Printf("%s successfully unfollowed!\n", currentFeed.Name)
	return nil
}
