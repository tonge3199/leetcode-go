package hashMap

import "strings"

/*
Every valid email consists of a local name and a domain name, separated by the '@' sign. Besides lowercase letters, the email may contain one or more '.' or '+'.

For example, in "alice@leetcode.com", "alice" is the local name, and "leetcode.com" is the domain name.
If you add periods '.' between some characters in the local name part of an email address, mail sent there will be forwarded to the same address without dots in the local name. Note that this rule does not apply to domain names.

For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.
If you add a plus '+' in the local name, everything after the first plus sign will be ignored. This allows certain emails to be filtered. Note that this rule does not apply to domain names.

For example, "m.y+name@email.com" will be forwarded to "my@email.com".
It is possible to use both of these rules at the same time.

Given an array of strings emails where we send one email to each emails[i], return the number of different addresses that actually receive mails.
*/

func numUniqueEmails(emails []string) int {
	emailSet := make(map[string]bool)

	for _, email := range emails {
		// 1. 解析为本地名称和域名两部分
		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			continue // 跳过无效的邮件地址
		}
		localName := parts[0]
		domainName := parts[1]

		// 2. 处理本地名称
		// 移除所有的 "."
		localName = strings.ReplaceAll(localName, ".", "")

		// 忽略第一个 "+" 之后的所有内容
		if plusIndex := strings.Index(localName, "+"); plusIndex != -1 {
			localName = localName[:plusIndex]
		}

		// 3. 域名保持不变，重新组合邮件地址
		normalizedEmail := localName + "@" + domainName

		// 使用 set 来统计唯一邮件地址
		emailSet[normalizedEmail] = true
	}

	return len(emailSet)
}

// TODO: 重新写一遍 并写ACM/测试文件
