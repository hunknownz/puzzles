#include <iostream>
#include <limits.h>
#include <vector>

using namespace std;

class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int minBuy= INT_MAX, maxProfit= 0;
        int pricesLen = prices.size();
        for (int i=0; i < pricesLen; i++) {
            int price = prices[i];
            maxProfit = price>minBuy?max(maxProfit, price-minBuy):maxProfit;
            minBuy = min(minBuy, price);
        }
    }
};