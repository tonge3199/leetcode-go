package tree

import (
	"testing"
)

// createTree creates a binary tree from a slice of interface{} values (allowing nil for null nodes)
func createTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func TestSymmetricTree(t *testing.T) {
	tests := []struct {
		name      string
		inputRoot *TreeNode
		want      bool
	}{
		{
			name:      "Symmetric tree [1,2,2,3,4,4,3]",
			inputRoot: createTree([]interface{}{1, 2, 2, 3, 4, 4, 3}),
			want:      true,
		},
		{
			name:      "Non-symmetric tree [1,2,2,null,3,null,3]",
			inputRoot: createTree([]interface{}{1, 2, 2, nil, 3, nil, 3}),
			want:      false,
		},
		{
			name:      "Empty tree",
			inputRoot: nil,
			want:      true,
		},
		{
			name:      "Single node [1]",
			inputRoot: createTree([]interface{}{1}),
			want:      true,
		},
		{
			name:      "Symmetric deeper tree [1,2,2,3,4,4,3,5,6,6,5]",
			inputRoot: createTree([]interface{}{1, 2, 2, 3, 4, 4, 3, 5, 6, 6, 5}),
			want:      false,
		},
		{
			name:      "Non-symmetric values [1,2,2,3,4,3,4]",
			inputRoot: createTree([]interface{}{1, 2, 2, 3, 4, 3, 4}),
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := isSymmetric_01(tt.inputRoot); got != tt.want {
			//	t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			//}
			if got := isSymmetric_02(tt.inputRoot); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
