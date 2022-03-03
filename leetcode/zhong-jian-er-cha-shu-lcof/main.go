package zhong_jian_er_cha_shu_lcof

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(inorder) == 0 {
		return nil
	}

	ans := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	idx := 0
	for ; idx < len(inorder); idx++ {
		if preorder[0] == inorder[idx] {
			break
		}
	}

	ans.Left = buildTree(preorder[1:1+idx], inorder[:idx])
	ans.Right = buildTree(preorder[idx+1:], inorder[1+idx:])
	return ans
}

// 迭代
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{}
	stack = append(stack, root)
	inorderIndex := 0
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				stack = stack[:len(stack)-1] // 出栈
				inorderIndex++
			}
			node = stack[len(stack)-1]
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}
