package recursive

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	// isVisited[i] == true : recursive tree depth visited
	// isVisited[i] == false: recursive tree width visited
	used := make([]bool, len(candidates))
	sort.Ints(candidates)

	var dfs func(start, sum, target int)
	dfs = func(start, sum, target int) {
		if sum == target {
			comb := make([]int, len(path))
			copy(comb, path)
			ans = append(ans, comb)
			return
		}
		for i := start; i < len(candidates); i++ {
			num := candidates[i]
			if num+sum > target {
				continue
			}
			if i > 0 && num == candidates[i-1] && !used[i-1] {
				continue
			}
			path = append(path, num)
			used[i] = true
			dfs(i+1, num+sum, target) // i->i+1
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0, target)
	return ans
}
