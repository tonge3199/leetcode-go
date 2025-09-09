package main

import (
	"fmt"
)

// TODO : Learning this.

/* monitor SQL:
有一个hive表，表名是t，字段是a和b，a和b的值都是正整数（不存在null），且a和b的最大值都小于100。t的数据行数小于1000。

编程实现下面SQL的逻辑，输出SQL的结果：
select sum(rn) from (
select row_number() over(PARTITION BY case when t1.a is not null then t1.a else t2.b end) as rn
from t as t1 full join t as t2
on t1.a = t2.b and t1.b>10 and t2.a>10
) tmp
*/

// Row 表示连接后的行
type Row struct {
	t1a *int // t1.a，可能为null
	t1b *int // t1.b，可能为null
	t2a *int // t2.a，可能为null
	t2b *int // t2.b，可能为null
}

// GetPartitionKey 获取分组键：coalesce(t1.a, t2.b)
func (r Row) GetPartitionKey() int {
	if r.t1a != nil {
		return *r.t1a
	}
	return *r.t2b
}

func main() {
	// 从标准输入读取数据（竞赛格式）
	var n int
	fmt.Scan(&n)
	t := make([][]int, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		t[i] = []int{a, b}
	}

	// 示例数据用于测试（可注释掉上面的输入代码，启用下面的测试数据）
	/*
		var t = [][]int{
			{1, 15},  // a=1, b=15
			{2, 12},  // a=2, b=12
			{1, 8},   // a=1, b=8
			{3, 20},  // a=3, b=20
			{15, 5},  // a=15, b=5
			{12, 25}, // a=12, b=25
		}
	*/

	result := simulateSQL(t)
	fmt.Println(result)
}

func simulateSQL(t [][]int) int {
	// 1. 模拟 full join
	var rows []Row
	n := len(t)

	// 记录哪些行已经匹配过
	leftMatched := make([]bool, n)
	rightMatched := make([]bool, n)

	// 找到所有匹配的行对
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			t1a, t1b := t[i][0], t[i][1]
			t2a, t2b := t[j][0], t[j][1]

			// 检查连接条件：t1.a = t2.b AND t1.b > 10 AND t2.a > 10
			if t1a == t2b && t1b > 10 && t2a > 10 {
				rows = append(rows, Row{
					t1a: &t1a,
					t1b: &t1b,
					t2a: &t2a,
					t2b: &t2b,
				})
				leftMatched[i] = true
				rightMatched[j] = true
			}
		}
	}

	// 添加左表未匹配的行（t1有值，t2为null）
	for i := 0; i < n; i++ {
		if !leftMatched[i] {
			t1a, t1b := t[i][0], t[i][1]
			rows = append(rows, Row{
				t1a: &t1a,
				t1b: &t1b,
				t2a: nil,
				t2b: nil,
			})
		}
	}

	// 添加右表未匹配的行（t2有值，t1为null）
	for j := 0; j < n; j++ {
		if !rightMatched[j] {
			t2a, t2b := t[j][0], t[j][1]
			rows = append(rows, Row{
				t1a: nil,
				t1b: nil,
				t2a: &t2a,
				t2b: &t2b,
			})
		}
	}

	// 2. 按分组键分组
	groups := make(map[int][]Row)
	for _, row := range rows {
		// 只有当分组键有效时才处理（t1.a不为null或t2.b不为null）
		if row.t1a != nil || row.t2b != nil {
			key := row.GetPartitionKey()
			groups[key] = append(groups[key], row)
		}
	}

	// 3. 计算每个分组内的row_number()并求和
	totalSum := 0
	for _, group := range groups {
		// 为每个组内的行分配行号（从1开始）
		for i := 0; i < len(group); i++ {
			rowNumber := i + 1
			totalSum += rowNumber
		}
	}

	return totalSum
}
