package begpack

import (
	"fmt"
	"sort"
)

// 不能重复使用硬币
// 背包倒序是防止一个硬币被放进去多次，coin
func CoinChangeDpNoDup(coins []int, amount int) int {
	tmpArr := make([]int, amount+1)
	tmpArr[0] = 0
	for i:=0;i<len(coins);i++ {
		for j:= amount;j>0;j-- {
			tmpAmount := j - coins[i]

			if tmpAmount == 0 {
				tmpArr[j] = 1
			} else if tmpAmount > 0 && tmpArr[j-coins[i]] > 0 {
				tmpArr[j] = tmpArr[j-coins[i]]+1
			}
		}
	}
	if tmpArr[amount] == 0 {
		return -1
	}
	return tmpArr[amount]
}

//动态规划自底向上 可以重复使用硬币 最少使用硬币 完全背包
func CoinChangeDp(coins []int, amount int) int {
	memo := make([]int, amount+1)
	memo[0] = 0
	//可以重复
	for i:=1; i<= amount;i++ {
		for j:=0; j< len(coins);j++ {
			tmpAmount := i-coins[j]
			if tmpAmount == 0 {
				memo[i] = 1
				break
			}else if  tmpAmount >0 && memo[tmpAmount]>0  {
				memo[i] = memo[tmpAmount]+1
			}
		}
	}

	return memo[amount]
}

//组合数 换硬币的组合数 不能重复使用 先遍历物品是组合数，先遍历钱包是排列数，先遍历物品遍历过得物品不会再被遍历到，先遍历背包所有的coins
// 都会遍历一遍
func GetCoinChangeCombineNoDup (coins []int, target int) int {
	dpArr := make([]int,target+1)
	dpArr[0] = 1
	for i:=0;i<len(coins);i++ {
		for j:= target;j>0;j-- {
			tmp := j-coins[i]
			if tmp >=0 {
				//dp[j]new = dp[j]old(不考虑coins[i]的情况下的和为j的组合数) + 考虑coins[i]的组合数
				dpArr[j] = dpArr[j]+dpArr[j-coins[i]]
			}
		}
	}
	return dpArr[target]
}
//组合数可以重复使用 如果是排列数是先遍历背包在遍历物品
func GetCoinChangeCombine (coins []int, target int) int {
	dpArr := make([]int,target+1)
	dpArr[0] = 1
	for i:=0;i<len(coins);i++ {
		for j:= 0;j<=target;j++ {
			tmp := j-coins[i]
			if tmp >=0 {
				//dp[j]new = dp[j]old(不考虑coins[i]的情况下的和为j的组合数) + 考虑coins[i]的组合数
				dpArr[j] = dpArr[j]+dpArr[j-coins[i]]
			}
		}
	}
	return dpArr[target]
}

//求和等于target的组合 元素可重复使用
func GetCombineSumEqTargetDp (arr []int, target int) [][]int {
	dpArr := make([][][]int, target+1)
	dpArr[0] = [][]int{}
	dpArr[0] = append(dpArr[0],[]int{})
	for i:=0;i<len(arr);i++ {
		for j:= 0;j<=target;j++ {
			tmp := j-arr[i]
			if tmp >=0{

				for k:=0;k<len(dpArr[tmp]);k++ {
					tmpArr := []int{}
					tmpArr = append(tmpArr,dpArr[tmp][k]...)
					tmpArr = append(tmpArr,arr[i])
					dpArr[j] = append(dpArr[j],tmpArr)
				}

			}
		}
	}
	return dpArr[target]
}

func GetCombineSumNoDupEqTargetDp (arr []int, target int) [][]int {
	dpArr := make([][][]int, target+1)
	dpArr[0] = [][]int{}
	dpArr[0] = append(dpArr[0],[]int{})
	for i:= 0; i<len(arr); i++ {
		for j:=target;j>0;j--{
			tmp := j-arr[i]
			if tmp >=0 && len(dpArr[tmp]) > 0{
				for k:=0; k< len(dpArr[tmp]); k++ {
					tmpArr := make([]int, 0)
					tmpArr = append(tmpArr, dpArr[tmp][k]...)
					tmpArr = append(tmpArr, arr[i])
					dpArr[j] = append(dpArr[j], tmpArr)
				}
			} else {

			}
		}
	}
	return dpArr[target]
}

//给定一个元素不重复的数组，找出所有和为target的组合
/*所有数字（包括 target）都是正整数。和选硬币是一样的 这个就是求所有的组合
解集不能包含重复的组合，求所有组合只能是递归了 求组合数可以用背包 可以重复选择*/
//可以每次都选择是使用下一个还是当前的方法 不可重复选择元素,数组中也没有重复元素



func FindCombineSumEqTargetNoDup1 (arr[]int, target int) {
	tmpArr := []int{}

	FindCombineSumEqTargetNoDupDFS1(arr ,target, tmpArr, 0)
}

func FindCombineSumEqTargetNoDupDFS1 (arr[]int, target int, tmpArr []int, idx int) {
	if target == 0 {
		println(fmt.Sprintf("tmpArr=%v",tmpArr))
		return
	}
	if idx == len(arr) {
		return
	}

	//不选当前的元素
	FindCombineSumEqTargetNoDupDFS1(arr, target, tmpArr, idx+1)
	if target-arr[idx] >= 0 {
		tmpArr = append(tmpArr, arr[idx])
		FindCombineSumEqTargetNoDupDFS1(arr, target-arr[idx], tmpArr, idx+1)
	}
}

func FindCombineSumEqTargetNoDup2 (arr[]int, target int , k int) {
	tmpArr := []int{}

	FindCombineSumEqTargetNoDupDFS2(arr ,target, tmpArr, 0, k)
}

func FindCombineSumEqTargetNoDupDFS2 (arr[]int, target int, tmpArr []int, idx int , k int) {
	if len(tmpArr) == k && target ==0 {
		println(fmt.Sprintf("arr=%v", tmpArr))
		return
	}

	if len(tmpArr) > k || idx == len(arr) {
		return
	}
	FindCombineSumEqTargetNoDupDFS2 (arr, target, tmpArr, idx+1, k)
	if target-arr[idx] >= 0 {
		tmpArr = append(tmpArr, arr[idx])
		FindCombineSumEqTargetNoDupDFS2 (arr, target-arr[idx], tmpArr, idx+1, k)
	}
}


//给定一个元素不重复的数组，找出所有和为target的组合
/*所有数字（包括 target）都是正整数。和选硬币是一样的 这个就是求所有的组合
解集不能包含重复的组合，求所有组合只能是递归了 求组合数可以用背包 可以重复选择*/
//可以每次都选择是使用下一个还是当前的方法 可重复循选但是不能有负数
func FindCombineSumEqTarget (arr[]int, target int) {
	tmpArr := []int{}
	FindCombineSumEqTargetDFS(arr ,target,tmpArr, 0)
}

func FindCombineSumEqTargetDFS (arr[]int, target int, tmpArr []int, idx int) {
	if idx == len(arr) {
		return
	}
	if target == 0 {
		println(fmt.Sprintf("tmpArr=%v",tmpArr))
		return
	}
	//不选当前的元素
	FindCombineSumEqTargetDFS(arr, target, tmpArr, idx+1)
	if target-arr[idx] >= 0 {
		tmpArr = append(tmpArr, arr[idx])
		FindCombineSumEqTargetDFS(arr, target-arr[idx], tmpArr, idx)
	}
}

func FindCombineSumEqTargetNoDupInRArr (arr []int, target int) {
	sort.Sort(sort.IntSlice(arr))
	println(fmt.Sprintf("arr=%v", arr))
	FindCombineSumEqTargetNoDupInRArrDfs(arr, target, 0, []int{})
}

func FindCombineSumEqTargetNoDupInRArrDfs (arr []int, target int, idx int, path []int) {
	if target == 0 {
		println(fmt.Sprintf("FindCombineSumEqTargetNoDupInRArrDfs=%+v", path))
		return
	}

	if idx == len(arr) {
		return
	}


	for i:=idx; i < len(arr); i++ {
		if i > idx && arr[i] ==  arr[i-1] {
			continue
		}
		FindCombineSumEqTargetNoDupInRArrDfs(arr, target-arr[i], i+1, append(path, arr[i]))
	}
}

