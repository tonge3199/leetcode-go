package main

import "fmt"

/*
题目描述：

老张爱好爬山。不过老张认为过频繁的爬山对膝盖不太好。

老张给自己定了一个规则，原则上只能每隔一天爬山一次，如果今天爬山了，那么明天就休息一天不爬山了。但老张认为凡事都有例外，所以他给了自己k次机会，在昨天已经爬山的情况下，今天仍然连续爬山！换句话说就是老张每天最多爬山一次，原则上如果昨天爬山了那么今天就不爬山，但有最多k次机会打破这一原则。

爬山让人心情愉悦，所以老张每天爬山都能获得一定的愉悦值，请帮老张规划一下爬山计划来获得最大的愉悦值之和。

输入描述

第一行两个整数n和k，表示老张正在计划未来n天的爬山计划以及k次打破原则的机会。

第二行n个整数a_{1},a_{2},...,a_{n}，其中a_{i}表示接下来第i天如果进行爬山可以获得的愉悦值。

1≤n≤2000，1≤k≤1000，1≤a≤10000

输出描述

输出一行一个数，表示老张能在最佳爬山计划下获得的愉悦值之和。

样例输入

7 1

1 2 3 4 5 6 7

样例输出

19

提示

样例解释

最优的方案是选择选择第2、4、6天爬山，并在第7天打破一次原则一次（因为第6天已经爬过了，原则上不能继续爬山，需要使用一次打破原则的机会）。
*/

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	result := myHappy(arr, k)
	fmt.Println(result)
}

// happy 使用动态规划解决爬山规划问题
// arr: 每天爬山的愉悦值数组
// k: 可以打破原则的次数
func happy(arr []int, k int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	// dp[i][j][state] 表示：
	// 前i天，使用了j次打破原则的机会，第i天状态为state时的最大愉悦值
	// state: 0表示第i天不爬山，1表示第i天爬山
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, 2)
			// 初始化为-1表示不可达状态
			dp[i][j][0] = -1
			dp[i][j][1] = -1
		}
	}

	// 边界条件：第0天（初始状态）
	dp[0][0][0] = 0 // 第0天不爬山，使用0次机会，愉悦值为0

	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			// 第i天不爬山的情况
			// 可以从前一天的任何状态转移而来
			if dp[i-1][j][0] != -1 {
				dp[i][j][0] = max(dp[i][j][0], dp[i-1][j][0])
			}
			if dp[i-1][j][1] != -1 {
				dp[i][j][0] = max(dp[i][j][0], dp[i-1][j][1])
			}

			// 第i天爬山的情况
			// 从前一天不爬山的状态转移（正常情况）
			if dp[i-1][j][0] != -1 {
				dp[i][j][1] = max(dp[i][j][1], dp[i-1][j][0]+arr[i-1])
			}

			// 从前一天爬山的状态转移（需要使用一次打破原则的机会）
			if j > 0 && dp[i-1][j-1][1] != -1 {
				dp[i][j][1] = max(dp[i][j][1], dp[i-1][j-1][1]+arr[i-1])
			}
		}
	}

	// 寻找最终答案：第n天的所有可能状态中的最大值
	maxHappiness := 0
	for j := 0; j <= k; j++ {
		if dp[n][j][0] != -1 {
			maxHappiness = max(maxHappiness, dp[n][j][0])
		}
		if dp[n][j][1] != -1 {
			maxHappiness = max(maxHappiness, dp[n][j][1])
		}
	}

	return maxHappiness
}

// max 返回两个整数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func myHappy(arr []int, k int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	
	// ✅ 修复：确保k+1的大小，处理k=0的情况
	dp := make([][][]int, n)

	// init
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)  // ✅ 修复：k+1而不是k
		for j := 0; j <= k; j++ {   // ✅ 修复：j <= k
			dp[i][j] = make([]int, 2)
			dp[i][j][0] = -1
			dp[i][j][1] = -1
		}
	}
	
	// ✅ 修复：正确的初始化
	dp[0][0][0] = 0  // 第一天不爬山
	dp[0][0][1] = arr[0]  // 第一天爬山

	// state transfer
	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {  // ✅ 修复：j <= k
			// 第i天不爬山：可以从前一天任何状态转移
			if dp[i-1][j][0] != -1 {
				dp[i][j][0] = max(dp[i][j][0], dp[i-1][j][0])
			}
			if dp[i-1][j][1] != -1 {
				dp[i][j][0] = max(dp[i][j][0], dp[i-1][j][1])
			}
			
			// 第i天爬山：正常情况（前一天不爬山）
			if dp[i-1][j][0] != -1 {
				dp[i][j][1] = max(dp[i][j][1], dp[i-1][j][0] + arr[i])
			}
			
			// 第i天爬山：打破规则（前一天爬山，需要使用1次机会）
			if j > 0 && dp[i-1][j-1][1] != -1 {
				dp[i][j][1] = max(dp[i][j][1], dp[i-1][j-1][1] + arr[i])
			}
		}
	}

	// ✅ 修复：只查找最后一天的状态
	maxHappiness := 0
	for j := 0; j <= k; j++ {
		if dp[n-1][j][0] != -1 {
			maxHappiness = max(maxHappiness, dp[n-1][j][0])
		}
		if dp[n-1][j][1] != -1 {
			maxHappiness = max(maxHappiness, dp[n-1][j][1])
		}
	}

	return maxHappiness
}
