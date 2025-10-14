package strings

func ZeroMatrix(matrix [][]int) {
	var rows []int
	var cols []int
	for i, arr := range matrix {
		for j := range arr {
			if matrix[i][j] == 0 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}
	// Now we need to stablish that rows and columns to zero
	setRowsToZero(rows, matrix)
	setColsToZero(cols, matrix)
}

func setRowsToZero(rows []int, matrix [][]int) {
	for i := range rows {
		// len(matris[i]) give the number of cols, then for each col in a row i, loop the cols...
		for j := range len(matrix[i]) {
			matrix[rows[i]][j] = 0
		}
	}
}

func setColsToZero(cols []int, matrix [][]int) {
	for i := range cols {
		// len(matrix) give the number of rows, then for each col, loop the rows...
		for j := range len(matrix) {
			matrix[j][cols[i]] = 0
		}
	}
}
