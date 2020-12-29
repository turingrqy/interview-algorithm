package other

import (
	"fmt"
	"math"
	"renqiyang/interview/tree"
	"strconv"
)

//两数相加之和
func GetSumIndex (arr []int, target int)  {
	indexMap := make(map[int]int)
	for k,v :=range arr {
		indexMap[v] = k
	}

	for k,v :=range arr {
		coTarget := target-v
		if i1,ok := indexMap[coTarget];ok {
			fmt.Println(fmt.Sprintf("i1:%d,i2:%d", k,i1))
		}
	}
}

//s := arr[startIndex:endIndex]
//将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
//最长回文子串

//动态规划 如果aba是 cabac也是
//s[i+1,j-1]是回文的话 如果 s[i]==s[j] name s[i,j]也是
func GetLongestPalindrome (s string) {
	res := make([][]bool,0,len(s))
	ans := ""
	//l 限制子串的长度 l+1 是当前子串的长度
	for l:=0; l<len(s); l++ {
		for i:=0;i+l < len(s);i++ {
			//i和j代表不同长度子串的左边界和又边界
			j := i+l
			if l==0 {
				res[i][j]= true
			} else if l == 1 {
				res[i][j]= true
			} else {
				if s[i] == s[j] {
					res[i][j] = res[i-1][j-1]
				}
			}

			if res[i][j] && l+1 > len(ans) {
				ans = s[i:i+l+1]
			}
		}
	}
	fmt.Println(fmt.Sprintf("longest Palindrome= %s", ans))
}

//是否是回文数字

//最长非重复子串的长度
//滑动窗口

func GetNodupSubStrLen(s string) int {
	dupMap := make(map[byte]bool)
	maxLen := 0
	right :=0
	for left:=0; left< len(s);left++ {
		for ;right < len(s);right++ {
			if _,ok:=dupMap[s[right]];ok {
				break
			} else {
				dupMap[s[right]] = true
			}

		}
		if right-left > maxLen {
			maxLen = right-left
		}
		if right >= len(s) {
			break
		}
		delete(dupMap,s[left])
	}
	return maxLen
}

//哪个容器的水最多,和最长非重复子串类似使用双指针法

func GetMostbigCup (arr []int) int {
	left := 0
	right := len(arr)-1
	maxContainer := 0
	for;left <= right; {
		height := 0
		if arr[left] > arr[right] {
			height = arr[left]
			right--
		} else {
			height = arr[right]
			left++
		}

		container :=  height * (right-left)
		if container > maxContainer {
			maxContainer = container
		}
	}
	return maxContainer
}


//爬楼梯多少种方法
func GetUpstairsWaysNum (k int) int {
	if k==0 {
		return 0
	}
	if k==1 {
		return 1
	}
	if k==2 {
		return 2
	}

	return GetUpstairsWaysNum(k-1) + GetUpstairsWaysNum(k-2)
}
//最大子序列和
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
var numMap = map[byte]bool{
	'1':true,
	'2':true,
	'3':true,
	'4':true,
	'5':true,
	'6':true,
	'7':true,
	'8':true,
	'9':true,
	'0':true,
}
//解码a2[a2[ab]]c
func DecodeStr(str string) string {

	newStr := ""

	for i:=0;i<len(str);i++ {
		if _,ok := numMap[str[i]];ok {
			stack := tree.Stack{}
			var num int64
			num,i = getNumbyStr(str,i)
			stack.Push(num)
			newsubStr := ""
			needPush := true
			for !stack.IsEmpty() {

				if _,ok := numMap[str[i]];ok {
					var repeatNum int64
					repeatNum,i = getNumbyStr(str,i)
					stack.Push(repeatNum)
				} else if str[i] == '[' {
					i++
					continue
				} else if str[i] == ']' {
					needPush = false
					value := stack.Pop()
					subSubStr := value.(string)
					if newsubStr != "" {
						subSubStr =subSubStr+newsubStr
					}


					value = stack.Pop()
					srepeatNum := value.(int64)
					newsubStr = ""
					for j:=0;j< int(srepeatNum);j ++ {
						newsubStr =subSubStr+newsubStr
					}
					i++
					suffix :=""
					suffix,i = getSubStr(str,i)
					newsubStr = newsubStr+suffix
				} else {
					var subStr string
					subStr,i = getSubStr(str,i)
					if needPush {
						stack.Push(subStr)
					} else {
						newsubStr += subStr
					}
				}
			}
			newStr +=newsubStr
		} else {
			newStr += string(str[i])
		}
	}
	return newStr
}

func getNumbyStr (str string, i int) (int64,int) {
	numStr := ""
	for;i<len(str);i++ {
		if _,ok := numMap[str[i]];ok {
			numStr += string(str[i])
		} else {
			break
		}
	}
	num,_:=strconv.ParseInt(numStr,10,64)
	return num,i
}
func getSubStr (str string, i int) (string,int) {
	numStr := ""
	for;i<len(str);i++ {
		if _,ok := numMap[str[i]];!ok && str[i] != '[' && str[i] != ']' {
			numStr += string(str[i])
		} else {
			break
		}
	}
	//num,_:=strconv.ParseInt(numStr,10,64)
	return numStr,i
}
//是否是回文数
func IsPalindromeNum (num int64) bool  {
	suffixNum := int64(0)
	prefixNum := num

	round := 0

	for ;prefixNum > suffixNum;round++ {
		prefixNum /=10
		suffixNum +=(prefixNum%10) * int64(math.Pow10(round))
	}

	if prefixNum == suffixNum {
		return true
	}

	if suffixNum/10 == prefixNum {
		return true
	}
	return false
}

//删除数组中的重复元素
func DeleteDupInArr (arr []int) {
	prev := 0
	next := prev+1
	for ; next < len(arr); {
		for ;next < len(arr)-1;  {
			if arr[next] == arr[prev] {
				next ++
			} else {
				break
			}
		}

		if next > len(arr) {
			break
		}
		dupNum := next-prev-1
		for ;next<len(arr);next++ {
			arr[next-dupNum] = arr[next]

		}
		prev++
		next = prev+1
	}
}

// b^0=b
// a^b^a = b^(a^a)=b^0=b
//数组中只出现一次的数字
func SingleNumInArr (arr []int) int {
	res := 0
	for i:=0; i< len(arr);i++ {
		res ^=arr[i]
	}
	return res
}
//全排列
//深度优先遍历 树
// 一个保存已选择数组的栈 path,搜索树的深度 depth，已经选择的map
func Permute(nums []int) *[][]int {
	path := []int{}
	depth := 0
	usedMap := map[int]bool{}
	res := [][]int{}
	permuteDfs(nums,path, depth, usedMap, &res)
	return &res
}

func permuteDfs (nums []int, path []int ,depth int, usedMap map[int]bool, res*[][]int) {
	if depth == len(nums) {
		tmp := []int{}
		tmp = append(tmp,  path...)
		*res = append(*res, tmp)
		return
	}
	//这块也是可以优化的
	for i:=0; i< len(nums);i++ {
		if _,ok := usedMap[nums[i]]; ok && usedMap[nums[i]] {
			continue
		}
		path = append(path,nums[i])
		usedMap[nums[i]] = true
		permuteDfs(nums, path,depth+1,usedMap,res)
		usedMap[nums[i]] = false
		path = path[:len(path)-1]
	}
}

//给定不同的硬币种类 选择最少的硬币数量组合为需要的钱
//排列组合搜索查找找路径都应该考虑树形图解法这个和全排列是类似的 也是常说的回溯法
//最重要的几个点是要有 1.递归终止的条件 2.当前遍历的深度
/*输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1*/
func CoinChangeNormal(coins []int, amount int) int {
	minNum := math.MaxInt32
	CoinChangeDFS(coins, amount, 0, &minNum)
	return minNum
}

func CoinChangeDFS (coins []int, amount int, depth int, minDepth *int) {
	if amount == 0 {
		if *minDepth > depth {
			*minDepth = depth
		}
		return
	} else if amount < 0 {
		return
	}

	for i := 0; i< len(coins); i++ {
		CoinChangeDFS(coins, amount-coins[i], depth + 1, minDepth)
	}
}

//从上面看我们深度遍历中有很多无用的遍历，我们可以事先记录 n 块钱最少需要的钱币的数目 memo 记录n元钱 最少需要几步
//自顶向下 记忆迭代
func CoinChangeInMemo(coins []int, amount int) int {
	memo := make([]int, amount)
	return CoinChangeDFSInMemo(coins, amount, &memo)
}
func CoinChangeDFSInMemo (coins []int, amount int, memo *[]int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}
	if (*memo)[amount-1] != 0 {
		return (*memo)[amount-1]
	}

	var minStep = math.MaxInt32
	for i := 0; i < len(coins);i++ {
		res := CoinChangeDFSInMemo(coins,amount-coins[i], memo)
		if res >=0 && res < minStep {
			minStep= res+1
		}
	}
	if minStep == math.MaxInt32 {
		(*memo)[amount-1] = -1
	} else {
		(*memo)[amount-1] = minStep
	}
	return (*memo)[amount-1]
}
//动态规划自底向上
func CoinChangeDp(coins []int, amount int) int {
	memo := make([]int, amount+1)
	memo[0] = 0

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


/*输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。*/
func CatchRainNormal (arr []int) int {
	res := 0
	for i:=0; i< len(arr); i++ {
		maxLeft := 0
		maxRight := 0
		for left := i-1;left>=0;left-- {
			if arr[left] > arr[i] && arr[left] > maxLeft {
				maxLeft = arr[left]
			}
		}

		for right := i+1; right<len(arr); right++ {
			if arr[right] > arr[i] && arr[right] > maxRight {
				maxRight = arr[right]
			}
		}
		var height int
		if maxRight< maxLeft {
			height = maxRight
		} else {
			height = maxLeft
		}
		if height == 0 {
			continue
		}
		res += height-arr[i]
	}
	return res
}

func CatchRainQ (arr[]int) int {
	res := 0
	maxLeft := []int{}
	maxRight := make([]int,0,len(arr))
	maxLeft[0] = 0
	maxRight[len(arr)-1] = 0
	for i:=1;i<len(arr)-1;i++ {
		if arr[i] > maxLeft[i-1] {
			maxLeft = append(maxLeft, arr[i])
		} else {
			maxLeft = append(maxLeft, maxLeft[i-1])
		}
	}

	for j:= len(arr)-2;j>=0;j-- {
		if arr[j] > maxRight[j+1] {
			maxRight[j] = arr[j]
		} else {
			maxRight[j] = maxRight[j+1]
		}
	}

	for i:=0; i< len(arr); i++ {
		if maxLeft[i] < maxRight[i] && maxLeft[i] !=0 {
			res+=maxLeft[i]-arr[i]
		}else if maxRight[i] <= maxLeft[i] && maxRight[i] !=0 {
			res+=maxRight[i]-arr[i]
		}
	}
	return res
}

//买股票
func buyStokToMax(arr []int) int {
	minPrice := 0
	maxProfit := 0
	for i:= 0; i< len(arr); i++ {
		if arr[i] < minPrice {
			minPrice =arr[i]
		}
		Profit := arr[i]-minPrice
		if Profit > maxProfit {
			maxProfit =Profit
		}
	}
	return maxProfit
}

//获取最长递增子序列的长度 递归
func GetLongestUpNum(arr []int) int {
	return GetLongestUpNumNormal(arr, len(arr)-1)
}

func GetLongestUpNumNormal (arr []int, i int) int {
	max := 1
	for j:=0;j<i;j++ {
		if arr[i] > arr[j] {
			res := GetLongestUpNumNormal(arr, j)
			if res +1 >max {
				max = res+1
			}
		}
	}
	return max
}
func GetLongestUpInMemo(arr []int, i int, dp *[]int) int {
	if i == 0  {
		return 0
	}
	if (*dp)[i] > 0 {
		return (*dp)[i]
	}

	max := 1
	for j:=0;j<i;j++ {
		if arr[i] > arr[j] {
			res := GetLongestUpNumNormal(arr, j)
			if res >max {
				max = res+1
			}
		}

	}
	(*dp)[i] = max
	return (*dp)[i]
}




