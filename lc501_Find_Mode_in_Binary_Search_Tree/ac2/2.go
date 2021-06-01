package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris
// 时间O(n)
// 空间O(1)
// 不用栈空间，直接修改前驱节点的右子树为当前节点
func findMode(root *TreeNode) []int {
	var (
		ans []int
		cur, count, maxCount = -1, 0, 0
		update func(x int)
		curNode, preNode *TreeNode = root, nil
	)

	update = func(x int) {
		if x != cur {
			count = 1
			cur = x
		} else {
			count++
		}
		if count > maxCount {
			ans = []int{cur}
			maxCount = count
		} else if count == maxCount {
			ans = append(ans, cur)
		}
	}

	// 中序遍历
	for curNode != nil {
		if curNode.Left == nil { // 如果当前节点的左子树为空，则没有前驱节点，直接遍历，并进入右子树
			update(curNode.Val)
			curNode = curNode.Right
			continue
		}

		// 否则，找到当前节点的左子树的最右下方的节点
		preNode = curNode.Left
		for preNode.Right != nil && preNode.Right != curNode {
			preNode = preNode.Right
		}
		if preNode.Right == nil { // 满足该条件说明 当前找到了当前节点的前驱节点
			preNode.Right = curNode // 则把前驱节点的右子树置为当前节点，这样遍历完这个前驱节点之后就可以直接遍历到当前节点，不需要占空间
			curNode = curNode.Left // 继续中序遍历
		} else { // 否则，当前节点的前驱节点已经找过了，不需要再遍历左子树
			preNode.Right = nil // 将前驱节点的右子节点置空
			update(curNode.Val)
			curNode = curNode.Right // 继续中序遍历
		}
	}
	return ans
}
