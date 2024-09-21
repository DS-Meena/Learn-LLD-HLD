#include <bits/stdc++.h>

using namespace std;

enum class Direction {UP = 1, DOWN = -1, IDLE = 0};
enum class State {MOVING=1, STOPPED=0};

struct Request {
    int sourceFloor;
    int destinationFloor;
};

struct compareRequest {
    bool operator()(pair<Request, int> a, pair<Request, int> b) {
        return a.second > b.second;
    }
};

class Elevator {
    private:
        int currentFloor;
        Direction direction;
        State state;
        priority_queue<pair<Request, int>, vector<pair<Request, int>>, compareRequest> requests;
        int maxCapacity;

    public:
        Elevator(int initialfloor) {
            currentFloor = initialfloor;
            direction = Direction::IDLE;
            state = State::STOPPED;
            maxCapacity = 10;
        }

        void move() {
            if (state == State::MOVING) {
                currentFloor += static_cast<int>(direction);
            }
        }

        void addRequest(const Request& request) {
            if (requests.size() < maxCapacity) {
                // first finish the closest request
                requests.push({request, abs(currentFloor - request.sourceFloor)});
                updateDirection();
            }
        }

        void processNextRequest() {
            if(!requests.empty()) {

                auto nextRequest = requests.top().first;

                // after finishing current request
                if (currentFloor == nextRequest.destinationFloor) {
                    requests.pop();
                    state = State::STOPPED;
                    updateDirection();   // update direction acc to closest call
                } else {

                    // keep moving
                    state = State::MOVING;
                }
            } else {
                // if no request, then stop
                direction = Direction::IDLE;
                state = State::STOPPED;
            }
        } 

        // change direction to complete closest request
        void updateDirection() {
            if (!requests.empty()) {
                auto nextRequest = requests.top().first;
                
                if (nextRequest.destinationFloor > currentFloor) {
                    direction = Direction::UP;
                } else if (nextRequest.destinationFloor < currentFloor) {
                    direction = Direction::DOWN;
                }
            }
        }

        int getCurrentFloor() {
            return currentFloor;
        }
};