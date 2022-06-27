package main

import (
	"log"
)

/**
插入排序的种类：直接插入排序、折半插入排序和希尔排序3类
插入排序：插入排序，一般也被称为直接插入排序。对于少量元素的排序，它是一个有效的算法。
插入排序是一种最简单的排序方法，它的基本思想是将一个记录插入到已经排好序的有序表中，从而一个新的、记录数增1的有序表。
在其实现过程使用双层循环，外层循环对除了第一个元素之外的所有元素，内层循环对当前元素前面有序表进行待插入位置查找，并进行移动
1.双层for
2.外层从i=1开始
3.内层遍历当前元素之前的有序区间

稳定性分析:如果待排序的序列中存在两个或两个以上具有相同关键词的数据，排序后这些数据的相对次序保持不变，即它们的位置保持不变，通俗地讲，就是两个相同的数的相对顺序不会发生改变，则该算法是稳定的；
		 如果排序后，数据的相对次序发生了变化，则该算法是不稳定的。

结论:
1.关键词相同的数据元素将保持原有位置不变，所以该算法是稳定的。
2.时间复杂度:正序的时候[1,2,3,4,5]是O(N),逆序时O(N^2)最坏;整体:O(N^(1-2))
3.空间复杂度:常数阶O(1)
建议:插入排序适用于已经有部分数据已经排好，并且排好的部分越大越好。一般在输入规模大于1000的场合下不建议使用插入排序

i++和i--在Go语言中是语句，不是表达式，因此不能赋值给另外的变量。此外没有++i和--i
*/
func main() {
	nums := []int{4, 3, 5, 1, 2, 9}
	nums1 := []int{4, 3, 5, 1, 2, 9}
	log.Println("插入排序-哨兵法:", InsertSort(nums))
	log.Println("插入排序-交换法:", InsertSort1(nums1))
}

// InsertSort 直接插入排序 最好时间复杂度O(n),最坏时间复杂度O(n^2),空间复杂度O(1)
// 稳定排序算法
func InsertSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		temp := nums[i]

		// j代表了第一个比temp小的值的位置
		var j = i - 1
		// 此处控制是从小到大排还是从大到小排
		for j >= 0 && nums[j] > temp {
			// 不停的推动占位
			nums[j+1] = nums[j]
			j--
		}

		// j+1自然就是temp应该排序的位置
		nums[j+1] = temp
	}

	return nums
}

// InsertSort1 交换法
func InsertSort1(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0 && nums[j] > nums[j+1]; j-- {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
	}

	return nums
}
