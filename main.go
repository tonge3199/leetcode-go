package main

// 使用 fmt 包进行 ACM 风格的 stdin 输入处理
import (
	"fmt"

	"github.com/tonge3199/leetcode/dp"
)

func main() {
	var m, n int

	// Read dimensions from stdin
	fmt.Print("Enter m and n (rows and columns): ")
	fmt.Scanf("%d %d", &m, &n)

	// Create and initialize m×n grid
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	// Read the grid with obstacles from stdin
	fmt.Printf("Enter the %d×%d grid (0 for empty, 1 for obstacle):\n", m, n)
	for i := 0; i < m; i++ {
		fmt.Printf("Row %d: ", i+1)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &grid[i][j])
		}
	}

	// Display the input grid
	fmt.Printf("\nInput grid:\n")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}

	// Test both UniquePath and UniquePathII functions
	fmt.Println("\n=== Testing UniquePath (ignores obstacles) ===")
	paths1 := dp.UniquePath(grid)
	fmt.Printf("Number of unique paths: %d\n", paths1)

	fmt.Println("\n=== Testing UniquePathII (considers obstacles) ===")
	paths2 := dp.UniquePathII(grid)
	fmt.Printf("Number of unique paths with obstacles: %d\n", paths2)
}
