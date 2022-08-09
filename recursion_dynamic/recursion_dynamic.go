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
	for j:=2;j<len(arr);j++ {
		if arr[j] + dpArr[j-2] > dpArr[j-1] {
			dpArr[j] = arr[j] + dpArr[j-2]
		} else {
			dpArr[j] = dpArr[j-1]
		}
	}

	return dpArr[len(arr)-1]
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

//最长连递增子序列 看位置关系
func GetMaxLengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] =1
	res := 0
	for i:=1; i<len(nums); i++ {
		dp[i] = 1
		for j := 0; j<i; j++ {
			if nums[i] > nums[j] {
				tmp := dp[j]+1
				if tmp > dp[i] {
					dp[i] = tmp
				}
			}
		}
		res = int(math.Max(float64(dp[i]), float64(res)))
	}
	return res
}
func GetNumberofLIS (nums []int) int {
	dp := make([]int, len(nums))
	dp[0] =1
	counts := make([]int, len(nums))
	counts[0]=1
	maxLen := 1
	countRes := 1
	for i:=0; i<len(nums);i++ {
		dp[i] = 1
		counts[i] = 1
		for j:=0; j<i; j++ {
			if nums[i] > nums[j] {
				tmp := dp[j] + 1
				if tmp > dp[i] {
					dp[i] = tmp
					counts[i] = counts[j]
				} else if tmp == dp[i] {
					counts[i] += counts[j]
				}
			}
		}
		if dp[i] == maxLen {
			countRes += counts[i]
		} else if dp[i] > maxLen {
			maxLen = dp[i]
			countRes = counts[i]
		}
	}
	return countRes
}

func GetLongetLIS (nums []int) [][]int {
	dp := make([]int,len(nums))
	list := make([][][]int, 0, len(nums))
	maxLen := 1
	dp[0] = 1
	list[0] = [][]int{}
	list[0] = append(list[0], []int{nums[0]})
	res := make([][]int, 0, 0)

	for i:=0; i<len(nums);i++ {
		dp[i]=1
		list[i] = [][]int{{nums[i]}}

		for j:=0; j<i; j++ {
			if nums[i] > nums[j] {
				tmp := dp[j]+1
				if tmp > dp[i] {
					dp[i] = tmp
					list[i] = [][]int{}
					for k := 0; k<len(list[j]); j++ {
						var tmpArr []int
						tmpArr = append(tmpArr, list[j][k]... )
						tmpArr = append(tmpArr, nums[i])
						list[i] = append(list[i], tmpArr)
					}
				} else if tmp > dp[i] {
					for k := 0; k<len(list[j]); j++ {
						var tmpArr []int
						tmpArr = append(tmpArr, list[j][k]... )
						tmpArr = append(tmpArr, nums[i])
						list[i] = append(list[i], tmpArr)
					}
				}
			}
		}
		if dp[i] > maxLen {
			res = make([][]int, 0, 0)
			for _, v :=  range list[i] {
				var tmpArr []int
				tmpArr = append(tmpArr, v...)
				res = append(res, tmpArr)
			}
		} else if dp[i] == maxLen {
			for _, v :=  range list[i] {
				var tmpArr []int
				tmpArr = append(tmpArr, v...)
				res = append(res, tmpArr)
			}
		}
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


/*最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
遇到最优解需要想到动态规划
dp[i][j] 代表[0:i] text2[0:j] 最长长度
*/

func GetLongestCommonSubsequence (text1, text2 string) int {
	dp := make ([][]int, len(text1)+1)
	for i:=0; i<len(text1)+1; i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i:=1; i<=len(text1); i++ {
		for j:=1; j<=len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1]+1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

/*
91. 解码方法
一条包含字母 A-Z 的消息通过以下映射进行了 编码 ：

'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
求最多能有多少解码方法
dp 大是因为 前两个字符也会有多种可能
 */

func DecodeAZ (text string) int {
	dp := make([]int, len(text) + 1)
	dp[0] = 1
	for i:= 1; i <= len(text); i++ {
		if text[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && text[i-2] != '0' && (text[i-1]-'0' + (text[i-2]-'0') * 10) <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(text)]
}

/*
最长重复子数组
718. 最长重复子数组
给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。 子数组是位置连续的
动态规划
 */
func GetLongestSubArr (nums1,nums2 []int) []int {
	dp := make([][]int, len(nums1)+1)

	for i:=0; i< len(nums1)+1; i++ {
		dp[i] = make([]int, len(nums2)+1)
	}

	maxEnd := 0
	maxLen := 0
	for i:=1; i<=len(nums1); i++ {
		for j:=1; j<=len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
					maxEnd = i
				}
			}
		}
	}
	println(maxEnd, maxLen)
	start := maxEnd-maxLen
	return nums1[start:maxEnd]
}






