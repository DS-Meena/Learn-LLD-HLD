# class ParkingLot:
    

#     def park_vehicle(self, vehicle):
#         for floor in self.floors:


class Floor:
    floorNumber = 0
    parkingSpots = []

    def find_available_spot(self, vehicle_type):
        for spot in self.parkingSpots:
            if spot.type == vehicle_type and not spot.is_occupied:
                return spot
        return None
