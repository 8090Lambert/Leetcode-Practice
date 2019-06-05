package permute

import "fmt"

//var res [][]int

func Permute(nums []int) [][]int {
	res := [][]int{}
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	// 方法一
	backTrack1(nums, []int{}, &res)

	// 方法二
	//backTrack2(nums, &res, 0)
	fmt.Println(res)
	return res
}

func backTrack1 (left, ans []int, res *[][]int) {
	if len(left) == 0 {
		tmp := append([]int{}, ans...)
		*res = append(*res, tmp)
		return
	}

	for key, value := range left {
		cleft := append([]int{}, left...)
		cans := append([]int{}, ans...)

		backTrack1(append(cleft[:key], cleft[key+1:]...), append(cans, value), res)
	}
}

func backTrack2 (nums []int, res *[][]int, start int) {
	if len(nums) == start {
		item := append([]int{}, nums...)
		*res = append(*res, item)
		return
	}

	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		backTrack2(nums, res, start+1)
		nums[start], nums[i] = nums[i], nums[start]
	}
}
