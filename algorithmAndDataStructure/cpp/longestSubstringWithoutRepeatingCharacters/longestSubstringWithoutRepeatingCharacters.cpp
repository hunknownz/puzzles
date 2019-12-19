#include <stdio.h>
#include <iostream>
#include <memory>

using namespace std;

class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        const int MAX_CHARS = 256;
        int alphabet[MAX_CHARS];
        memset(alphabet, -1, sizeof(alphabet));
        int lastPos = -1, max = 0;
        for (int i = 0; i < s.size(); i++) {
            int value = s[i];
            if (alphabet[value] != -1 && lastPos < alphabet[value]) {
                lastPos = alphabet[value];
            }
            alphabet[value] = i;
            if (i - lastPos > max) {
                max = i - lastPos;
            }
        }
        return max; 
    }
};