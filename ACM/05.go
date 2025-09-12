package main

import (
	"fmt"
	"strings"
)

/*
题目
提交记录
tong
子序列(golang)
时间限制：C/C++语言1000MS；其他语言3000MS
内存限制：C/C++语言262144KB；其他语言786432KB

题目描述：
定义两个字符串是等价的，当且仅当其中一个串可以通过重新排列这些字符得到另一个串。例如，abccb 和 cbcba 等价，a 和 a 等价，abba 和 baab 等价。而 abc 和 aac 不等价，a 和 b 不等价。
现在输入 n 个仅由小写字母组成的字符串 s₁, s₂, …, sₙ，你需要找到一个长度最长的字符串 t，使得每个串都能找到一个子序列，这个子序列形成的字符串与 t 等价。
如果有多个答案，请输出字典序最小的串。如果找不到，则输出 -1。

输入描述
第一行输入一个正整数 T (1 ≤ T ≤ 5)，表示数据组数。
对于每一组数据：
第一行输入一个正整数 n (1 ≤ n ≤ 10)，表示字符串的个数。
第二行输入 n 个仅由小写字母组成的字符串 s₁, s₂, …, sₙ。每两个字符串之间用一个空格隔开，末尾没有多余空格。
保证同一组数据的字符串长度之和不超过 10⁵。

输出描述
对于每一组数据，输出一行。如果有多个答案，请输出字典序最小的串。如果找不到，则输出 -1。
*/

func main() {
	var T int
	fmt.Scan(&T) // 读取测试用例数量

	for t := 0; t < T; t++ {
		var n int
		fmt.Scan(&n) // 读取字符串个数

		inputStrings := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&inputStrings[i]) // 读取每个字符串
		}

		// 找到最长的等价子序列
		result := myFindSubseq(inputStrings)
		fmt.Println(result)
	}
}

// findLongestEquivalentSubsequence 找到最长的字符串t，使得每个输入字符串都包含与t等价的子序列
// 解题思路：
// 1. 等价字符串意味着字符频次相同
// 2. 子序列可以从原字符串中按顺序选取字符
// 3. 关键：t中每个字符的数量不能超过任何输入字符串中该字符的最少数量
// 4. 为保证字典序最小，按a-z顺序构建结果
func findLongestEquivalentSubsequence(inputStrings []string) string {
	if len(inputStrings) == 0 {
		return "-1"
	}

	// 对于每个字符a-z，找到所有字符串中该字符的最小数量
	// 这表示我们最多可以在结果字符串t中使用该字符多少次
	minCount := make([]int, 26)

	// 用第一个字符串的字符计数初始化
	firstCount := getCharCount(inputStrings[0])
	copy(minCount, firstCount)

	// 遍历其余字符串，找到每个字符的最小出现次数
	for i := 1; i < len(inputStrings); i++ {
		count := getCharCount(inputStrings[i])
		for j := 0; j < 26; j++ {
			if count[j] < minCount[j] {
				minCount[j] = count[j] // 更新为更小的计数
			}
		}
	}

	// 使用最小计数构建结果字符串
	// 为了得到字典序最小的结果，按a-z的顺序添加字符
	var result strings.Builder
	for i := 0; i < 26; i++ {
		char := byte('a' + i)
		for j := 0; j < minCount[i]; j++ {
			result.WriteByte(char) // 添加字符到结果中
		}
	}

	resultStr := result.String()
	if resultStr == "" {
		return "-1" // 如果没有公共字符，返回-1
	}

	return resultStr
}

// getCharCount 返回字符串中每个字符a-z的频次计数
// 输入：字符串s
// 输出：长度为26的数组，count[0]表示'a'的数量，count[1]表示'b'的数量，以此类推
func getCharCount(s string) []int {
	count := make([]int, 26)
	for _, c := range s {
		count[c-'a']++ // 将字符转换为数组索引并增加计数
	}
	return count
}

func myFindSubseq(inputStrings []string) string {

	if len(inputStrings) == 0 {
		return "-1"
	}

	// a-z min 找到所有字符串中该字符的最小数量
	// 这表示我们最多可以在结果字符串t中使用该字符多少次
	minCount := make([]int, 26)

	firstCount := getCharCount(inputStrings[0])
	// copy(dst, src []Type) copies elements from a source slice into a destination slice
	copy(minCount, firstCount)

	for i := 0; i < len(inputStrings); i++ {
		s := inputStrings[i]
		count := getCharCount(s)
		for j := 0; j < 26; j++ {
			if count[j] < minCount[j] {
				minCount[j] = count[j]
			}
		}

	}

	var result strings.Builder
	for i := 0; i < 26; i++ {
		c := byte(i + 'a')
		for j := 0; j < minCount[i]; j++ {
			result.WriteByte(c)
		}
	}

	resultStr := result.String()
	if resultStr == "" {
		return "-1"
	}
	return resultStr
}
