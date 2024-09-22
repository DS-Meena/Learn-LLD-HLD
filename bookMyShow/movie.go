package main

import "fmt"

type Movie struct {
	ID          string
	Title       string
	Description string
	Duration    int
}

// move service handles movie-related operations
type MovieService struct{}

func (s *MovieService) GetMovie(id string) (*Movie, error) {
	fmt.Printf("Fetching movie with ID: %s from database\n", id)

	// simulating database query
	movie := &Movie{
		ID:          id,
		Title:       "The Godfather",
		Description: "The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.",
		Duration:    175,
	}

	return movie, nil
}

func (s *MovieService) ListMovies() ([]*Movie, error) {
	fmt.Println("Fetching all movies from database")

	// simulating database query"
	movies := []*Movie{
		{ID: "1", Title: "Don", Description: "Surya", Duration: 120},
		{ID: "2", Title: "Tiger", Description: "Tigris", Duration: 180},
		{ID: "3", Title: "Padmaavat", Description: "All is well", Duration: 150},
	}

	return movies, nil
}
