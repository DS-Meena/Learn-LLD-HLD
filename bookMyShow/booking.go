package main

import "fmt"

type Booking struct {
	ID     string
	UserId string
	ShowId string
	Seats  []string
	Status string
}

type BookingService struct{}

func (s *BookingService) BookSeats(userId, showId string, seats []string) error {
	fmt.Printf("Booking seats for user %s, show %s, seats %v\n", userId, showId, seats)
	fmt.Println("Starting database stransaction")

	for _, seatNumber := range seats {
		fmt.Printf("Checking seat availability for %s\n", seatNumber)
		fmt.Printf("Reserving seat %s for user %s\n", seatNumber, userId)
	}

	fmt.Println("Creating booking entry in database")
	fmt.Println("Committing database transaction")

	return nil
}
