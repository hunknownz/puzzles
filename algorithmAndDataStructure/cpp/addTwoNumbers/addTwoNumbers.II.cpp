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
    ListNode *addTwoNumbers(ListNode *l1, ListNode *l2) {
        int carry=0;
        ListNode *p=NULL,*res=NULL;
        while(l1!=NULL || l2!=NULL || carry!=0){
            int sum=carry;
            
            if(l1!=NULL) sum+=l1->val,l1=l1->next;
            if(l2!=NULL) sum+=l2->val,l2=l2->next;
            
            if(res==NULL){
                res=new ListNode(sum%10);
                p=res;
            }
            else{
                ListNode *tmp=new ListNode(sum%10);
                p->next=tmp;
                p=p->next;
            }
            carry=sum/10;
        }
        return res;
    }
};