package generate

func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)
	for i := 1; i <= numRows; i++ {
		item := make([]int, 0, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				item = append(item, 1)
			} else {
				item = append(item, res[i-2][j-1], res[i-2][j])
			}
		}
		res = append(res, item)
	}
	return res
}
