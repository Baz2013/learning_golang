package lcof

import (
	"fmt"
	"testing"
)

func TestRecreateBinaryTree(t *testing.T) {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(root)

	//前序遍历结果
	preorder := []int{3, 9, 20, 15, 7}
	//中序遍历结果
	inorder := []int{9, 3, 15, 20, 7}

	/**

	    3
	   / \
	  9  20
	    /  \
	   15   7

	*/

	RecreateBinaryTree(preorder, inorder)

}
