#include <iostream>
#include <vector>
#include <limits.h>

using namespace std;

int min(int x, int y, int z) {
    return min(x, min(y, z));
}

int max(int x, int y, int z) {
    return max(x, max(y, z));
}

class Solution {
public:
    int maxProduct(vector<int>& nums) {
        int prevMin = nums[0], prevMax = nums[0];
        int curMin = nums[0], curMax = nums[0], maxValue=nums[0];
        int numsLen = nums.size();
        for (int i=1; i<numsLen; i++) {
            curMin = min(prevMin*nums[i], prevMax*nums[i], nums[i]);
            curMax = max(prevMin*nums[i], prevMax*nums[i], nums[i]);

            prevMin = curMin, prevMax = curMax;
            maxValue = max(maxValue, curMax);
        }
        return maxValue;
    }
};