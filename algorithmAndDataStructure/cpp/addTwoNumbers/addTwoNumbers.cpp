#include <stdio.h>
#include <algorithm>
#include <vector>
#include <iostream>

using namespace std;

// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int x = 0 , y = 0, sum = 0, carry=0;
        ListNode *head=NULL, **tmp=&head;
        while(l1 != NULL || l2 != NULL || carry>0) {
            x = getValueAndNext(l1);
            y = getValueAndNext(l2);
            sum = x + y + carry;
            ListNode *node = new ListNode(sum%10);
            *tmp = node;
            tmp = &node->next;

            carry = sum/10;
        }
        return head;
    }
private:
    int getValueAndNext(ListNode* &node) {
        int value=0;
        if(node != NULL) {
            value = node->val;
            node = node->next;
        }
        return value;
    }
};