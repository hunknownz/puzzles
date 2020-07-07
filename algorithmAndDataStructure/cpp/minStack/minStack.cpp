#include <iostream>
#include <stack>
#include <limits.h>

using namespace std;

class MinStack {
public:
    /** initialize your data structure here. */
    stack<int> values;
    int minValue;
    MinStack() {
        minValue = INT_MAX;
    }
    
    void push(int x) {
        if (x <= minValue) {
            values.push(minValue);
            minValue = x;
        }
        values.push(x);
    }
    
    void pop() {
        int value = values.top();
        values.pop();
        if (value == minValue) {
            minValue = values.top();
            values.pop();
        }
    }
    
    int top() {
        return values.top();
    }
    
    int getMin() {
        return minValue;
    }
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(x);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->getMin();
 */