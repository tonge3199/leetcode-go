package string

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// 以第一个字符串作为初始公共前缀
	commonPrefix := strs[0]

	for i := 1; i < len(strs); i++ {
		// 找到当前公共前缀和下一个字符串的公共前缀
		j := 0
		minLen := len(commonPrefix)
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}

		for j < minLen && commonPrefix[j] == strs[i][j] {
			j++
		}

		// 更新公共前缀
		commonPrefix = commonPrefix[:j]

		// 如果公共前缀已经为空，直接返回
		if commonPrefix == "" {
			return ""
		}
	}

	return commonPrefix
}
