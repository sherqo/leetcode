class Solution {
public:
    bool isUpper(char c) {
        if (c >= 'A' && c <= 'Z') {
            return true;
        }
        return false;
    }
    bool isLower(char c) {
        if (c >= 'a' && c <= 'z') {
            return true;
        }
        return false;
    }

    bool isPalindrome(string s) {
        int n = s.length(); int l = 0, r = n - 1;

        while (l < n && r >= l) {
            if (isalnum(s[l]) && isalnum(s[r])) {
                if (std::tolower(s[l]) == std::tolower(s[r])) {
                    l++;
                    r--;
                    continue;
                } else {
                    return false;
                }
            } else {
                if (isUpper(s[l]))
                    s[l] = s[l] + 32;
                if (isUpper(s[r]))
                    s[r] = s[r] + 32;

                if (!(isUpper(s[l]) || isLower(s[l]) || std::isdigit(s[l])))
                    l++;
                if (!(isUpper(s[r]) || isLower(s[r]) || std::isdigit(s[r])))
                    r--;
            }
        }

        return true;
    }
};
