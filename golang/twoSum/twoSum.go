package twoSum

func TwoSum( numbers []int, target int) []int {
	res := make([]int, 0)
	sumMap := map[int]int{}
	for index := 0; index < len(numbers); index++ {
		need := target - numbers[index]
		if _, ok := sumMap[need]; ok {
			res = append(res, index+1, sumMap[need])
			break;
		}
		sumMap[numbers[index]] = index+1
	}

	return res
}