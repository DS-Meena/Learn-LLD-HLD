package main

import (
	"fmt"
	"time"
)

func main() {
	movieService := &MovieService{}
	bookingService := &BookingService{}

	movie, err := movieService.GetMovie("1")
	if err != nil {
		fmt.Printf("Error getting movie: %v\n", err)
	} else {
		fmt.Printf("Retrieved movie: %+v\n", movie)
	}

	movies, err := movieService.ListMovies()
	if err != nil {
		fmt.Printf("Error listing movies: %v\n", err)
	} else {
		fmt.Printf("Retrieved movies:\n")
		for _, m := range movies {
			fmt.Printf("- %+v\n", m.Title)
		}
	}

	err = bookingService.BookSeats("user123", "show456", []string{"A1", "A2"})
	if err != nil {
		fmt.Printf("Error booking seats: %v\n", err)
	} else {
		fmt.Println("Seats booked successfully")
	}

	// concurrent bookings
	go bookingService.BookSeats("user1", "show1", []string{"B1", "B2"})
	go bookingService.BookSeats("user2", "show1", []string{"B2", "B3"})

	time.Sleep(time.Second)
}
