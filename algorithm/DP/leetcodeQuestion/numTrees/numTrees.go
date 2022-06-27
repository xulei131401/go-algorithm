package main

import "fmt"

/**
题目:https://leetcode-cn.com/problems/unique-binary-search-trees/

不同的二叉搜索树
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

提示：

1 <= n <= 19


*/
func main() {
	n := 3
	fmt.Println("不同的二叉搜索树种类：", numTrees(n))
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
