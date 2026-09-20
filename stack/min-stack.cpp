class MinStack {
private:
    std::vector<int> st;
    std::vector<int> mn;

public:
    MinStack() {
    }
    
    void push(int value) {
        st.push_back(value);
        if (mn.empty()) {
            mn.push_back(value);
        } else {
            mn.push_back(std::min(value, mn.back()));
        }
    }
    
    void pop() {
        st.pop_back();
        mn.pop_back();
    }
    
    int top() {
        return st.back();
    }
    
    int getMin() {
        return mn.back();
    }
};
