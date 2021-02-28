package binary_search

import (
	"math"
	"renqiyang/interview/heap_sort"
)
//二维排序数组查找1 下一行第一个元素大于上一行最后一个元素
//将二维数组看成一维排序数组，line =index/3 col index%3-1
func FindInTwoDimensionStrictSorted(arr[][]int, target int) (int,int) {
	return -1,-1
}
//二维排序数组查找2
func FindInTwoDimensionSorted(arr[][]int, target int) (int,int) {
	startRow := len(arr)-1
	startCol := 0
	maxCol := len(arr[0])-1
	for startRow >=0 && startCol <=maxCol {
		if arr[startRow][startCol] > target {
			startRow--
		} else if arr[startRow][startCol] < target {
			startCol++
		} else {
			return startRow,startCol
		}
	}
	return -1,-1
}

//排序数组要想到二分查找

func BinarySearch (arr []int64, lookFor int64) int {
	low, high := 0, len(arr)-1

	for low <= high  {
		mid := (low+high)/2

		if arr[mid] > lookFor {
			high = mid-1
		} else if arr[mid] < lookFor {
			low = mid+1

		} else {
			return mid
		}

	}
	return -1
}

func HashBinarySearch (arr []int64, node int64) int64 {
	low := 0
	high := len(arr) -1

	index := 0
	for low <= high {
		mid := (low+high)/2
		if arr[mid] == node{
			return arr[mid]
		}

		if arr[mid] < node {
			low = mid +1
		} else if arr[mid] > node {
			index = mid
			high = mid-1
		}
	}

	return arr[index]
}

//旋转数组查找
func BinarySearchRotate (arr []int64, lookFor int64) int {
	low := 0
	high := len(arr)-1

	for low <= high {
		mid := (low + high)/2
		if arr[mid] == lookFor{
			return mid
		}
		if arr[mid] >= arr[low] {

			if lookFor >= arr[low] && arr[mid] > lookFor {
				high = mid -1
			} else{
				low = mid +1
			}
		} else {
			if lookFor <= arr[high] && arr[mid] < lookFor {
				low = mid +1
			} else {
				high = mid -1
			}
		}
	}
	return -1
}

//旋转数组最小值
func BinarySearchMinInRotate (arr []int64) int {
	low := 0
	high := len(arr)-1
	midIndex := low
	for arr[low] >= arr[high] {
		if high-low == 1 {
			midIndex = high
			break
		}

		midIndex = (low+high)/2
		if arr[midIndex] == arr[low] && arr[midIndex] == arr[high] {
			//顺序查找
			GetMinInRatateOrder(arr, low, high)
		}

		if arr[midIndex] >= arr[low] {
			low = midIndex
		} else if arr[midIndex] <= arr[high] {
			high = midIndex
		}

	}
	return midIndex
}

func GetMinInRatateOrder (arr []int64, low, high int) int {

	index := low
	min := arr[low]
	for i:=index +1; i<= high;i++ {
		if min > arr[i] {
			index = i
			min = arr[i]
		}
	}

	return index
}

func getDpBinarySearch(arr []int64, target int64) int {
	low := 0;
	high := len(arr)-1

	for low <= high {
		if target < arr[low] {
			return low
		}
		if target > arr[high] {
			return high +1
		}

		mid := (low+high)/2

		if arr[mid] == target {
			return -1
		}
		if arr[mid] > target {
			high = mid -1
		} else {
			low = mid + 1
		}
	}
	return -1
}
// 最长递增子序列顺序和位置要求和子数组相同
func GetLongestUpSubArr (arr []int64) [][]int64 {
	dp := make([][]int64,len(arr))
	en := []int64{}

	dp = append(dp, []int64{arr[0]})
	en = append(en, arr[0])
	for key, val := range arr {
		p := getDpBinarySearch(en, val)
		if p >= 0 {
			en = append(en[0:p], val)
			tmp := []int64{}
			dp[key] = append(tmp, en...)
		}
	}

	return dp
}
//两个递增数组找到两个数组中的中位数
//log(m+n) log 一般都要想到二分查找
func GetMidIntwosortedArr(arr1 []int64, arr2 []int64) float64 {
	totalLen := len(arr1)+len(arr2)
	if totalLen%2 == 0 {
		return (float64(GetKthNumInTwoArr(arr1,arr2, totalLen/2)) + float64(GetKthNumInTwoArr(arr1,arr2, totalLen/2+1)))/2.0
	} else {
		return GetKthNumInTwoArr(arr1,arr2, totalLen/2+1)
	}
}

func GetKthNumInTwoArr(arr1 []int64, arr2 []int64, k int) float64 {
	index1, index2 := 0, 0
	for {
		if index1 == len(arr1) {
			return float64(arr2[index2+k-1])
		}
		if index2 == len(arr2) {
			return float64(arr1[index1+k-1])
		}
		if k == 1 {
			if arr1[index1] < arr2[index2] {
				return float64(arr1[index1])
			} else {
				return float64(arr2[index2])
			}
		}
		var newIndex1,newIndex2 int
		half := k/2
		if index1 + half < len(arr1) {
			newIndex1 = index1 + half-1
		} else {
			newIndex1 = len(arr1)-1
		}
		if index2 + half < len(arr2) {
			newIndex2 = index2 + half-1
		} else {
			newIndex2 = len(arr2)-1
		}

		if arr1[newIndex1] < arr2[newIndex2] {
			k -= (newIndex1-index1 +1)
			index1 = newIndex1+1
		} else {
			k -= (newIndex2-index2 +1)
			index2 = newIndex2+1
		}
	}

	return 0
}
//寻找峰值
func FindPeek(arr []int) int {
	left, right := 0, len(arr)-1

	for ;left <= right; {
		half := (left + right)/2
		if half+1 <= len(arr)-1 {
			if arr[half+1] > arr[half] {
				left = half+1
			} else {
				if half - 1 >=0  {
					if arr[half-1] < arr[half] {
						return half
					} else {
						right = half-1
					}
				} else {
					return half
				}
			}
		} else {
			if half-1 >= 0 {
				if arr[half] > arr[half-1] {
					return half
				} else {
					right = half-1
				}
			} else {
				return half
			}
		}
	}
	return -1
}
// 查找唯一重复的数字
/*给定一个包含n + 1 个整数的数组nums ，其数字都在 1 到 n之间（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。
二分法 + 抽屉原理
*/
func GetOnceDupNumInArr (arr []int ) int {
	left := 1
	right := len(arr)-1
	var ans int
	for left <= right {
		mid :=  (left + right)/2
		cnt := 0
		for _,v:=range arr {
			if v <= mid {
				cnt ++
			}
		}
		if cnt <= mid {
			left = mid + 1
		} else {
			right = mid-1
			ans = mid
		}
	}
	return ans
}

/*
给定一个整数数组找到3个数的和最接近于target
*/
func GetClosestThreeNumSum(arr []int64, target int64) (a,b,c,sum int64) {
	//1.先排序
	//2.遍历选择第一个数据a
	heap_sort.HeapSort(arr)
	var aBest,bBest,cBest,sumBest int64
	sumBest = math.MinInt64
	for i:= 0; i< len(arr)-2;i++ {
		k := i+1
		j := len(arr)-1
		needReduce := false
		needPlus := false
		sum := int64(0)
		for k<j {
			sum = arr[i] + arr[k] + arr[j]
			if  sum < target {
				k++
				needReduce = true
			} else if sum > target {
				j--
				needPlus = true
			} else {
				return arr[i],arr[k],arr[j],target
			}
		}
		if needReduce {
			k--
		}
		if needPlus {
			j++
		}
		if getDistance(sumBest,target) > getDistance (sum,target) {
			aBest,bBest,cBest,sumBest = arr[i],arr[k],arr[j],sum
		}
	}
	return aBest,bBest,cBest,sumBest
}
func getDistance (a, target int64) int64 {
	d := target-a
	if d<0 {
		return -d
	}
	return d
}

// x的平方根 二分查找
func GetSqrt(x int) int {
	low := 1
	high := x
	var ans int
	for low<=high {
		mid := (low+high)/2
		if mid * mid > x {
			high = mid-1
		} else if mid * mid < x  {
			ans = mid
			low = mid+1

		} else {
			return mid
		}
	}
	return ans
}

//完全2叉树几点数量，直接遍历二叉树是o(n)的操作
//使用的完全二叉树的性质和二分查找计算节点个数






