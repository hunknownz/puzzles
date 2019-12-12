#include <stdio.h>
#include <algorithm>
#include <vector>
#include <iostream>

using namespace std;

class Node {
public:
    int val, pos;

    Node(){
        val = pos = 0;
    }

    Node(int v, int p){
        val=v, pos=p;
    }

    bool operator<(const Node& other) const {
        return val < other.val;
    }

};

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<Node> nodes;
        int len = nums.size();
        for (int i=0; i < len; ++i) {
            nodes.push_back(Node(nums[i], i));
        }
        sort(nodes.begin(), nodes.end());
        int l=0, r = len - 1;
        vector<int> res;
        while(l < r) {
            int sum = nodes[l].val + nodes[r].val;
            if (sum == target) {
                res.push_back(nodes[l].pos);
                res.push_back(nodes[r].pos);
                break;
            } else if (sum < target) {
                ++l;
            } else {
                --r;
            }
        }
        return res;
    }
};