package ac2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 树哈希
// 设定一棵树 x 的哈希值为 hash(x)
// hash(x) = x.Val + 31 * hash(x.Left) * primes[N(x.Left)] + 179 * hash(x.Right) * primes[N(x.Right)]
// 其中 primes[n] 表示第 n 个素数，N(x) 表示树 x 的节点个数
// 由于 subRoot 最多只有 1000 个节点，可以通过埃式筛法求出前 1000 个素数
// 然后通过递归计算 root 中每个节点的 hash 值并保存在 map 中
// 再计算 subRoot 的哈希值
// 遍历一遍结果 map，如果存在与 subRoot 相同的哈希值，说明 subRoot 是 root 的子树
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	primes := getPrime()
	hashMap := make(map[*TreeNode]int)

	hash(root, hashMap, primes)
	s, _ := hash(subRoot, nil, primes)

	for _, val := range hashMap {
		if val == s {
			return true
		}
	}
	return false
}

func getPrime() []int {
	vis := make(map[int]struct{})
	vis[0] = struct{}{}
	vis[1] = struct{}{}
	primes := []int{-1}
	for i := 2; i < 100000; i++ {
		if _, ok := vis[i]; !ok {
			primes = append(primes, i)
		}
		if len(primes) > 1000 {
			break
		}
		for j := 1; j < len(primes) && primes[j]*i < 100000; j++ {
			vis[i*primes[j]] = struct{}{}
		}
	}
	return primes
}

func hash(node *TreeNode, hashMap map[*TreeNode]int, primes []int) (int, int) {
	if node.Left == nil && node.Right == nil {
		h := node.Val
		if hashMap != nil {
			hashMap[node] = node.Val
		}
		return h, 1
	}
	var hl, hr, nl, nr int
	if node.Left != nil {
		hl, nl = hash(node.Left, hashMap, primes)
	}
	if node.Right != nil {
		hr, nr = hash(node.Right, hashMap, primes)
	}
	h := 0
	if nl <= 100 && nr <= 1000 {
		h = node.Val + 31*hl*primes[nl] + 179*hr*primes[nr]
	}
	if hashMap != nil {
		hashMap[node] = h
	}
	return h, nl + nr + 1
}
