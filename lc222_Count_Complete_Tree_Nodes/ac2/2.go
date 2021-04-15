package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二分，位运算
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	h := 0
	for tmp := root; tmp.Left != nil; tmp = tmp.Left {
		h++
	}
	if h == 0 {
		return 1
	}

	hasNode := func(k int) bool {
		tmp := root
		bits := 1 << (h - 1)
		for tmp != nil && bits > 0 {
			if k & bits == 0 {
				tmp = tmp.Left
			} else {
				tmp = tmp.Right
			}
			bits >>= 1
		}
		return tmp != nil
	}

	l, r := 1 << h, 1 << (h + 1)
	ans := 0
	for l < r {
		m := (r - l) >> 1 + l
		if hasNode(m) {
			ans = m
			l = m + 1
		} else {
			r = m
		}
	}
	return ans
}
