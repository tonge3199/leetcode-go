package array

import "sort"

// Given an array nums of n integers, are there elements a, b, c
// in nums such that a + b + c = 0?
// Find all unique triplets in the array which gives the sum of zero.

/*
Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

Try demo:

  For each element, treat it as first number (a)
  two ptrs find a pair of number sum to -a.

  Handle Duplicates: TODO

*/

// double ptr + sort
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	length := len(nums)

	for i := 0; i < len(nums)-2; i++ {
		// skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// two-pointer
		// demo: nums[i] < start < end
		start := i + 1
		end := length - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum == 0 { // Found a valid triplet
				result = append(result, []int{nums[i], nums[start], nums[end]})
				// Skip duplicates for start
				for start < end && nums[start] == nums[start+1] {
					start++
				}
				// Skip duplicates for end
				for start < end && nums[end] == nums[end-1] {
					end--
				}
                                // Move ptr after skipping duplicates , Fixed
				start++
				end--
			} else if sum < 0 {
				start++
			} else {
				end--
			}
		}
	}
	return result
}

//
