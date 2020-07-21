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


func Four (nums []int, target int) [][]int {
	sort.Ints(nums)
	count := len(nums)
	res := make([][]int, 0)
	uniqueMap := map[string]struct{}{}
	for i := 0; i < count; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		newTarget := target - nums[i]
		for j := i+1; j < count; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j+1
			right := count-1
			for left < right {
				sum := nums[j] + nums[left] + nums[right]
				if sum == newTarget {
					key := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[left]) + strconv.Itoa(nums[right])
					if _, ok := uniqueMap[key]; !ok {
						res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
						uniqueMap[key] = struct{}{}
					}
					left++
					right--
				} else if sum < newTarget {
					left++
				} else {
					right--
				}
			}
		}
	}

	return res
}