package main

import (
	"sort"
)

func myTwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res :=[]int{i, j,target}
				sort.Ints(res)
				return res
			}
		}
	}
	return nil
}

//输入：nums = [-4 -1 -1 0 1 2]
//输出：[[-1,-1,2],[-1,0,1]]
func threeSum(nums []int) [][]int {
	var res [][]int
	if nums == nil || len(nums) < 3 || nums[0] > 0 {
		return res
	}
	//m := &sync.Map{}

	sort.Ints(nums)
	for i := 2; i < len(nums); i++ {
		if nums[i] < 0 {
			continue
		}

	}
	return res
}

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}
