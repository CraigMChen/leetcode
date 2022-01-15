package ac1

// 找规律（动态规划）
// 遍历1-n，设把i作为根节点的BST的数量为G(i)
// 对于i为根节点的BST，不同左子树的数量为G(i - 1)，不同右子树的数量为G(n - i)
// 则G(i) = G(i - 1) * G(n - i)
// G(0) = 1
// G(1) = 1
//func numTrees(n int) int {
//	nums := []int{1, 1}
//	for i := 2; i <= n; i++ {
//		count := 0
//		for j := 1; j <= i; j++ {
//			count += nums[j - 1] * nums[i - j]
//		}
//		nums = append(nums, count)
//	}
//	return nums[n]
//}

var nums []int

func init() {
	nums = make([]int, 20)
	nums[0] = 1
	nums[1] = 1
	for i := 2; i <= 19; i++ {
		for j := 1; j <= i; j++ {
			nums[i] += nums[j-1] * nums[i - j]
		}
	}
}

func numTrees(n int) int {
	return nums[n]
}