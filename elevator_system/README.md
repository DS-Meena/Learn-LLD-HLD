# Low Level Design: Elevator System 🛗

Our today's question is to design a Elevator system. with following requirements: 📋

- Multiple elevators in a building 🏢
- Users can request elevators from any floor 🔢
- Elevators can move up and down ⬆️⬇️
- Elevators have a maximum capacity 🧍‍♂️🧍‍♀️
- System should be efficient in handling requests ⚡

# Example

Let's understand our approach with an example. Let's say, a elevator system is their with 1 elevator. He is at 9th floor, wants to go to 3rd floor. The elevator is at 3rd floor.

Let's walk through this scenario step-by-step to illustrate how the elevator system would handle this request:

1. You're on the 9th floor and press the button to go down. This creates a new Request `Request{source = 9, destination = 9}`.
2. The ElevatorSystem receives this request and finds the nearest available elevator, which is currently at the 2nd floor.
3. The selected elevator adds this request to its queue. Its direction is set to UP since it needs to go up to reach you `state=MOVING and direction=UP(1)`.
4. The elevator starts moving up, incrementing its current_floor in each step until it reaches the 9th floor.
5. When the elevator reaches the 9th floor, it stops `state = STOPPED` and opens its doors.
6. You enter the elevator and press the button for the 3rd floor. This creates a new Request `Request{source = 9, destination = 3}`.
7. The elevator adds this new request to its queue, updates its direction  `direction = DOWN(-1)`, and closes its doors.
8. The elevator starts moving down, decrementing its current_floor in each step until it reaches the 3rd floor.
9. When the elevator reaches the 3rd floor, it stops again, opens its doors, and you exit.

# UML Diagram

Let's see what classes we would require:

```python
class ElevatorSystem{
    list of elevators

    # add reqeust to closest elevator's queue
    # and update its direction, depends on next closets request 
    requestElevator(floor)  
    
    # move all elevators
    # and udpate their direction if required (on reaching destination)
    step()     
}

class Elevator {
    currentFloor, direction (up, down, idle), state(moving, stopped)

    # contains the requests assigned, we process the requests in priority to closest destination
    # this requests could of picking up or dropping off passengers (treated same)
    Priority_queue
    
    # go to next floor (decided by direction enum)
    move()
    
    # add request to a priority queue (closest destinations first)
    # update direction 
    addRequest(request)
    
    # If reaches to current destination, then stops and update direction according to next request in queue
    # otherwise keep moving
    processNextRequest()
}

class Request {
    source and destination floor
}
```

![UML Diagram](image.png)

# Sequence Diagram

This diagram shows, how the system keeps running. 

![Sequence Diagram](image-1.png)

Let's break down the sequence diagram to understand how the elevator system works:
1. **User Requests an Elevator:** The process starts when a user requests an elevator by calling the `requestElevator(floor)` method on the ElevatorSystem.
2. **Finding the Nearest Elevator:** The ElevatorSystem then calls its internal method `findNearestElevator(floor)` to determine which elevator is closest to the requested floor.
3. **Adding the Request:** Once the nearest elevator is identified, the ElevatorSystem calls the `addRequest(request)` method on that specific Elevator, adding the user's request to the elevator's queue.
4. **Continuous Operation:** The diagram shows a loop that represents the continuous operation of the elevator system. In each time step:
- The ElevatorSystem calls the `move()` method on each Elevator, which updates the elevator's position.
- The ElevatorSystem then calls `processNextRequest()` on each Elevator, which handles picking up or dropping off passengers as needed.

This diagram is useful in understanding how the system keeps running.

## Running Code

```cpp
Elevator positions: 0 0 0  // request comes from 5th floor
Elevator positions: 1 0 0 
Elevator positions: 2 0 0 
Elevator positions: 3 0 0 
Elevator positions: 4 0 0 
Elevator positions: 5 0 0  // request comes from 1st floor
Elevator positions: 5 1 0  // request comes from 8th floor
Elevator positions: 6 1 0 
Elevator positions: 7 1 0 
Elevator positions: 8 1 0 
Elevator positions: 8 1 0 
Elevator positions: 8 1 0 
Elevator positions: 8 1 0 
Elevator positions: 8 1 0
```