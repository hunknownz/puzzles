#include <iostream>

using namespace std;

class Solution {
public:
    int climbStairs(int n) {
        int lastOneStepWays = 1, lastTwoStepWays = 0;
        for (int i = 1; i < n; i++) {
            int ways = lastTwoStepWays + lastOneStepWays;
            lastTwoStepWays = lastOneStepWays;
            lastOneStepWays = ways;
        }
        return lastOneStepWays + lastTwoStepWays;
    }
};