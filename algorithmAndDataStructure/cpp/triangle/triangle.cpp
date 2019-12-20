#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int minimumTotal(vector<vector<int>>& triangle) {
        int rows = triangle.size();
        vector<int> d(triangle[rows-1]);
        for (int i = rows-2; i >= 0; i--) {
            int cols = triangle[i].size();
            for (int j=0; j <cols; j++) {
                d[j] = min(d[j], d[j+1]) + triangle[i][j];
            }
        }
        return d[0];
    }
};