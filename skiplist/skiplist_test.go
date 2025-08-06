package skiplist

import (
	"fmt"
	"testing"
)

// TestBasicOperations 测试基本操作
func TestBasicOperations(t *testing.T) {
	fmt.Println("=== 测试跳跃表基本操作 ===")

	// 创建跳跃表
	sl := NewSkiplist()

	// 测试初始状态
	if sl.Length() != 0 {
		t.Errorf("新创建的跳跃表长度应为0，实际为%d", sl.Length())
	}

	if sl.Level() != 1 {
		t.Errorf("新创建的跳跃表层数应为1，实际为%d", sl.Level())
	}

	fmt.Printf("初始化跳跃表成功，长度: %d, 层数: %d\n", sl.Length(), sl.Level())
}

// TestInsert 测试插入操作
func TestInsert(t *testing.T) {
	fmt.Println("\n=== 测试插入操作 ===")

	sl := NewSkiplist()

	// 插入测试数据
	testData := []struct {
		score float64
		obj   interface{}
	}{
		{1.0, "o1"},
		{2.0, "o2"},
		{3.0, "o3"},
		{1.5, "o4"},
		{2.5, "o5"},
	}

	for _, data := range testData {
		node := sl.Insert(data.score, data.obj)
		if node == nil {
			t.Errorf("插入节点失败: score=%.1f, obj=%v", data.score, data.obj)
		}
		fmt.Printf("插入成功: 分值=%.1f, 对象=%v\n", data.score, data.obj)
	}

	// 检查长度
	expectedLength := int64(len(testData))
	if sl.Length() != expectedLength {
		t.Errorf("插入后长度不正确，期望%d，实际%d", expectedLength, sl.Length())
	}

	fmt.Printf("插入完成，当前长度: %d, 层数: %d\n", sl.Length(), sl.Level())
	sl.Print()
}

// TestSearch 测试搜索操作
func TestSearch(t *testing.T) {
	fmt.Println("\n=== 测试搜索操作 ===")

	sl := NewSkiplist()

	// 插入测试数据
	sl.Insert(1.0, "o1")
	sl.Insert(2.0, "o2")
	sl.Insert(3.0, "o3")

	// 测试搜索存在的节点
	node := sl.Search(2.0, "o2")
	if node == nil {
		t.Error("搜索存在的节点失败")
	} else {
		fmt.Printf("搜索成功: 分值=%.1f, 对象=%v\n", node.score, node.obj)
	}

	// 测试搜索不存在的节点
	node = sl.Search(4.0, "o4")
	if node != nil {
		t.Error("搜索不存在的节点应返回nil")
	} else {
		fmt.Println("搜索不存在的节点正确返回nil")
	}
}

// TestDelete 测试删除操作
func TestDelete(t *testing.T) {
	fmt.Println("\n=== 测试删除操作 ===")

	sl := NewSkiplist()

	// 插入测试数据
	sl.Insert(1.0, "o1")
	sl.Insert(2.0, "o2")
	sl.Insert(3.0, "o3")
	sl.Insert(4.0, "o4")

	fmt.Println("删除前:")
	sl.Print()

	// 删除中间节点
	success := sl.Delete(2.0, "o2")
	if !success {
		t.Error("删除存在的节点失败")
	} else {
		fmt.Println("删除节点 [分值:2.0, 对象:o2] 成功")
	}

	fmt.Println("删除后:")
	sl.Print()

	// 验证删除结果
	node := sl.Search(2.0, "o2")
	if node != nil {
		t.Error("删除后节点仍然存在")
	}

	if sl.Length() != 3 {
		t.Errorf("删除后长度不正确，期望3，实际%d", sl.Length())
	}

	// 删除不存在的节点
	success = sl.Delete(5.0, "o5")
	if success {
		t.Error("删除不存在的节点应返回false")
	} else {
		fmt.Println("删除不存在的节点正确返回false")
	}
}

// TestRank 测试排名操作
func TestRank(t *testing.T) {
	fmt.Println("\n=== 测试排名操作 ===")

	sl := NewSkiplist()

	// 插入测试数据（按分值排序）
	testData := []struct {
		score float64
		obj   interface{}
		rank  int64
	}{
		{1.0, "o1", 1},
		{2.0, "o2", 2},
		{3.0, "o3", 3},
		{4.0, "o4", 4},
		{5.0, "o5", 5},
	}

	for _, data := range testData {
		sl.Insert(data.score, data.obj)
	}

	fmt.Println("插入完成:")
	sl.Print()

	// 测试获取排名
	for _, data := range testData {
		rank := sl.GetRank(data.score, data.obj)
		if rank != data.rank {
			t.Errorf("排名不正确，期望%d，实际%d", data.rank, rank)
		} else {
			fmt.Printf("节点 [分值:%.1f, 对象:%v] 排名: %d\n", data.score, data.obj, rank)
		}
	}

	// 测试根据排名获取元素
	fmt.Println("\n根据排名获取元素:")
	for i := uint64(1); i <= uint64(len(testData)); i++ {
		node := sl.GetElementByRank(i)
		if node == nil {
			t.Errorf("获取排名%d的元素失败", i)
		} else {
			fmt.Printf("排名%d: [分值:%.1f, 对象:%v]\n", i, node.score, node.obj)
		}
	}
}

// TestComplexOperations 测试复杂操作
func TestComplexOperations(t *testing.T) {
	fmt.Println("\n=== 测试复杂操作 ===")

	sl := NewSkiplist()

	// 插入大量数据
	fmt.Println("插入大量数据...")
	for i := 0; i < 100; i++ {
		score := float64(i % 10) // 0.0 - 9.0
		obj := fmt.Sprintf("obj%d", i)
		sl.Insert(score, obj)
	}

	fmt.Printf("插入完成，长度: %d, 层数: %d\n", sl.Length(), sl.Level())

	// 测试搜索性能
	fmt.Println("测试搜索...")
	found := 0
	for i := 0; i < 100; i++ {
		score := float64(i % 10)
		obj := fmt.Sprintf("obj%d", i)
		if sl.Search(score, obj) != nil {
			found++
		}
	}
	fmt.Printf("搜索完成，找到 %d 个元素\n", found)

	// 测试删除一些元素
	fmt.Println("删除部分元素...")
	deleted := 0
	for i := 0; i < 50; i += 5 {
		score := float64(i % 10)
		obj := fmt.Sprintf("obj%d", i)
		if sl.Delete(score, obj) {
			deleted++
		}
	}
	fmt.Printf("删除完成，删除了 %d 个元素，剩余长度: %d\n", deleted, sl.Length())
}

// TestBackwardTraversal 测试后向遍历
func TestBackwardTraversal(t *testing.T) {
	fmt.Println("\n=== 测试后向遍历 ===")

	sl := NewSkiplist()

	// 插入测试数据
	scores := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	for i, score := range scores {
		sl.Insert(score, fmt.Sprintf("o%d", i+1))
	}

	fmt.Println("正向遍历:")
	sl.Print()

	fmt.Println("后向遍历:")
	node := sl.tail
	fmt.Print("表尾 -> ")
	for node != nil {
		fmt.Printf("[分值:%.1f, 对象:%v] -> ", node.score, node.obj)
		node = node.backward
	}
	fmt.Println("表头")
}

// BenchmarkInsert 插入操作性能测试
func BenchmarkInsert(b *testing.B) {
	sl := NewSkiplist()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		score := float64(i % 1000)
		obj := fmt.Sprintf("obj%d", i)
		sl.Insert(score, obj)
	}
}

// BenchmarkSearch 搜索操作性能测试
func BenchmarkSearch(b *testing.B) {
	sl := NewSkiplist()

	// 预先插入数据
	for i := 0; i < 10000; i++ {
		score := float64(i % 1000)
		obj := fmt.Sprintf("obj%d", i)
		sl.Insert(score, obj)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		score := float64(i % 1000)
		obj := fmt.Sprintf("obj%d", i%10000)
		sl.Search(score, obj)
	}
}
