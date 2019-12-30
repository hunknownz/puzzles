#include <iostream>
#include <vector>

using namespace std;

class UnionFindSet {
public:

    vector<int> p;
    int count;

    UnionFindSet(vector<vector<char> >& grid) {
        int rowSize = grid.size(), columnSize = grid[0].size();
        p = vector<int>(rowSize*columnSize);
        count = 0;
        for (int i = 0; i < rowSize; i++) {
            for (int j = 0; j < columnSize; j++) {
                if (grid[i][j] == '1') {
                    p[i*columnSize+j] = i*columnSize + j;
                    count = count + 1;
                }
            }
        }
    }

    int FindRecursion(int x) {
        return p[x] == x ? x : p[x] = FindRecursion(p[x]);
    }

    int FindLoop(int x) {
        int root = x;
        while (root != p[root]) {
            root = p[root];
        }
        while (x != p[x]) {
            int tmp  = p[x];
            p[x] = root;
            x = tmp;
        }
        return root;
    }

    void Union(int x, int y) {
        int rootX = FindRecursion(x),rootY = FindRecursion(y);
        if (rootX != rootY) {
            p[rootX] = rootY;
            count -= 1;
        }
    }
};

class Solution {
public:
    int numIslands(vector<vector<char> >& grid) {
        int rowSize = grid.size();
        if (rowSize == 0 || grid[0].size() == 0) {
            return 0;
        }
        int columnSize = grid[0].size();
        UnionFindSet *unionFindSet = new UnionFindSet(grid);
        int directions[2][2] = {{0, -1}, {-1, 0}};
        for (int i = 0; i < rowSize; i++) {
            for (int j = 0; j < columnSize; j++) {
                if (grid[i][j] == '0') {
                    continue;
                }
                for (int d = 0; d < 2; d++) {
                    int prevI = i + directions[d][0], prevJ = j + directions[d][1];
                    if (prevI >= 0 && prevJ >= 0 && grid[prevI][prevJ] == '1') {
                        unionFindSet->Union(prevI*columnSize+prevJ, i*columnSize+j);
                    }
                }
            }
        }
        return unionFindSet->count;
    }
};