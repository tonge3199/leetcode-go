package main

import (
	"fmt"
)

// Finish movingLongestPath
// Question : 寻找任意节点k到t的最长路径
// stdin and stdout
// eg input
// line1 : n,m
// n is 节点总数量
// m is 最多可经过值为1的节点数量
// line2 : 0 1 0 1 0 1 ...
// line2 is ai (i from 1 to n), the value of node
// and after n-1 line is the ui -> vi ，构建节点ui连接vi
// eg 1 2
// 3 4
// ...

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	// 读取节点值
	values := make([]int, n+1) // 1-indexed
	for i := 1; i <= n; i++ {
		fmt.Scan(&values[i])
	}

	// 构建邻接表
	graph := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u) // 无向图
	}

	result := movingLongestPath(n, m, values, graph)
	fmt.Println(result)
}

// movingLongestPath 寻找满足约束条件的最长路径
// 需要考虑所有可能的起点和终点组合
func movingLongestPath(n, m int, values []int, graph [][]int) int {
	maxPath := 0

	// 对每个节点作为起点
	for start := 1; start <= n; start++ {
		// 使用DFS从起点开始寻找所有可能的路径
		visited := make([]bool, n+1)
		dfsAllPaths(start, -1, 0, m, values, graph, visited, &maxPath)
	}

	return maxPath
}
func movingLongestPathII(n, m int, value []int, graph [][]int) int {
	maxPath := 0

	// each node as start
	for start := 1; start <= n; start++ {
		// DFS all possible paths
		visited := make([]bool, n+1)
		dfsAllPathII(start, -1, 0, m, graph, value, visited, &maxPath)
	}

	return maxPath
}

// dfsAllPaths 深度优先搜索所有可能的路径
func dfsAllPaths(current, parent, currentLength, remainingOnes int, values []int, graph [][]int, visited []bool, maxPath *int) {
	// 检查当前节点是否可访问
	currentOnes := remainingOnes
	if values[current] == 1 {
		if currentOnes <= 0 {
			return // 不能访问这个节点
		}
		currentOnes--
	}

	visited[current] = true

	// 更新最大路径长度
	if currentLength > *maxPath {
		*maxPath = currentLength
	}

	// 继续探索邻居节点
	for _, neighbor := range graph[current] {
		if neighbor != parent && !visited[neighbor] {
			dfsAllPaths(neighbor, current, currentLength+1, currentOnes, values, graph, visited, maxPath)
		}
	}

	visited[current] = false // 回溯
}

// dfsAllPathII 修正版本：正确处理路径长度和回溯
func dfsAllPathII(current, parent, currentLength, m int, graph [][]int, values []int, visited []bool, maxPath *int) {
	// 检查当前节点是否可访问
	currentOnes := m
	if values[current] == 1 {
		if currentOnes <= 0 {
			return // 不能访问这个节点
		}
		currentOnes--
	}

	// 标记为已访问
	visited[current] = true

	// 更新最大路径长度
	if currentLength > *maxPath {
		*maxPath = currentLength
	}

	// 继续探索邻居节点
	for _, nextNode := range graph[current] {
		if !visited[nextNode] && nextNode != parent {
			dfsAllPathII(nextNode, current, currentLength+1, currentOnes, graph, values, visited, maxPath)
		}
	}

	// 回溯：取消访问标记
	visited[current] = false
}
