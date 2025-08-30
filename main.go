package main

// 使用 fmt 包进行 ACM 风格的 stdin 输入处理
import (
	"fmt"

	array "github.com/tonge3199/leetcode/array"
)

func main() {
	var nums []int

	// 读取一行输入，用空格分隔
	/*
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			fields := strings.Fields(line)
			for _, field := range fields {
				if num, err := strconv.Atoi(field); err == nil {
					nums = append(nums, num)
				}
			}
		}
	*/

	array.MoveZero(nums)
	fmt.Println(nums)
}
