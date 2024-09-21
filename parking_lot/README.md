# Parking Lot 🅿️🚗

Q: Create a LLD of parking lot

## Requirements 📋

We need to consider following requirements while designing the parking system:
- Multiple floors in the parking lot 🏢
- Different types of vehicles (car, bike, bus) 🚗🏍️🚌
- Multiple entry and exit points ⬅️➡️
- Parking spots of different sizes 🔲🔳
- Pricing based on vehicle type and duration ⏱️💰

# Core Classes

```python
Vehicle class {
    licensePlate, vehicleType
}

Floor class {
    floorNumber
    parkingSpots list
    isParkingAvailable(vehicleType)
} 

parkingSpot class {
    id, vehicleType, isOccupied
    occup()
    vacate()
}

Ticket class {
    id, vehicleType, spot, entryTime
    calculateFee()
}

EntryPoint class {
    id
}
ExitPoint class {
    id
}

ParkingLot class {
    list of floors, entryPoints, exitPoint
    parkVehicle(vehicle)
    removeVehicle(Ticket)
}
```

Let's check the UML for the above classes and requirements:
![UML Diagram](image.png)

# Parking Process Flow 🔄

Check the flow chart for parking a car in the parking lot:
![Parking Flow](image-1.png)


## How to Run 

You can use `python3 main.py` command to run the code.

Example:
```bash
# If vehicles don't exit
python3 main.py
2024-09-21 09:25:55 - Vehicle Parked: Vechile ABC123 parked in 0-car-0
2024-09-21 09:25:55 - Vehicle Parked: Vechile XYZ456 parked in 0-bike-0
2024-09-21 09:25:55 - Vehicle Parked: Vechile DEF789 parked in 0-truck-0
2024-09-21 09:25:55 - Vehicle Parked: Vechile ABC789 parked in 0-car-1
2024-09-21 09:25:55 - Vehicle Parked: Vechile ABC456 parked in 0-car-2
2024-09-21 09:25:55 - Vehicle Parked: Vechile ABC741 parked in 1-car-0
```

If vehicles exit
```bash
python3 main.py
2024-09-21 09:24:52 - Vehicle Parked: Vechile ABC123 parked in 0-car-0
2024-09-21 09:24:54 - Vehicle Exited: Vechile ABC123 exited. Fee: 20.0
2024-09-21 09:24:54 - Vehicle Parked: Vechile XYZ456 parked in 0-bike-0
2024-09-21 09:24:56 - Vehicle Exited: Vechile XYZ456 exited. Fee: 10.0
2024-09-21 09:24:56 - Vehicle Parked: Vechile DEF789 parked in 0-truck-0
2024-09-21 09:24:58 - Vehicle Exited: Vechile DEF789 exited. Fee: 30.0
2024-09-21 09:24:58 - Vehicle Parked: Vechile ABC789 parked in 0-car-0
2024-09-21 09:25:00 - Vehicle Exited: Vechile ABC789 exited. Fee: 20.0
2024-09-21 09:25:00 - Vehicle Parked: Vechile ABC456 parked in 0-car-0
2024-09-21 09:25:02 - Vehicle Exited: Vechile ABC456 exited. Fee: 20.0
2024-09-21 09:25:02 - Vehicle Parked: Vechile ABC741 parked in 0-car-0
2024-09-21 09:25:04 - Vehicle Exited: Vechile ABC741 exited. Fee: 20.0
```

# Troubleshooting

## How to Import method from file

1. Direct import: If the file you want to import is in the same directory, you can simply use: 
    
    `from parking_lot import Ticket`
2. Importing from subdirectory: If the file is in a subdirectory, you can use dot notation: 

    `from subdir.parking_lot import Ticket`

## How to import datetime

You can import datetime in two ways:
```python
import datetime
# to call a method
datetime.datetime.now()

# or
from datetime import datetime
datetime.now()
```

