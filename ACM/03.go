package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 问题：Java 代码注释去除
// 输入 stdin 一段Java代码
// 输出 去除注释后的代码，以及空行不输出（空行：只包含空格和\t的行）

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	// 读取所有输入行
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// 处理注释去除
	// result := removeComments(lines)
	result := removeComments2(lines)

	// 输出结果
	for _, line := range result {
		fmt.Println(line)
	}
}

func removeComments(lines []string) []string {
	var result []string
	inBlockComment := false

	for _, line := range lines {
		processedLine := ""
		i := 0

		for i < len(line) {
			if inBlockComment {
				// 在块注释中，寻找结束标记 */
				if i < len(line)-1 && line[i] == '*' && line[i+1] == '/' {
					inBlockComment = false
					i += 2
				} else {
					i++
				}
			} else {
				// 不在块注释中
				if i < len(line)-1 && line[i] == '/' && line[i+1] == '/' {
					// 遇到单行注释，忽略该行剩余部分
					break
				} else if i < len(line)-1 && line[i] == '/' && line[i+1] == '*' {
					// 遇到块注释开始
					inBlockComment = true
					i += 2
				} else {
					// 普通字符
					processedLine += string(line[i])
					i++
				}
			}
		}

		// 检查是否为空行（只包含空格和制表符）
		if !inBlockComment { // 只有不在块注释中时才添加行
			trimmed := strings.TrimSpace(processedLine)
			if trimmed != "" {
				result = append(result, processedLine)
			}
		}
	}

	return result
}

func removeComments2(lines []string) (result []string) {
	isBlockComment := false // block can be multi-lines.
	for _, line := range lines {
		processedLine := ""
		var i int = 0
		for i < len(line) {

			if !isBlockComment {
				if i < len(line)-1 && line[i] == '/' && line[i+1] == '/' {
					break // in lineComment
				}
				if i < len(line)-1 && line[i] == '/' && line[i+1] == '*' {
					isBlockComment = true
					i += 2
				} else {
					processedLine += string(line[i]) // optimize 01
					i++
				}
			} else { // isBlockComment
				if i < len(line)-1 && line[i] == '*' && line[i+1] == '/' {
					isBlockComment = false
					i += 2
				} else {
					i += 1
				}
			}

		}

		trimmed := strings.TrimSpace(processedLine)
		if trimmed != "" {
			result = append(result, processedLine)
		}
	}
	return
}
