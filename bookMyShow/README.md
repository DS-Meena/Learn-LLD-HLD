# Low Level Design of BookMyShow 🍿

Let's start by understanding the requirements of a movie ticket booking system.

- Uesr registration and authentication
- Brose movies and shows
- Book tickets with seat selection 
- Payment processing 💳
- Handle concurrent bookings

Our ticket booking system will have following components: 🏗️

- User Service: Handles user registration and authentication
- Movie Service: Manages movie and show information 🎬
- Booking Service: Handles ticket booking and seat selection 💺
- Payment Service: Processes payments
- Notification Service: Sends booking confirmations and updates

## Classes

We will only implmenet the functionalities of Movie service and Booking service.

```python
class movie {
    # features of any movie
    id, title, description, duration 
}  

class movieService {
    # fetch movie from database
    getMovie(id) (movie, error)

    # list all movies available in database
    listMovies() (movie[], error)
}

class booking {
    # features of any booking by user
    id, userId, showId, seats[], status
}

class bookingService {
    # book the given seats if available
    bookSeats(userId, showId, seatNumbers[])
}

```

## Flow Chart 

![Flow chart](image-1.png)


## Sequene Diagram


![alt text](image-2.png)
*Sequence Diagram for bookMyShow sytem*