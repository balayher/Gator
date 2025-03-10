package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {

	err := s.db.ResetDB((context.Background()))
	if err != nil {
		return fmt.Errorf("database could not be reset: %w", err)
	}

	fmt.Println("Database reset successfully!")
	return nil
}
