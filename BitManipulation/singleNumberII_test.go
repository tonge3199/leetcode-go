package bitmanipulation

import (
	"testing"
)

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
*/

func TestSingleNumberII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1: [2,2,3,2]",
			nums:     []int{2, 2, 3, 2},
			expected: 3,
		},
		{
			name:     "example 2: [0,1,0,1,0,1,99]",
			nums:     []int{0, 1, 0, 1, 0, 1, 99},
			expected: 99,
		},
		{
			name:     "single element",
			nums:     []int{42},
			expected: 42,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -1, -1, -2, -2, -2, 5},
			expected: 5,
		},
		{
			name:     "large numbers",
			nums:     []int{1000000, 999999, 1000000, 999999, 1000000, 999999, 123456},
			expected: 123456,
		},
		{
			name:     "zero as single number",
			nums:     []int{7, 7, 7, 0},
			expected: 0,
		},
	}

	// Test all three implementations
	implementations := []struct {
		name string
		fn   func([]int) int
	}{
		{"singleNumberII", singleNumberII},
		{"singleNumberII_v2", singleNumberII_v2},
		{"singleNumberII_general", singleNumberII_general},
	}

	for _, impl := range implementations {
		t.Run(impl.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					result := impl.fn(tt.nums)
					if result != tt.expected {
						t.Errorf("%s(%v) = %d, want %d", impl.name, tt.nums, result, tt.expected)
					}
				})
			}
		})
	}
}

// TestStateMachineVisualization 可视化三进制状态机的工作过程
func TestStateMachineVisualization(t *testing.T) {
	nums := []int{2, 2, 3, 2} // 期望结果是3

	t.Logf("=== 三进制状态机可视化 ===")
	t.Logf("输入数组: %v", nums)
	t.Logf("期望结果: 3")
	t.Logf("")

	ones, twos := 0, 0

	t.Logf("初始状态: ones=%08b, twos=%08b", ones, twos)
	t.Logf("")

	for i, num := range nums {
		oldOnes, oldTwos := ones, twos

		// 状态机更新
		ones = (ones ^ num) & ^twos
		twos = (twos ^ num) & ^ones

		t.Logf("步骤 %d: 处理数字 %d (%08b)", i+1, num, num)
		t.Logf("  之前: ones=%08b, twos=%08b", oldOnes, oldTwos)
		t.Logf("  之后: ones=%08b, twos=%08b", ones, twos)
		t.Logf("  ones=%d, twos=%d", ones, twos)
		t.Logf("")
	}

	t.Logf("最终结果: ones=%d (%08b)", ones, ones)

	if ones != 3 {
		t.Errorf("期望结果3，实际结果%d", ones)
	}
}

// TestBitLevelAnalysis 分析每一位的状态变化
func TestBitLevelAnalysis(t *testing.T) {
	nums := []int{2, 2, 3, 2} // 二进制: [10, 10, 11, 10]

	t.Logf("=== 位级分析 ===")
	t.Logf("数字 2 的二进制: %08b", 2)
	t.Logf("数字 3 的二进制: %08b", 3)
	t.Logf("")

	// 手动分析每一位
	t.Logf("位0 (最低位) 出现1的次数:")
	t.Logf("  2: 0, 2: 0, 3: 1, 2: 0 → 总计1次 → 1%%3=1 → 结果位0为1")
	t.Logf("")

	t.Logf("位1 出现1的次数:")
	t.Logf("  2: 1, 2: 1, 3: 1, 2: 1 → 总计4次 → 4%%3=1 → 结果位1为1")
	t.Logf("")

	t.Logf("所以最终结果: 位1为1, 位0为1 → 11(二进制) = 3(十进制)")

	result := singleNumberII(nums)
	if result != 3 {
		t.Errorf("期望结果3，实际结果%d", result)
	}
}

// TestGeneralApproachStepByStep 演示通用解法的逐步计算
func TestGeneralApproachStepByStep(t *testing.T) {
	nums := []int{2, 2, 3, 2}

	t.Logf("=== 通用解法逐步计算 ===")
	t.Logf("输入: %v", nums)

	result := 0

	for bit := 0; bit < 4; bit++ { // 只检查前4位就够了
		count := 0

		t.Logf("\n检查第%d位:", bit)
		for _, num := range nums {
			bitValue := (num >> bit) & 1
			count += bitValue
			t.Logf("  数字%d的第%d位: %d", num, bit, bitValue)
		}

		remainder := count % 3
		t.Logf("  第%d位总计: %d次, %d%%3 = %d", bit, count, count, remainder)

		if remainder != 0 {
			result |= (1 << bit)
			t.Logf("  → 结果的第%d位设为1", bit)
		} else {
			t.Logf("  → 结果的第%d位保持0", bit)
		}
	}

	t.Logf("\n最终结果: %d (二进制: %08b)", result, result)

	if result != 3 {
		t.Errorf("期望结果3，实际结果%d", result)
	}
}

// TestStateTransitions 测试状态转换表
func TestStateTransitions(t *testing.T) {
	t.Logf("=== 三进制状态转换表 ===")
	t.Logf("状态 (ones,twos) + 输入 → 新状态")
	t.Logf("(0,0) + 0 → (0,0)  // 0次+0 = 0次")
	t.Logf("(0,0) + 1 → (1,0)  // 0次+1 = 1次")
	t.Logf("(1,0) + 0 → (1,0)  // 1次+0 = 1次")
	t.Logf("(1,0) + 1 → (0,1)  // 1次+1 = 2次")
	t.Logf("(0,1) + 0 → (0,1)  // 2次+0 = 2次")
	t.Logf("(0,1) + 1 → (0,0)  // 2次+1 = 0次(模3)")

	// 验证状态转换
	testCases := []struct {
		ones, twos, input          int
		expectedOnes, expectedTwos int
	}{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 1},
		{0, 1, 0, 0, 1},
		{0, 1, 1, 0, 0},
	}

	for _, tc := range testCases {
		ones := (tc.ones ^ tc.input) & ^tc.twos
		twos := (tc.twos ^ tc.input) & ^ones

		if ones != tc.expectedOnes || twos != tc.expectedTwos {
			t.Errorf("状态转换错误: (%d,%d)+%d → (%d,%d), 期望(%d,%d)",
				tc.ones, tc.twos, tc.input, ones, twos, tc.expectedOnes, tc.expectedTwos)
		}
	}
}

func BenchmarkSingleNumberII(b *testing.B) {
	nums := []int{2, 2, 3, 2, 5, 5, 5, 7, 7, 7}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		singleNumberII(nums)
	}
}

func BenchmarkSingleNumberII_General(b *testing.B) {
	nums := []int{2, 2, 3, 2, 5, 5, 5, 7, 7, 7}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		singleNumberII_general(nums)
	}
}
