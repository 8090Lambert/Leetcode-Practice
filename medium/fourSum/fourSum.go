package fourSum

import (
	"fmt"
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
					tmp := []int{nums[index], nums[k], nums[left], nums[right]}
					sort.Ints(tmp)
					uniqueFlag := strconv.Itoa(tmp[0]) + strconv.Itoa(tmp[1]) + strconv.Itoa(tmp[2]) + strconv.Itoa(tmp[3])
					fmt.Println(uniqueMap, tmp)
					if _, ok := uniqueMap[uniqueFlag]; !ok {
						res = append(res, []int{tmp[0], tmp[1], tmp[2], tmp[3]})
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
