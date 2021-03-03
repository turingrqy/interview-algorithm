package recursion_dynamic

import (
	"math"
)

//打家劫舍
/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func RobMax(arr []int) int {
	dpArr := make([]int,len(arr),len(arr))
	dpArr[0] = arr[0]
	if dpArr[0] > dpArr[1] {
		dpArr[1] = dpArr[0]
	} else {
		dpArr[1] = dpArr[1]
	}
	for j:=2;j<len(arr)-1;j++ {
		if arr[j] + dpArr[j-2] > dpArr[j-1] {
			dpArr[j] = arr[j] + dpArr[j-2]
		} else {
			dpArr[j] = dpArr[j-1]
		}
	}

	return dpArr[len(arr)-1]
}

//最大子数组和
func GetMaxSumInArray(arr []int64) int64 {
	maxSum :=arr[0]
	sum := int64(0)
	for i:=0;i< len(arr);i++ {
		sum += arr[i]
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}
//乘积最大子数组 乘积
func GetMaxMulti(arr []int) int {
	max := arr[0]
	min := arr[0]
	maxRes := arr[0]
	for i:=1; i< len(arr);i++ {
		multiMin := min * arr[i]
		multiMax := max * arr[i]
		if multiMax > multiMin {
			max = multiMax
			min = multiMin
		} else {
			max = multiMin
			min = multiMax
		}

		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
		if max > maxRes {
			maxRes = max
		}
	}
	return maxRes
}

func GetMaxLengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] =1
	for i:=1;i<len(nums);i++ {
		max := 1
		for j:= 0;j<i;j++ {
			if nums[i]>nums[j] {
				tmpMax := dp[j] + 1
				if tmpMax > max {
					max = tmpMax
				}
			}
		}
		dp[i] = max
	}
	max := 0
	for i:=0;i<len(dp);i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
func GetNumberofLIS (nums []int) int {
	dp := make([]int, len(nums))
	dp[0] =1
	counts := make([]int, len(nums))
	counts[0]=1
	for i:=1;i<len(nums);i++ {
		max := 1
		counts[i] = 1
		for j:= 0;j<i;j++ {
			if nums[i]>nums[j] {
				tmpMax := dp[j] + 1
				if tmpMax > max {
					max = tmpMax
					counts[i] = counts[j]
				} else if tmpMax==max {
					counts[i] += counts[j]
				}
			}
		}
		dp[i] = max
	}
	maxCount := 0
	maxLen := 0
	for i:=0;i<len(dp);i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
			maxCount= counts[i]
		} else if dp[i] == maxLen {
			maxCount+= counts[i]
		}
	}
	return maxCount
}

func GetLongetLIS (nums []int) [][]int {
	dp := make([]int,len(nums))
	lisArr := make([][][]int, len(nums))
	lisArr[0] = make([][]int, 1)
	lisArr[0][0] = []int{nums[0]}

	for i:= 0; i< len(nums);i++ {
		max := 1
		var maxIndex []int
		for j:= 0; j<i;j++ {
			if nums[i] > nums[j] {
				tmpMax := dp[j]+1
				if tmpMax > max {
					max = tmpMax
					maxIndex = []int{}
					maxIndex = append(maxIndex,j)
				} else if tmpMax== max{
					maxIndex = append(maxIndex,j)
				}
			}
		}
		dp[i] = max
		if len(maxIndex) > 0 {
			lisArr[i] = [][]int{}
			for _,index := range maxIndex {
				for _,arr := range lisArr[index] {
					tmpArr := []int{}
					tmpArr = append(tmpArr, arr...)
					tmpArr = append(tmpArr, nums[i])
					lisArr[i] = append(lisArr[i],tmpArr)
				}
			}
		} else {
			lisArr[i] = make([][]int, 1)
			lisArr[i][0] = []int{nums[i]}
		}
	}
	maxLen:=0
	var maxIndex []int
	for i:=0;i<len(lisArr);i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
			maxIndex = []int{}
			maxIndex = append(maxIndex, i)
		} else {
			maxIndex = append(maxIndex, i)
		}
	}
	res := [][]int{}
	for _,index :=range maxIndex {
		res = append(res,lisArr[index]...)
	}
	return res
}

// 买股票最佳时机1 只能买一次
func StokmaxProfit(arr[] int) {
	maxProfit := 0
	minPrice := math.MaxInt32
	for i:=0;i<len(arr);i++ {
		if arr[i] < minPrice {
			minPrice = arr[i]
		}
		profit := arr[i]-minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
}
// 买股票最佳时机2 能买多次但买之前手里的股票必须卖出 二维数组dp[i][j] i 第i天 j 0，1 是否持有股票 值为收益
func MultiByStokMaxProfit(arr[] int) int {
	dp := make([][]int, len(arr))
	for i:=0;i<len(dp);i++ {
		dp[i] = make([]int,2)
	}

	dp[0][0] = 0
	dp[0][1] = -arr[0]

	for i:= 1;i<len(arr);i++ {
		//没有持有股票 要不是前一天就没持有，要么昨天今天卖掉
		dp[i][0] = max(dp[i-1][0],dp[i-1][1]+arr[i])
		//没有持有股票 要不是前一天就持有，要么是昨天没有今天刚买
		dp[i][1] = max(dp[i-1][1],dp[i-1][0]-arr[i])
	}
	return dp[len(arr)-1][0]
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

