package main

import "fmt"

/*
有效的重复字符
详细描述
给定一个经过编码的字符串，按照特定规则返回它解码后的字符串。

编码规则为：k{string}，表示大括号内部的 string 经过解码后重复 k 次，k 保证为正整数，string 经过解码后为由 a-z 之间的字符组成的字符串，即大括号可能会有嵌套的情况。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的括号总是符合格式要求的。

原始数据不包含数字，所有的数字只表示重复的次数 k，例如不会出现像 3a 或 2{4} 的输入，但是会出现像 2{a3{b4{c}d}e} 的情况。

其他
时间限制：3000ms

内存限制：256.0MB

输入输出示例
示例1
输入："2{ac}"

输出："acac"

示例2
输入："2{ab3{ac}}"

输出："abacacacabacacac"
*/
func main() {
	// 竞赛格式的输入输出
	var s string
	fmt.Scan(&s)
	result := decodeString2(s)
	fmt.Println(result)

	// 测试示例（可以注释掉）
	fmt.Println("测试 '2{ac}':", decodeString2("2{ac}"))
	fmt.Println("测试 '2{ab3{ac}}':", decodeString2("2{ab3{ac}}"))
	fmt.Println("测试 '3{2{ac}b}':", decodeString2("3{2{ac}b}"))
}

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	currentNum := 0
	currentStr := ""

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c >= '0' && c <= '9' {
			// 构建数字（可能是多位数）
			currentNum = currentNum*10 + int(c-'0')
		} else if c == '{' {
			// 遇到左括号，将当前数字和字符串压入栈
			numStack = append(numStack, currentNum)
			strStack = append(strStack, currentStr)
			currentNum = 0
			currentStr = ""
		} else if c == '}' {
			// 遇到右括号，弹出栈顶元素进行解码
			prevNum := numStack[len(numStack)-1]
			prevStr := strStack[len(strStack)-1]
			numStack = numStack[:len(numStack)-1]
			strStack = strStack[:len(strStack)-1]

			// 重复currentStr prevNum次，然后拼接到prevStr后面
			repeated := ""
			for j := 0; j < prevNum; j++ {
				repeated += currentStr
			}
			currentStr = prevStr + repeated
		} else {
			// 普通字符，直接拼接
			currentStr += string(c)
		}
	}

	return currentStr
}

func decodeString2(s string) string {
	numStack := []int{}
	strStack := []string{}
	currentNum := 0
	currentStr := ""

	for i := 0; i < len(s); i++ {
		c := s[i]
		if '0' <= c && c <= '9' {
			currentNum = currentNum*10 + int(c-'0')
		} else if c == '{' {
			// flash in prevStr
			strStack = append(strStack, currentStr)
			currentStr = ""
			// flash in prevNum too
			numStack = append(numStack, currentNum)
			currentNum = 0
		} else if c == '}' {
			// outStack , use other val to store state avoid confusing.
			prevNum := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			// construct the answer
			repeated := ""
			for i := 0; i < prevNum; i++ { // currentStr repeated 1.
				repeated += currentStr
			}
			currentStr = prevStr + repeated
		} else {
			// 'a' - 'z'
			currentStr = currentStr + string(c)
		}
	}
	return currentStr
}
