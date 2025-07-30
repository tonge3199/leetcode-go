package bitmanipulation

/*
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
题目大意 #
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。说明：解集不能包含重复的子集。

位运算思路：
- 数组长度为 n，则子集总数为 2^n（包括空集）。
- 可以把 0 到 (2^n - 1) 之间的每一个整数看作一个“掩码”（mask），
  掩码的第 i 位为 1 表示 nums[i] 在当前子集中，应当加入；否则跳过。
- 枚举所有掩码，就能得到所有子集。
*/

// 位运算
func subsets(nums []int) [][]int {
	n := len(nums)

	// The total number of subsets is 2^n (including the empty subset)
	totalSubsets := 1 << n

	result := make([][]int, 0, totalSubsets)

	// 0 -- 2^n -1 all masks
	for mask := 0; mask < totalSubsets; mask++ {
		var subset []int
		// 掩码每一位都检查是否 1
		// mask >> i & 1 取出第i位的值 (0 or 1)
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				// 如果第i位是1，说明nums[i]要加入当前子集
				subset = append(subset, nums[i])
			}
		}
		result = append(result, subset)
	}

	return result
}
