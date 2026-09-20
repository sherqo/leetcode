#include <vector>
#include <algorithm>

class Solution {
public:
    int carFleet(int target, std::vector<int>& position, std::vector<int>& speed) {
        int n = position.size();
        if (n == 0) return 0;

        std::vector<std::pair<int, int>> cars(n);
        for (int i = 0; i < n; ++i) {
            cars[i] = {position[i], speed[i]};
        }

        std::sort(cars.begin(), cars.end(), [](const std::pair<int, int>& a, const std::pair<int, int>& b) {
            return a.first > b.first;
        });

        int fleets = 0;
        double maxTime = 0.0;

        for (int i = 0; i < n; ++i) {
            double curTime = (double)(target - cars[i].first) / cars[i].second;

            if (curTime > maxTime) {
                fleets++;
                maxTime = curTime;
            }
        }

        return fleets;
    }
};
