package bitmanipulation

/*
Given a non-empty array of integers, every element appears three times except for one, which appears exactly once. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,3,2]
Output: 3
Example 2:

Input: [0,1,0,1,0,1,99]
Output: 99

题目大意 #
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。
找出那个只出现了一次的元素。要求算法时间复杂度是线性的，并且不使用额外的辅助空间。

解题思路 #
这一题是第 136 题的加强版。这类题也可以扩展，在数组中每个元素都出现 5 次，找出只出现 1 次的数。

本题中要求找出只出现 1 次的数，出现 3 次的数都要被消除。第 136 题是消除出现 2 次的数。
这一题也会相当相同的解法，出现 3 次的数也要被消除。

核心思想：模拟三进制状态机
- 定义状态：00、01、10，这 3 个状态分别表示某位上1的个数 % 3 的结果
- 当一个数出现 3 次，那么它每个位置上的 1 出现的次数肯定是 3 的倍数
- 所以当 1 出现 3 次以后，就归零清除

状态转换：
- 00 + 1 → 01 (0次变1次)
- 01 + 1 → 10 (1次变2次)
- 10 + 1 → 00 (2次变0次，完成3次循环)

使用两个变量 ones 和 twos 来表示状态：
- ones: 记录每个位上出现1的次数 % 3 == 1 的位
- twos: 记录每个位上出现1的次数 % 3 == 2 的位
*/

// 解法一：经典的三进制状态机
func singleNumberII(nums []int) int {
	ones, twos := 0, 0

	for _, num := range nums {
		// 更新 ones: 当前位在 ones 中为1但在 twos 中为1时，说明出现了1次
		ones = (ones ^ num) & ^twos
		// 更新 twos: 当前位在 twos 中为1但在 ones 中为0时，说明出现了2次
		twos = (twos ^ num) & ^ones
	}

	return ones
}

// 解法二：另一种三进制状态机实现 (修复版)
func singleNumberII_v2(nums []int) int {
	ones, twos := 0, 0

	for _, num := range nums {
		// 计算新的 twos：当前 ones 中的位遇到相同的位时进位到 twos
		twos = twos ^ (ones & num)

		// 更新 ones
		ones = ones ^ num

		// 清除出现3次的位（同时在 ones 和 twos 中为1的位）
		threes := ones & twos
		ones &= ^threes
		twos &= ^threes
	}

	return ones
}

// 解法三：通用解法 - 可以扩展到任意k次
func singleNumberII_general(nums []int) int {
	result := 0

	// 对每一位进行处理
	for i := 0; i < 32; i++ {
		count := 0

		// 统计第i位上1的个数
		for _, num := range nums {
			if (num>>i)&1 == 1 {
				count++
			}
		}

		// 如果count % 3 != 0，说明单独出现的数在第i位为1
		if count%3 != 0 {
			result |= (1 << i)
		}
	}

	return result
}
