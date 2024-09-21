def calculate_fee_based_on_duration_and_vehicle_type(duration, vehicle_type):
    if vehicle_type == "car":
        base_fee = 2.0
        per_minute_fee = 0.5
    elif vehicle_type == "bike":
        base_fee = 1.0
        per_minute_fee = 0.25
    else:
        return "Invalid vehicle type"

    total_fee = base_fee + (duration * per_minute_fee)
    return total_fee