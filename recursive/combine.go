package recursive

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。

示例 1：

输入：n = 4, k = 2
输出：
[

	[2,4],
	[3,4],
	[2,3],
	[1,2],
	[1,3],
	[1,4],

]
示例 2：

输入：n = 1, k = 1
输出：[[1]]

提示：

1 <= n <= 20
1 <= k <= n
*/

func Combine(n int, k int) [][]int {
	var result [][]int
	var path []int // 存放符合条件结果

	backtracking(n, k, 1, path, &result)
	return result
}

// trim 剪枝
func backtracking(n, k, startIdx int, path []int, result *[][]int) {
	if len(path) == k {
		// 创建 path 的副本，避免切片引用问题
		combination := make([]int, len(path))
		copy(combination, path)
		*result = append(*result, combination)
		return
	}

	// trim n-i < k
	for i := startIdx; i <= n; i++ {
		path = append(path, i)
		backtracking(n, k, i+1, path, result)
		path = path[:len(path)-1]
	}
}
