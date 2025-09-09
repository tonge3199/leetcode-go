package dp

/*
  m*n map
  right , down only
  从左上角 -> 右下角
*/

func UniquePath(Map [][]int) int {
	m := len(Map)
	if m == 0 {
		return 0
	}
	n := len(Map[0])
	if n == 0 {
		return 0
	}

	// init
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1]
}

// added obstacle
func UniquePathII(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	// init
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// consider obstacle now
	for i := 0; i < m; i++ {
		if grid[i][0] == 1 { // Is obstacle position , Fix: 第一列障碍物下面的格子都无法到达
			// dp[i][0] = 0 冗余
			break
		} else {
			dp[i][0] = 1
		}
	}
	for j := 0; j < n; j++ {
		if grid[0][j] == 1 { // Is obstacle position , Fix: 第一行障碍物右边的格子都无法到达
			// dp[0][j] = 0
			break
		} else {
			dp[0][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 1 { // Is obstacle position
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}

	return dp[m-1][n-1]
}
