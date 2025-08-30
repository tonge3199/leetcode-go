package main

import (
	"fmt"
)

// 游戏词条有正也有负，需要从一组词条里选择其中连续的一段，词条数量任意，选取后可在这部分词条里删掉一条或不删除，但是删除后，词条数量不能为0，总属性越大越强。
// eg input :
// -2 -1 3 -2 3
// output :
// 6
// explain : chosen last three and delete -2
// eg input :
// -2 -1 3 4 5 -3
// output :
// 12
// chosen middles 3 4 5 , do not delete.
// 我的思路：滑动窗口 + 其他额外方法
func main() {
	num := 0
	records := make([]int, 0)
	for {
		n, err := fmt.Scan(&num)
		if n == 0 || err != nil {
			break
		} else {
			records = append(records, num)
		}
	}

	result := maxSubarraySum(records)
	myresult := maxSubarraySum2(records)
	fmt.Printf("my answer %d\n", myresult)
	fmt.Printf("answer %d\n", result)
}

func maxSubarraySum(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0] // 只有一个元素，必须选它
	}

	// rightMax[i] 表示以位置i结尾的最大子数组和
	// 这是经典的Kadane算法的dp数组
	rightMax := make([]int, n)
	rightMax[0] = nums[0]
	for i := 1; i < n; i++ {
		// 要么单独开始一个新的子数组，要么延续前面的子数组
		rightMax[i] = max(nums[i], rightMax[i-1]+nums[i])
	}

	// leftMax[i] 表示从位置i开始的最大子数组和
	// 这是Kadane算法的反向版本
	leftMax := make([]int, n)
	leftMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		// 要么单独开始一个新的子数组，要么延续后面的子数组
		leftMax[i] = max(nums[i], leftMax[i+1]+nums[i])
	}

	// 情况1：不删除任何元素，就是经典的最大子数组和问题
	result := rightMax[0]
	for i := 1; i < n; i++ {
		result = max(result, rightMax[i])
	}

	// 情况2：删除一个元素
	// 枚举每个可能删除的位置
	for i := 0; i < n; i++ {
		var candidate int

		if i == 0 {
			// 删除第一个元素，只考虑从位置1开始的部分
			if n > 1 {
				candidate = leftMax[1]
			}
		} else if i == n-1 {
			// 删除最后一个元素，只考虑到位置n-2结束的部分
			candidate = rightMax[n-2]
		} else {
			// 删除中间元素，左右两部分可以拼接
			// 注意：这里要求左右两部分都必须非空
			// rightMax[i-1]是以i-1结尾的最大子数组和
			// leftMax[i+1]是从i+1开始的最大子数组和
			candidate = rightMax[i-1] + leftMax[i+1]
		}

		result = max(result, candidate)
	}

	return result
}

func maxSubarraySum2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0] // 只有一个元素，必须选它
	}

	// 修复Bug1: 完整计算所有DP值
	// leftMax[i] 表示以位置i结尾的最大子数组和
	leftMax := make([]int, n)
	leftMax[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1]+nums[i], nums[i])
	}

	// rightMax[i] 表示从位置i开始的最大子数组和
	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1]+nums[i], nums[i])
	}

	// 修复Bug3: 先考虑不删除元素的情况
	maxSum := leftMax[0]
	for i := 1; i < n; i++ {
		maxSum = max(maxSum, leftMax[i])
	}

	// 修复Bug2: 正确处理删除元素的逻辑
	// 注意：删除边界元素的情况已经被"不删除"情况覆盖，只需考虑删除中间元素
	for i := 1; i < n-1; i++ {
		// 删除中间元素，左右拼接
		candidate := leftMax[i-1] + rightMax[i+1]
		maxSum = max(maxSum, candidate)
	}

	return maxSum
}

func arrSum(arr []int) (sum int) {
	for _, e := range arr {
		sum += e
	}
	return
}
