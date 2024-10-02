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

### Concurrency Handling 🔄

Before we look at how concurrency is handled. First we will check what problems could arise if concurrency is not handled properly.

1. **Dirty Read 💩**

    If Transaction A is reading the data which is writing by Transaction B and not yet even committed. If transaction B does the rollback, then whatever data read by Transaction A is known as dirty read.

2. **Non-Repeatable Read 🔁**

    If suppose Transaction A, reads the same row several times and there is a chance that it reads different value.

3. **Phantom Read 👻**

    If suppose Transaction A, executes same query several times and there is a chance that the rows returned are different.

Shared lock is used for reading and Exclusive lock is used for Writing. These locking mechanism helps us in avoiding the above above problems.

| Lock Type | Another Shared Lock | Another Exclusive Lock |
| --- | --- | --- |
| Have Shared Lock | Yes | No |
| Have Exclusive Lock | No | No |

For our problem of booking system, we will use Optimistic concurrency control (OCC) in the database transactions.

#### Optimistic Concurrency control (OCC)

- Isolation level used is below Repeatable Read i.e. Read committed.
- It has much higher concurrency than Pessimistic concurrency control.
- Also there is no change of deadlock with OCC.
- In case of conflic, overhead of transaction rollback and retry logic is implemented.

| Isolation Level 📊 | Locking Strategy 🔐 | Dirty Read Possible | Non-Repeatable Read Possible | Phantom Read Possible |
| --- | --- | --- | --- | --- |
| Read Uncommitted  (highest concurrency) | Read: No Lock acquired & Write: No Lock acquired | Yes | Yes | Yes |
| Read Committed | Read: Shared Lock acquired and Released as soon as Read is done & Write: Exclusive Lock acquired and keep till the end of the transaction | No | Yes | Yes |
| Repeatable Read | Read: Shared Lock acquired and Released only at the end of the Transaction & Write: Exclusive Lock acquired and Released only at the end of the Transaction | No | No | Yes |
| Serializable  (Least concurrency) | Same as Repeatable Read Locking Strategy & Write: Apply range Lock and lock is released only at the end of the transaction | No | No | No |

From above table, we can check what types of locking mechanism is used by Read committed Isolation level.

## Flow Chart 

![Flow chart](image-1.png)


## Sequene Diagram 🔢

![alt text](image-2.png)
*Sequence Diagram for bookMyShow sytem*

Let's explain the sequene diagram for the bookMyShow system:

1. User 🙋: Initiates the process by browsing movies
2. MovieService: Responds with a list of available movies
3. User: Selects a show and seats
4. BookingService 🎟️: Receives the user's selection
5. SeatService: Checks seat availability upon BookingService's request
6. BookingService: Receives seat availability confirmation
7. PaymentService: Processes the payment for the booking
8. BookingService: Receives payment confirmation
9. SeatService: Updates the seat status (booked)
10. NotificationService 📨: Sends booking confirmation to the user
11. User 🙋: Receives the booking confirmation
