#include <iostream>
#include <vector>
#include <limits.h>

using namespace std;

class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
        vector<int> dp(amount+1, -1);
        dp[0] = 0;

        int coinsNum = coins.size();
        for (int i = 1; i <= amount; i++) {
            for (int j = 0; j < coinsNum; j++) {
                if (i < coins[j]) {
                    continue;
                }
                int tmp = dp[i-coins[j]];
                if (tmp == -1) {
                    continue;
                }
                if (dp[i] > tmp+1 || dp[i] == -1) {
                    dp[i] = tmp + 1;
                }
            }
        }

        return dp[amount];
    }
};