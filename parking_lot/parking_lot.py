from datetime import datetime

from utils import calculate_fee_based_on_duration_and_vehicle_type
from vehicle import Vehicle
import uuid

class ParkingLot:
    def __init__(self):
        self.floors = []
        self.entry_points = []
        self.exit_points = []

    def park_vehicle(self, vehicle):
        for floor in self.floors:

            # find available spot if any
            spot = floor.isParkingAvailable(vehicle.type)
            if spot:
                # generate random ticket id
                ticket_id = uuid.uuid4()

                ticket = Ticket(ticket_id, vehicle, spot)
                spot.occupy()
                return ticket
        
        raise Exception("No available spots for the vehicle")
    
    def remove_vehicle(self, ticket):
        fee = ticket.calculate_fee()
        ticket.spot.vacate()
        return fee

class Floor:
    def __init__(self, floor_number):
        self.floor_number = floor_number
        self.parking_spots = []

    # return any available spot for the given vehicle type if any
    def isParkingAvailable(self, vehicle_type):
        for spot in self.parking_spots:
            if spot.type == vehicle_type and not spot.is_occupied:
                return spot
        return None

# each floor will have many parking spots
class ParkingSpot:
    def __init__(self, spot_id, type):
        self.spot_id = spot_id
        self.type = type
        self.is_occupied = False

    def occupy(self):
        self.is_occupied = True

    def vacate(self):
        self.is_occupied = False

class Ticket:
    def __init__(self, ticket_id, vehicle, spot):
        self.ticket_id = ticket_id
        self.vehicle = vehicle
        self.spot = spot
        self.start_time = datetime.now()

    def calculate_fee(self):
        duration = datetime.now() - self.start_time
        return calculate_fee_based_on_duration_and_vehicle_type(duration, self.vehicle.type)
    
class EntryPoint:
    def __init__(self, entry_id):
        self.entry_id = entry_id
    
class ExitPoint:
    def __init__(self, exit_id):
        self.exit_id = exit_id