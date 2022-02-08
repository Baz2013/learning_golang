package lcof

/**
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
	data := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
*/

func FindNumberIn2DArray(matrix [][]int, f int) bool {

	if len(matrix) == 0 {
		return false
	}

	m := len(matrix)    // 行数
	n := len(matrix[0]) // 列数

	for j, i := n-1, 0; j > -1 && i < m; {
		if matrix[i][j] == f {
			return true
		}

		if matrix[i][j] > f {
			j--
		}

		if matrix[i][j] < f {
			i++
		}
	}

	return false
}
