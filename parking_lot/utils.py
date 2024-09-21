from datetime import datetime


def calculate_fee_based_on_duration_and_vehicle_type(duration, vehicle_type):
    # convert time delta to hours
    hours = duration.total_seconds() / 3600

    base_fee = 10.0
    hourly_rate = 5.0

    if vehicle_type == "car":
        base_fee *= 2.0
        hourly_rate *= 0.5
    elif vehicle_type == "bike":
        base_fee *= 1.0
        hourly_rate *= 0.25
    elif vehicle_type == "truck":
        base_fee *= 3.0
        hourly_rate *= 1.0
    else:
        return "Invalid vehicle type"

    if hours <= 1:
        return base_fee
    return base_fee + (hours - 1) * hourly_rate

class Logger:
    def log_event(self, event_type, details):
        timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        log_entry = f"{timestamp} - {event_type}: {details}"
        print(log_entry)