package binary_search

import (
	"fmt"
	"math"
	"renqiyang/interview/heap_sort"
)
//二维排序数组查找1 下一行第一个元素大于上一行最后一个元素
//将二维数组看成一维排序数组，line =index/3 col index%3-1
func FindInTwoDimensionStrictSorted(arr[][]int, target int) (int,int) {
	rows := len(arr)
	col := len(arr[0])
	low := 0
	high := rows*col -1
	for low <= high {
		mid := (low + high)/2
		midRow := mid/col
		midCol := mid%col

		if arr[midRow][midCol] > target {
			high = mid-1
		} else if arr[midRow][midCol] == target {
			return midRow,midCol
		} else {
			low = mid+1
		}
	}
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
	ans := int(0)
	for low <= high {
		mid := (low + high)/2
		if arr[mid] > arr[high] {
			ans = mid + 1
			low = mid + 1
		}else if arr[mid] < arr[high] {
			high = mid
		} else {
			high --
		}
	}
	return ans
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
	if k> len(arr1) + len(arr2)  {
		return 0
	}
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
		if index1 + half <= len(arr1) {
			newIndex1 = index1 + half-1
		} else {
			newIndex1 = len(arr1)-1
		}
		if index2 + half <= len(arr2) {
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
//寻找峰值 相邻元素不相等
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
有序数组，有2N+1个数，其中N个数成对出现，仅1个数单独出现，找出那个单独出现的数.,时间复杂度

1，1，2，2，3，4，4，5，5，6，6，7，7
*/
// 也是二分只不过要计算下先检查 是不是和两边的数相同

func GetNoDupNumInSortedDoubleArr (arr []int) int {
	low := 0
	high := len(arr)-1

	if len(arr) == 1 {
		return arr[0]
	}
	for low <= high {
		mid := (low + high)/2
		if mid + 1 < len(arr) && arr[mid] == arr[mid+1] {
			itemNum := high-mid+1
			if itemNum % 2 == 0 {
				high = mid -1
			} else {
				low = mid +2
			}
			continue
		}
		if mid -1 >= 0 && arr[mid-1] == arr[mid] {
			itemNum := mid-low+1
			if itemNum % 2 == 0 {
				low = mid +1
			} else {
				high = mid-2
			}
			continue
		}

		if mid -1 < 0 && arr[mid] != arr[mid+1] {
			return arr[mid]
		}

		if mid +1 >= len(arr) && arr[mid] != arr[mid-1] {
			return arr[mid]
		}
		if mid -1>=0 && mid + 1 < len(arr) && arr[mid] != arr[mid-1] && arr[mid] != arr[mid+1] {
			return arr[mid]
		}
	}
	return -1
}

/*
给定一个配排序的先排序整数数组找到3个数的和最接近于target 3数之和 最接近
*/
func GetClosestThreeNumSum(arr []int64, target int64) (a,b,c,sum int64) {
	//1.先排序
	//2.遍历选择第一个数据a
	heap_sort.HeapSort(arr)
	var aBest,bBest,cBest,sumBest int64
	sumBest = math.MinInt64
	//ans := int(0)
	for i:= 0; i< len(arr)-2;i++ {
		k := i+1
		j := len(arr)-1
		var needReduce,needPlus bool

		sum := int64(0)
		for k<j {

			sum = arr[i] + arr[k] + arr[j]
			if  sum < target {
				//ans = k
				k++
				needReduce = true
				needPlus = false
			} else if sum > target {
				//ans = j
				j--
				needPlus = true
				needReduce = false
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

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/3sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func GetThreeNumSumEqTarget (arr []int64, target int64) {
	heap_sort.HeapSort(arr)
	for first := 0; first < len(arr)-2; first++ {
		if first > 0 && arr[first] == arr[first-1] {
			continue
		}
		coTarget := target - arr[first]
		third := len(arr)-1
		second := first+1
		for  ;second < third; second++ {
			if second > first + 1 && arr[second] == arr[second-1] {
				continue
			}

			for second < third && arr[second] + arr[third] > coTarget {
				third--
			}

			if second == third {
				break
			}
			if arr[second] + arr[third] == coTarget {
				println(fmt.Sprintf("first=%d, second=%d, third=%d", arr[first], arr[second], arr[third]))
			}
		}
	}
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

/*找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
置换
*/

func GetDupInArr(nums []int) int {
	for i:=0;i<len(nums);i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i],nums[nums[i]] = nums[nums[i]],nums[i]
		}
	}
	return -1
}


/*
34. 在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。
右边界就是找第一个大于target的下标
左边界是找第一个大于等于target的下标
*/
func BinarySearchRange (nums []int, target int, lower bool) int {
	low := 0
	high := len(nums)-1
	ans := len(nums)
	mid := (high+low)/2
	for low <= high {
		if nums[mid] > target || (lower && nums[mid] >= target) {
			ans = mid
			high = mid -1
		} else {
			low = mid+1
		}
	}
	return ans
}

func GetEqualRangeInSortedArr (nums []int, target int) (int, int) {
	leftIdx := BinarySearchRange (nums, target, true)
	rightIdx := BinarySearchRange (nums, target, false)-1

	if leftIdx <= rightIdx && rightIdx < len(nums) && nums[leftIdx] == target && nums[rightIdx] == target {
		return leftIdx, rightIdx
	} else {
		return -1, -1
	}
}







