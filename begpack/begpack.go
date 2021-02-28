package begpack
// 不能重复使用硬币 最少使用硬币 要凑成的数是背包，硬币候选池是 物品，所以遍历背包的时候后要 从后向前
func CoinChangeDpNoDup(coins []int, amount int) int {
	tmpArr := make([]int, amount+1)
	tmpArr[0] = 0
	for i:=0;i<len(coins);i++ {
		for j:= amount;j>0;j-- {
			tmpAmount := j - coins[i]
			if tmpAmount < 0 {
				continue
			}
			////min (tmpArr[j],tmpArr[j-coins[i]]+1)
			if tmpAmount == 0 {
				tmpArr[j] = 1
			} else if tmpArr[j-coins[i]] > 0 {
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
	//可以重复 背包可以从左到右遍历
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

//求和等于target的组合 候选集数据可重复
func GetCombineSumEqTargetDp (arr []int, target int) [][]int {
	dpArr := make([][][]int, target+1)
	dpArr[0] = [][]int{}
	dpArr[0] = append(dpArr[0],[]int{})
	for i:=0;i<len(arr);i++ {
		for j:= 0;j<=target;j++ {
			tmp := j-arr[i]
			if tmp >=0 && len(dpArr[tmp]) > 0 {

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
