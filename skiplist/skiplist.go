package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// 最大层数
	SKIPLIST_MAXLEVEL = 32
	// 概率因子，用于决定新节点的层数
	SKIPLIST_P = 0.25
)

// zskiplistNode 跳跃表节点结构
type zskiplistNode struct {
	obj      interface{}      // 成员对象
	score    float64          // 分值，用于排序
	backward *zskiplistNode   // 后退指针，指向前一个节点
	level    []zskiplistLevel // 层数组，每层包含前进指针和跨度
	// obj interface{} // 成员对象
	// score  float64 // sort used
	// backward *zskiplistNode
	// level 	[]zskiplistLevel
}

// zskiplistLevel 跳跃表节点的层结构
type zskiplistLevel struct {
	forward *zskiplistNode // 前进指针，指向下一个节点
	span    uint64         // 跨度，记录前进指针所指向节点和当前节点的距离
	// forward *zskiplistLevel
	// span  uint...
}

// zskiplist 跳跃表结构
type zskiplist struct {
	header *zskiplistNode // 指向表头节点
	tail   *zskiplistNode // 指向表尾节点
	length int64          // 跳跃表长度（不包括表头节点）
	level  int            // 当前跳跃表内层数最大的节点层数（表头节点层数不计算在内）
}

// NewSkiplist 创建新的跳跃表
func NewSkiplist() *zskiplist {
	sl := &zskiplist{
		level:  1,
		length: 0,
	}

	// 创建表头节点，层数为最大层数
	sl.header = &zskiplistNode{
		level: make([]zskiplistLevel, SKIPLIST_MAXLEVEL),
	}

	// 初始化表头节点的所有层
	for i := 0; i < SKIPLIST_MAXLEVEL; i++ {
		sl.header.level[i].forward = nil
		sl.header.level[i].span = 0
	}

	sl.header.backward = nil
	sl.tail = nil

	return sl
}

// randomLevel 随机生成节点层数
func (sl *zskiplist) randomLevel() int {
	level := 1
	for rand.Float64() < SKIPLIST_P && level < SKIPLIST_MAXLEVEL {
		level++
	}
	return level
}

// Insert 插入节点
func (sl *zskiplist) Insert(score float64, obj interface{}) *zskiplistNode {
	update := make([]*zskiplistNode, SKIPLIST_MAXLEVEL) // 记录每层需要更新的节点
	rank := make([]uint64, SKIPLIST_MAXLEVEL)           // 记录每层的排名

	x := sl.header

	// 从最高层开始，逐层向下查找插入位置
	for i := sl.level - 1; i >= 0; i-- {
		if i == sl.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		// 在当前层查找合适的插入位置
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score && compareObj(x.level[i].forward.obj, obj) < 0)) {
			rank[i] += x.level[i].span
			x = x.level[i].forward
		}
		update[i] = x
	}

	// 随机生成新节点的层数
	level := sl.randomLevel()

	// 如果新节点的层数大于当前跳跃表的最大层数，需要更新跳跃表的层数
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			rank[i] = 0
			update[i] = sl.header
			update[i].level[i].span = uint64(sl.length)
		}
		sl.level = level
	}

	// 创建新节点
	x = &zskiplistNode{
		obj:   obj,
		score: score,
		level: make([]zskiplistLevel, level),
	}

	// 更新各层的指针和跨度
	for i := 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x

		// 更新跨度
		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	// 更新未涉及到的层的跨度
	for i := level; i < sl.level; i++ {
		update[i].level[i].span++
	}

	// 设置后退指针
	if update[0] == sl.header {
		x.backward = nil
	} else {
		x.backward = update[0]
	}

	// 更新后退指针和表尾指针
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		sl.tail = x
	}

	sl.length++
	return x
}

// Delete 删除节点
func (sl *zskiplist) Delete(score float64, obj interface{}) bool {
	update := make([]*zskiplistNode, SKIPLIST_MAXLEVEL)

	x := sl.header

	// 查找要删除的节点
	for i := sl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score && compareObj(x.level[i].forward.obj, obj) < 0)) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward

	// 检查是否找到要删除的节点
	if x != nil && score == x.score && compareObj(x.obj, obj) == 0 {
		sl.deleteNode(x, update)
		return true
	}

	return false
}

// deleteNode 删除指定节点
func (sl *zskiplist) deleteNode(x *zskiplistNode, update []*zskiplistNode) {
	// 更新各层的指针和跨度
	for i := 0; i < sl.level; i++ {
		if update[i].level[i].forward == x {
			update[i].level[i].span += x.level[i].span - 1
			update[i].level[i].forward = x.level[i].forward
		} else {
			update[i].level[i].span--
		}
	}

	// 更新后退指针
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x.backward
	} else {
		sl.tail = x.backward
	}

	// 如果删除的是最高层的节点，需要更新跳跃表的层数
	for sl.level > 1 && sl.header.level[sl.level-1].forward == nil {
		sl.level--
	}

	sl.length--
}

// Search 搜索节点
func (sl *zskiplist) Search(score float64, obj interface{}) *zskiplistNode {
	x := sl.header

	// 从最高层开始搜索
	for i := sl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score && compareObj(x.level[i].forward.obj, obj) < 0)) {
			x = x.level[i].forward
		}
	}

	x = x.level[0].forward

	// 检查是否找到目标节点
	if x != nil && score == x.score && compareObj(x.obj, obj) == 0 {
		return x
	}

	return nil
}

// GetRank 获取节点排名（从1开始）
func (sl *zskiplist) GetRank(score float64, obj interface{}) int64 {
	x := sl.header
	rank := uint64(0)

	// 从最高层开始查找
	for i := sl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score && compareObj(x.level[i].forward.obj, obj) <= 0)) {
			rank += x.level[i].span
			x = x.level[i].forward
		}

		// 如果找到了目标节点
		if x.obj != nil && compareObj(x.obj, obj) == 0 && x.score == score {
			return int64(rank)
		}
	}

	return 0
}

// GetElementByRank 根据排名获取节点（排名从1开始）
func (sl *zskiplist) GetElementByRank(rank uint64) *zskiplistNode {
	x := sl.header
	traversed := uint64(0)

	// 从最高层开始查找
	for i := sl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (traversed+x.level[i].span) <= rank {
			traversed += x.level[i].span
			x = x.level[i].forward
		}

		if traversed == rank {
			return x
		}
	}

	return nil
}

// Length 获取跳跃表长度
func (sl *zskiplist) Length() int64 {
	return sl.length
}

// Level 获取跳跃表当前最大层数
func (sl *zskiplist) Level() int {
	return sl.level
}

// Print 打印跳跃表结构（用于调试）
func (sl *zskiplist) Print() {
	fmt.Printf("跳跃表信息: 长度=%d, 最大层数=%d\n", sl.length, sl.level)
	fmt.Print("表头 -> ")

	node := sl.header.level[0].forward
	for node != nil {
		fmt.Printf("[分值:%.1f, 对象:%v] -> ", node.score, node.obj)
		node = node.level[0].forward
	}
	fmt.Println("表尾")
}

// compareObj 比较两个对象，返回-1、0、1分别表示小于、等于、大于
func compareObj(a, b interface{}) int {
	switch av := a.(type) {
	case string:
		if bv, ok := b.(string); ok {
			if av < bv {
				return -1
			} else if av > bv {
				return 1
			}
			return 0
		}
	case int:
		if bv, ok := b.(int); ok {
			if av < bv {
				return -1
			} else if av > bv {
				return 1
			}
			return 0
		}
	}

	// 默认按字符串比较
	as := fmt.Sprintf("%v", a)
	bs := fmt.Sprintf("%v", b)
	if as < bs {
		return -1
	} else if as > bs {
		return 1
	}
	return 0
}

// init 初始化随机数种子
func init() {
	rand.Seed(time.Now().UnixNano())
}
