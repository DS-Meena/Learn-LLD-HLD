#include <bits/stdc++.h>
#include "elevator.cpp"

using namespace std;

class ElevatorSystem{
    private:
        vector<Elevator> elevators;
    public:
        ElevatorSystem(int numElevators){
            for(int i=0;i<numElevators;i++){
                elevators.push_back(Elevator(0));
            }
        }

        void requestElevator(int floor) {
            int mn_idx = 0;
            int minDistance = abs(elevators[mn_idx].getCurrentFloor() - floor);

            for (int i=1; i<elevators.size(); i++) {
                int distance = abs(elevators[i].getCurrentFloor() - floor);
                if (distance < minDistance) {
                    minDistance = distance;
                    mn_idx = i;
                }
            }

            // add request to elevator, that is closest
            elevators[mn_idx].addRequest(Request{floor, floor});
        }

        void step() {
            for (auto& elevator : elevators) {
                // move according to current request
                elevator.move();

                // check if need to change direction
                elevator.processNextRequest();
            }
        }

        void printElevatorPositions() {
            cout << "Elevator positions: ";
            for (auto& elevator : elevators) {
                cout << elevator.getCurrentFloor() << " ";
            }
            cout << endl;
        }
};