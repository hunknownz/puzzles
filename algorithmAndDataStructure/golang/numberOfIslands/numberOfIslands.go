package numberOfIslands

type UnionFindSet struct {
	p []int
	count int
}

func (ufs *UnionFindSet) findRecursion(x int) int {
	if ufs.p[x] == x {
		return x
	} else {
		ufs.p[x] = ufs.findRecursion(ufs.p[x])
		return ufs.p[x]
	}
}

func (ufs *UnionFindSet) findLoop(x int) (root int) {
	root = x
	for ufs.p[root] != root {
		root = ufs.p[root]
	}

	for ufs.p[x] != x {
        tmp := ufs.p[x]
		ufs.p[x] = root
		x = ufs.p[tmp]
	}
    
    return root
}

func (ufs *UnionFindSet) union(x, y int) {
	rootX, rootY := ufs.findRecursion(x), ufs.findRecursion(y)
	if rootX != rootY {
		ufs.p[rootX] = rootY
		ufs.count -= 1
	}
}

func numIslands(grid [][]byte) int {
	rowLen := len(grid)
	if rowLen == 0 || len(grid[0]) == 0 {
		return 0
	}
	columnLen := len(grid[0])
	unionFindSet := new(UnionFindSet)
	directions := [][]int {
        {0, -1},
        {-1, 0},
    }
    unionFindSet.p = make([]int, rowLen*columnLen, rowLen*columnLen)
	for i, row := range grid {
		for j, item := range row {
			if item == '0' {
				continue
			}
			unionFindSet.p[i*columnLen+j] = i * columnLen + j
			unionFindSet.count += 1
			for _, d := range directions {
				prevI, prevJ := i + d[0], j + d[1]
				if prevI >=0 && prevJ >=0 && grid[prevI][prevJ] == '1' {
					unionFindSet.union(i*columnLen+j, prevI*columnLen+prevJ)
				}
			}
		}
	}

	return unionFindSet.count
}