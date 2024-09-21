import time
from parking_lot import *
from utils import Logger


def main():
    # initialize parking lot
    parking_lot = ParkingLot()
    logger = Logger()

    # add floors and parking spots
    for i in range(3):
        floor = Floor(i)

        # add 3 spots for each type of vehicle
        for vehicle_type in ["car", "bike", "truck"]:
            for j in range(3):
                spot = ParkingSpot(f"{i}-{vehicle_type}-{j}", vehicle_type)
                floor.parking_spots.append(spot)

        parking_lot.floors.append(floor)
    
    # add entry and exit points
    parking_lot.entry_points.append(EntryPoint(1))
    parking_lot.exit_points.append(ExitPoint(1))

    # Simulate parking process
    car = Vehicle("ABC123", "car")
    bike = Vehicle("XYZ456", "bike")
    truck = Vehicle("DEF789", "truck")
    car2 = Vehicle("ABC789", "car")
    car3 = Vehicle("ABC456", "car")
    car4 = Vehicle("ABC741", "car")

    vehicles = [car, bike, truck, car2, car3, car4]

    for vehicle in vehicles:
        try:

            ticket = parking_lot.park_vehicle(vehicle)

            # if no exception, that means parked
            logger.log_event("Vehicle Parked", f"Vechile {vehicle.license_plate} parked in {ticket.spot.spot_id}")

            time.sleep(2)

            fee = parking_lot.remove_vehicle(ticket)
            logger.log_event("Vehicle Exited", f"Vechile {vehicle.license_plate} exited. Fee: {fee}")

        except Exception as e:
            print(f"Error: {str(e)}")

if __name__ == "__main__":
    main()