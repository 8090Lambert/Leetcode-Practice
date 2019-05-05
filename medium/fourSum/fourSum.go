package fourSum

import (
	"sort"
	"strconv"
)

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	count := len(nums)
	uniqueMap := map[string]string{}
	res := make([][]int, 0)

	for index := 0; index < count; index++ {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		newtarget := target - nums[index]
		for k := index + 1; k < count; k++ {
			if k > index+1 && nums[k] == nums[k-1] {
				continue
			}
			left := k + 1
			right := count - 1
			for left < right {
				sum := nums[k] + nums[left] + nums[right]
				if sum == newtarget {
					uniqueFlag := strconv.Itoa(nums[index]) + strconv.Itoa(nums[k]) + strconv.Itoa(nums[left]) + strconv.Itoa(nums[right])
					if _, ok := uniqueMap[uniqueFlag]; !ok {
						res = append(res, []int{nums[index], nums[k], nums[left], nums[right]})
						uniqueMap[uniqueFlag] = uniqueFlag
					}
					left++
					right--
				} else if sum < newtarget {
					left++
				} else {
					right--
				}
			}
		}
	}

	return res
}
