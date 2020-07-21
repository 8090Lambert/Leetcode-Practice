package _5_threeSum

import (
	"sort"
	"strconv"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	count := len(nums)
	uniqueMap := map[string]string{}
	res := make([][]int, 0)
	for index := 0; index < count; index++ {
		if index == 0 || nums[index] != nums[index-1] {
			left := index + 1
			right := count - 1
			for left < right {
				sum := nums[index] + nums[left] + nums[right]
				if sum == 0 {
					// 利用map去重
					uniqueFlag := strconv.Itoa(nums[index]) + strconv.Itoa(nums[left]) + strconv.Itoa(nums[right])
					if _, ok := uniqueMap[uniqueFlag]; !ok {
						uniqueMap[uniqueFlag] = uniqueFlag
						res = append(res, []int{nums[index], nums[left], nums[right]})
					}
					left++
					right--
				} else if sum > 0 {
					right--
				} else {
					left++
				}
			}
		}
	}

	return res
}

func threeSum1(nums []int) [][]int {
	count := len(nums)
	res := make([][]int, 0)
	uniqMap := map[string]struct{}{}
	for i := 0; i < count; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i+1
		right := count-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			uniqKey := strconv.Itoa(nums[i]) + strconv.Itoa(nums[left]) + strconv.Itoa(nums[right])
			if sum == 0 {
				if _, ok := uniqMap[uniqKey]; !ok {
					uniqMap[uniqKey] = struct{}{}
					res = append(res, []int{nums[i], nums[left], nums[right]})
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
