#include <stdio.h>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int total = nums1.size() + nums2.size();
        if (total&1) {
            return findKthValue(nums1, nums2, total/2 + 1);
        } else {
            return (findKthValue(nums1, nums2, total/2) + findKthValue(nums1, nums2, total/2+1))/2;
        }
    }
private:
    double findKthValue(vector<int> nums1, vector<int> nums2, int k) {
        if (nums1.empty()) {
            return nums2[k-1];
        }
        if (nums2.empty()) {
            return nums1[k-1];
        }
        if (k == 1) {
            return min(nums1[0], nums2[0]);
        }
        int mid1 = min(int(nums1.size()), k/2), mid2 = min(int(nums2.size()), k/2);
        if (nums1[mid1-1] < nums2[mid2-1]) {
            return findKthValue(vector<int>(nums1.begin()+mid1, nums1.end()), nums2, k-mid1);
        } else {
            return findKthValue(nums1, vector<int>(nums2.begin()+mid2, nums2.end()), k-mid2);
        }
    }
};
