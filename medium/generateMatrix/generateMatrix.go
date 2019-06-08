package generateMatrix

func GenerateMatrix(n int) [][]int {
	res := [][]int{}
	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		res = append(res, tmp)
	}
	rowStart, colStart := 0, 0
	rowEnd, colEnd := n-1, n-1
	nums := 1
	
	for rowStart <= rowEnd && colStart <= colEnd {
		for c := colStart; c <= colEnd; c++ {
			res[rowStart][c] = nums
			nums += 1
		}
		for r := rowStart+1; r <= rowEnd; r++ {
			res[r][colEnd] = nums
			nums += 1
		}
		if rowStart < rowEnd && colStart < colEnd {
			for c := colEnd-1; c > colStart; c-- {
				res[rowEnd][c] = nums
				nums += 1
			}
			
			for r := rowEnd; r > rowStart; r-- {
				res[r][colStart] = nums
				nums += 1
			}
		}
		rowStart++
		rowEnd--
		colStart++
		colEnd--
	}
	
	return res
}
