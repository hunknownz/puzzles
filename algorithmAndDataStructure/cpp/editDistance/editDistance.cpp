#include <iostream>
#include <vector>

using namespace std;

int min(int x, int y, int z) {
    return min(x, min(y, z));
}

class Solution {
public:
    int minDistance(string word1, string word2) {
        int word1Len = word1.size(), word2Len = word2.size();
        if (word1Len == 0) {
            return word2Len;
        }
        if (word2Len == 0) {
            return word1Len;
        }
        vector<int> prevDp(word1Len+1, 0), curDp(word1Len+1, 0);
        for (int i = 0; i <= word1Len; i++) {
            prevDp[i] = i;
        }

        for (int i = 1; i <= word2Len; i++) {
            curDp[0] = i;
            for (int j = 1; j <= word1Len; j++) {
                if (word2[i-1] == word1[j-1]) {
                    curDp[j] = prevDp[j-1];
                } else {
                    curDp[j] = min(prevDp[j-1], prevDp[j], curDp[j-1]) + 1;
                }
            }
            prevDp = curDp;
        }

        return curDp[word1Len];
    }
};