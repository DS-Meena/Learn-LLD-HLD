#include <bits/stdc++.h>
#include "elevator_system.cpp"

using namespace std;

int main() {
    ElevatorSystem system(3);

    // 3 persons want to take elevator
    // Initial requests
    system.requestElevator(5);
    for (int i=0; i<5; i++) {
        system.step();
        system.printElevatorPositions();
    }
   
    system.requestElevator(1);
    system.step();
    system.printElevatorPositions();

    system.requestElevator(8);
    system.step();
    system.printElevatorPositions();


    for (int i=0; i<7; i++) {
        system.step();
        system.printElevatorPositions();
    }

    return 0;
}