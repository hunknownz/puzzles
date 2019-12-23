#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
        int maxValue = 0, numsLen = nums.size();
        vector<int> dp(nums);
        for (int i = 0; i < numsLen; i++) {
            dp[i] = 1;
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j]) {
                    dp[i] = max(dp[i], dp[j]+1);
                }
            }
            maxValue = max(maxValue, dp[i]);
        }
        return maxValue;
    }

    int lengthOfLISII(vector<int>& nums) {
        int numsLen = nums.size();
        vector<int> lowEndings;
        for (int i = 0; i < numsLen; i++) {
            auto it = lower_bound(lowEndings.begin(), lowEndings.end(), nums[i]);
            if (it == lowEndings.end()) {
                lowEndings.push_back(nums[i]);
            } else {
                *it = nums[i];
            }
        }
        return lowEndings.size();
    }
};