package triangle

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func minimumTotal(triangle [][]int) int {
	rows := len(triangle)
	d := make([]int, rows, rows)
	copy(d, triangle[rows-1])
	for i := rows - 2; i >= 0; i-- {
		cols := len(triangle[i])
		for j := 0; j < cols; j++ {
			d[j] = min(d[j], d[j+1]) + triangle[i][j]
		}
	}
	return d[0]
}
