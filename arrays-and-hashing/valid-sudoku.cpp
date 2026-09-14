class Solution {
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        // for rows
        for (auto i = 0; i < 9; ++i) {
            auto row = board[i];
            std::unordered_set<int> st;

            for (auto j = 0; j < 9; j++) {
                if (row[j] == '.')
                    continue;
                if (st.contains(row[j]))
                    return false;
                st.insert(row[j]);
            }
        }

        // for cols
        for (auto i = 0; i < 9; ++i) {
            std::unordered_set<int> st;

            for (auto j = 0; j < 9; j++) {
                if (board[j][i] == '.')
                    continue;
                if (st.contains(board[j][i]))
                    return false;
                st.insert(board[j][i]);
            }
        }

        // for boxes
        for (auto i = 0; i < 9; ++i) {
            std::unordered_set<int> st;
            for (auto j = 0; j < 9; ++j) {
                int r = 3 * (i / 3) + (j / 3);
                int c = 3 * (i % 3) + (j % 3);
                if (board[r][c] == '.')
                    continue;

                if (st.contains(board[r][c]))
                    return false;
                st.insert(board[r][c]);
            }
        }

        return true;
    }
};
