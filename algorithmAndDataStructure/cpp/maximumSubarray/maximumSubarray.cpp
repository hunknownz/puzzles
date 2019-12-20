#include <iostream>
#include <vector>
#include <limits.h>

using namespace std;

class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int sum = 0, maxValue = INT_MIN, numsLen = nums.size();
        for (int i = 0; i < numsLen; i++) {
            sum += nums[i];
            maxValue = max(maxValue, sum);
            if (sum < 0) {
                sum = 0;
            }
        }
        return maxValue;
    }
};