package array

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it is able to trap after raining.

eg. elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!
*/

// Trap 双指针法解决接雨水问题
// 核心思想：对于任意位置，能接水的高度取决于左右两边最高柱子的最小值
// 时间复杂度: O(n), 空间复杂度: O(1)
func Trap(height []int) int {
	// 边界情况：少于3个柱子无法形成积水
	if len(height) <= 2 {
		return 0
	}

	left, right := 0, len(height)-1
	rightMax, leftMax := 0, 0
	// 累计积水量
	water := 0

	// 双指针向中间移动，直到相遇
	for left < right {
		// 关键判断：选择较小的一边进行处理
		// 这样可以确保当前位置的积水高度由较小一边的最大值决定
		if height[left] < height[right] {
			// 处理左边指针
			if height[left] >= rightMax {
				// 当前位置是新的最高点，更新rightMax
				rightMax = height[left]
			} else {
				// 当前位置可以积水，积水高度 = 左边最大高度 - 当前高度
				water += rightMax - height[left]
			}
			left++ // 左指针向右移动
		} else {
			// 处理右边指针
			if height[right] >= leftMax {
				// 当前位置是新的最高点，更新leftMax
				leftMax = height[right]
			} else {
				// 当前位置可以积水，积水高度 = 右边最大高度 - 当前高度
				water += leftMax - height[right]
			}
			right-- // 右指针向左移动
		}
	}

	return water
}

func Trap2(height []int) int {
	res := 0
	L, R := 0, len(height)-1
	leftMax, rightMax := 0, 0
	var currentH int
	var Area int
	// current position using small height

	for L < R {
		if height[L] <= height[R] {
			currentH = height[L]
			Area = leftMax - currentH
			if Area > 0 {
				res += Area
			}
			if leftMax < currentH {
				leftMax = currentH
			}
			L++
		} else {
			currentH = height[R]
			Area = rightMax - currentH
			if Area > 0 {
				res += Area
			}
			if rightMax < currentH {
				rightMax = currentH
			}
			R--
		}
	}
	return res
}
