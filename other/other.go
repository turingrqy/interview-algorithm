package other

import (
	"context"
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
	context.WithCancel(context.Background())
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
					res[i][j] = res[i+1][j-1]
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

func GetMaxLenOfNoDupStr (s string) int {
	max := 0
	indexMap := make(map[byte]int)
	left := 0
	right :=0
	for ;right < len(s);right++ {
		if index,ok := indexMap[s[right]];!ok {
			indexMap[s[right]] = right
		} else {
			for i:= left;i<=index;i++ {
				delete(indexMap,s[i])
			}
			left = index+1
			indexMap[s[right]] = right
		}
		len := right-left+1
		if len > max {
			max = len
		}
	}
	return max
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


//

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
func IsPalindromeNum (num int) bool  {
	prefix := num
	suffix := 0

	for ;suffix< prefix; {
		suffix = suffix * 10 +prefix%10
		prefix /=10
	}

	if prefix == suffix || suffix/10==prefix {
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
//相同取0，相异取1。（二进制）
func SingleNumInArr (arr []int) int {
	res := 0
	for i:=0; i< len(arr);i++ {
		res ^=arr[i]
	}
	return res
}
//全排列
//深度优先遍历 树
// 一个保存已选择数组的栈 path,搜索树的深度 depth，已经选择的map map 保证不能选已经选过的
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
//还有一种是求组合数



//给定一个元素不重复的数组，找出所有和为target的组合
/*所有数字（包括 target）都是正整数。和选硬币是一样的 这个就是求所有的组合
解集不能包含重复的组合，求所有组合只能是递归了 求组合数可以用背包*/
//可以每次都选择是使用下一个还是当前的方法
func FindCombineSumEqTarget (arr[]int, target int) [][]int {
	res := &[][]int{}
	tmpArr := []int{}
	FindCombineSumEqTargetDFS(arr,target,tmpArr, 0,res)
	return *res
}

func FindCombineSumEqTargetDFS (arr[]int, target int ,tmpArr []int, idx int, res *[][]int) {
	if target == 0 {
		//终止条件 找到了一个组合
		tmpRes := []int{}
		tmpRes = append(tmpRes, tmpArr...)
		*res = append(*res,tmpRes)
		return
	}
	if idx == len(arr) {
		//候选集被选完了
		return
	}

	/*这里并不是每次从头开始选了，从idx 初先选
	for i:=0;i<len(arr);i++ {
		tmpArr = append(tmpArr, arr[i])
		FindCombineSumEqTargetDFS (arr, target-arr[i], tmpArr, res)
		tmpArr = tmpArr[:len(tmpArr)-1]
	}*/
	//下一次从下一个数开始选
	// 不取当前的数 取下一个数
	FindCombineSumEqTargetDFS (arr, target, tmpArr, idx+1, res)
	if target-arr[idx] >=0 {
		tmpArr = append(tmpArr, arr[idx])
		//
		FindCombineSumEqTargetDFS (arr, target-arr[idx], tmpArr, idx, res)
		tmpArr = tmpArr[:len(tmpArr)-1]
	}
}

func Subsets(nums []int) [][]int {
	cur := []int{}
	res := [][]int{}
	SubSetsDFS(nums, 0,cur, &res)
	return res
}

func SubSetsDFS (nums []int, depth int, cur []int, res*[][]int) {
	if depth == len(nums) {
		tmpArr := []int{}
		tmpArr = append(tmpArr, cur...)
		*res = append(*res, tmpArr)
		return
	}

	SubSetsDFS(nums, depth+1, cur, res)
	cur = append(cur, nums[depth])
	SubSetsDFS(nums, depth+1, cur, res)
}


/*最小路劲和
还是回朔法给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
dp 计算到每格的最小和,回朔法也是ok的
 */
func GetMinRoadSumGrid (arr[][]int, rows,col int) (minSum int) {
	for i:=0;i<rows;i++ {
		for j:=0;j<col;j++ {
			if i==0 && j==0 {
				continue
			}
			if i==0 {
				arr[0][j] = arr[0][j-1]+arr[0][j]
				continue
			}
			if j==0 {
				arr[i][0] = arr[i-1][0]+arr[i][0]
				continue
			}
			min := 0
			if arr[i-1][j] < arr[i][j-1] {
				min = arr[i-1][j]
			} else {
				min = arr[i][j-1]
			}
			arr[i][j] =min+arr[i][j]
		}
	}
	minSum = arr[rows-1][col-1]
	return
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

//柱状图中最大的矩形面积
// 暴力解法
// 1.枚举宽度 先枚举左边界，在左边界到右边界之间遍历左边界确定最小高度 计算面积
// 2.枚举高度，遍历数组将每个元素作为矩形的高度，查看当以此为高度的时候面积最大是多少 找到左右第一个小于此高度的序列作为边界
/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。
 */

func largestRectangleArea(heights []int) int {
	maxArea := 0
	for i:=0;i<len(heights);i++ {
		var left,right = i,i

		for ;left -1 >=0; {
			if heights[left-1] >= heights[i] {
				left --
			} else {
				break
			}
		}
		for ;right +1 < len(heights); {
			if heights[right+1] >= heights[i] {
				right++
			} else {
				break
			}
		}
		area := (right-left+1) * heights[i]
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
//单调栈解法
type boarderItem struct {
	Height int
	index int
}
func largestRectangleAreaWithStack(heights []int) int {
	maxArea := 0
	stack := tree.Stack{}
	leftBorderArr := make([]int,len(heights))
	rightBoarderArr := make([]int,len(heights))

	for i:= 0; i< len(heights);i++ {
		tmpItem := boarderItem{}
		tmpItem.Height = heights[i]
		tmpItem.index = i
		if stack.IsEmpty() {
			leftBorderArr[i] = -1
		} else {
			for !stack.IsEmpty() {
				v := stack.Pop()
				item := v.(boarderItem)
				if item.Height < heights[i] {
					stack.Push(item)
					stack.Push(tmpItem)
					leftBorderArr[i] = item.index
					break
				}
			}
			if stack.IsEmpty() {
				stack.Push(tmpItem)
				leftBorderArr[i] = -1
			}
		}
	}
	stack = tree.Stack{}
	for j:= len(heights)-1; j>=0 ;j++ {
		tmpItem := boarderItem{}
		tmpItem.Height = heights[j]
		tmpItem.index = j
		if stack.IsEmpty() {
			rightBoarderArr[j] = len(heights)
		} else {
			for !stack.IsEmpty() {
				v := stack.Pop()
				item := v.(boarderItem)
				if item.Height < heights[j] {
					stack.Push(item)
					stack.Push(tmpItem)
					rightBoarderArr[j] = item.index
					break
				}
			}
			if stack.IsEmpty() {
				stack.Push(tmpItem)
				leftBorderArr[j] = len(heights)
			}
		}
	}
	for i:=0;i<len(heights);i++ {
		area := (rightBoarderArr[i]-leftBorderArr[i]-1)*heights[i]
		if area>maxArea{
			maxArea= area
		}
	}
	return  maxArea
}


/*
* 获取数组除自己以外的元素的乘积，不能使用除法
和接雨水相似，只不过接雨水是先找到左右两边的最大值在两个最大值中取最小，就是该节点能接的雨水
这个题是先计算 左边的乘积在计算右边的乘积
 */

func MultiExceptself (arr[]int) []int {
	multiLeft := make([]int, len(arr),len(arr))
	multiRight := make([]int, len(arr),len(arr))
	multiLeft[0] = 1
	multiRight[len(arr)-1] = 1
	for i:=1; i< len(arr);i++ {
		multiLeft[i] =  multiLeft[i-1] * arr[i]
	}

	for j:= len(arr)-2;j>0;j-- {
		multiRight[j] = multiRight[j+1] * arr[j]
	}

	res := make([]int, len(arr),len(arr))
	for i:=0; i< len(arr);i++ {
		res[i] = multiLeft[i] * multiRight[i]
	}
	return res
}

//买股票1 只能买一次

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

/*//获取最长递增子序列 位置和顺序和原数组相同长度 递归
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
}*/


//数组的所有子集 字符串的所有子串？ 这个得看个常规方法
func GetAllSubset (nums []int64) [][]int64 {
	res := [][]int64{}
	standardNum := (1 << len(nums))-1
	for i:=0; i<= standardNum;i++ {
		tmp := []int64{}
		for j := 0; j < len(nums);j++ {
			if (i&(1<<j)) > 0 {
				tmp = append(tmp,nums[j])
			}
		}
		res = append(res, tmp)
	}
	return res
}

//todo 200.岛屿数量
type point struct {
	x,y int
}
func GetIsLandNum(intMap [][]int) int {
	islandNum := 0
	maxLine := len(intMap)-1
	maxVolumn := len(intMap[0])-1
	for i:=0; i<=maxLine;i++ {
		for j:=0; j<= maxVolumn; j++ {
			if intMap[i][j] == 1 {
				islandNum ++
				queue := tree.Queue{}
				tmpPoint := point{
					i,j,
				}
				queue.In(tmpPoint)
				for !queue.IsEmpty() {
					tmp := queue.Out()
					curPoint := tmp.(point)
					if intMap[curPoint.x][curPoint.y] == 0 {
						continue
					}
					intMap[curPoint.x][curPoint.y] = 0
					if curPoint.y != 0 {
						tmpPoint := point{
							curPoint.x,curPoint.y-1,
						}
						queue.In(tmpPoint)
					}
					if curPoint.y !=maxVolumn {
						tmpPoint := point{
							curPoint.x,curPoint.y+1,
						}
						queue.In(tmpPoint)
					}
					if curPoint.x != maxLine {
						tmpPoint := point{
							curPoint.x+1,j,
						}
						queue.In(tmpPoint)
					}
				}
			}
		}
	}
	return islandNum
}

//括号生成

func GenerateParenthesis (target int) ([]string) {
	res := &([]string{})
	GenerateParenthesisDfs(0,target,'(',[]byte{},res)
	return *res
}
func GenerateParenthesisDfs (n, target int,v byte, tmpByte []byte, res *[]string) {
	tmpByte = append(tmpByte,v)
	if n == target*2-1 {
		if checkParenthesisValid(tmpByte) {
			*res = append(*res, string(tmpByte))
		}
		return
	}
	GenerateParenthesisDfs (n+1, target,'(', tmpByte,res)

	GenerateParenthesisDfs (n+1, target,')', tmpByte,res)
}

func checkParenthesisValid (tmpArr []byte) bool {
	stack := tree.Stack{}
	i:=0
	for ;i<len(tmpArr);i++ {
		if tmpArr[i] == '(' {
			stack.Push('(')
		} else {
			if !stack.IsEmpty() {
				v := stack.Pop()
				testByte := byte(v.(int32))
				if testByte != '(' {
					return false
				}
			} else {
				return false
			}
		}
	}
	if stack.IsEmpty() && i == len(tmpArr) {
		return true
	}
	return false
}
/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func MinWindow(s string, t string) string {
	charactorCursorMap := map[byte]int{}
	charactorPosition := map[byte][]int{}


	for i:=0; i< len(t);i++ {
		charactorCursorMap[t[i]] =0
		charactorPosition[t[i]] = []int{}
	}
	for i:=0; i< len(s);i++ {
		if _,ok:=charactorPosition[s[i]]; ok{
			charactorPosition[s[i]] = append(charactorPosition[s[i]], i)
		}
	}
	/*shortest := math.MaxInt32
	for _,v:=range charactorPosition {
		if len(v) < shortest {
			shortest = len(v)
		}
	}*/
	var minPosByte byte = t[0]
	needMvCursor := -1
	var mindistance = math.MaxInt32
	resPosStart,resPosEnd := 0,0
	for needMvCursor +1 < len(charactorPosition[minPosByte]) {
		charactorCursorMap[minPosByte] = needMvCursor +1
		minPos := math.MaxInt32
		maxPos := math.MinInt32
		for k,cur := range charactorCursorMap {
			if charactorPosition[k][cur] > maxPos {
				maxPos = charactorPosition[k][cur]
			}
			if charactorPosition[k][cur] < minPos {
				minPos = charactorPosition[k][cur]
				minPosByte = k
				needMvCursor = cur
			}
		}
		distance := maxPos - minPos
		if distance <mindistance {
			mindistance = distance
			resPosStart = minPos
			resPosEnd = maxPos
		}
	}

	return s[resPosStart:resPosEnd+1]
}

/*给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-image
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func RotateMatrix(matrix [][]int)  {
	n := len(matrix)
	for i:= 0; i< n/2; i++ {
		for j:= 0; j< n;j++ {
			matrix[i][j],matrix[n-i-1][j] = matrix[n-i-1][j],matrix[i][j]
		}
	}
	for i:=0;i<n;i++ {
		for j:=i;j<n;j++ {
			matrix[i][j],matrix[j][i] = matrix[j][i],matrix[i][j]
		}
	}
}

//下一个排列
/*实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须 原地 修改，只允许使用额外常数空间。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func GetNextSeq (arr []int) {
	var j = len(arr)-2
	for ;j>=0;j-- {
		if arr[j+1] < arr[j] {
			continue
		} else {
			break
		}
	}
	if j>=0 {
		//需要将较小数和较大数交换
		for i:=len(arr)-1;i>=(j+1);i-- {
			if arr[i] > arr[j] {
				arr[i],arr[j] = arr[j],arr[i]
				break
			}
		}
	}
	//重新修改为正序
	var low = j+1
	var high = len(arr)-1
	for low < high {
		arr[low],arr[high] = arr[high],arr[low]
		low ++
		high --
	}
}

//目标和
/*
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findTargetSumWays(nums []int, S int) int {
	res := 0
	findTargetSumWaysDfs(nums, 0, S, 0, &res)
	return res
}
//depth 控制选几次，内部的递归循环控制每次有几种选择
func findTargetSumWaysDfs(nums []int, sum int, target ,depth int,ways *int) {
	if depth == len(nums) {
		if sum == target {
			*ways++
		}
		return
	}
	//两次选择
	sum += nums[depth]
	findTargetSumWaysDfs(nums, sum, target,depth+1,ways)
	sum -= nums[depth]
	sum -= nums[depth]
	findTargetSumWaysDfs(nums, sum, target,depth+1,ways)
	sum += nums[depth]
}

// 不同路径
/*一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
如果有障碍

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
GetRobotWaysGridDp 动态规划
GetRobotWaysGridDFS 回溯法
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	return uniquePathsWithObstaclesDp(obstacleGrid)
}

func uniquePathsWithObstaclesDp (obstacleGrid [][]int) int {
	ways := make([][]int,len(obstacleGrid))
	for i:=0;i<len(ways);i++ {
		ways[i] = make([]int,len(obstacleGrid[0]))
	}
	ways[0][0] = 1
	for i:=0;i<len(ways);i++ {
		for j:=0;j<len(ways[0]);j++ {
			if i == 0 && j==0 {
				ways[i][j] = 1
			}else if i == 0 {
				if obstacleGrid[i][j] == 1 {
					ways[i][j] = 0
				} else {
					ways[i][j] = ways[i][j-1]
				}
			} else if j==0 {
				if obstacleGrid[i][j] == 1 {
					ways[i][j] = 0
				} else {
					ways[i][j] = ways[i-1][j]
				}
			} else {
				if obstacleGrid[i][j] == 1 {
					ways[i][j] = 0
				} else {
					ways[i][j] = ways[i-1][j] + ways[i][j-1]
				}
			}
		}
	}
	return ways[len(ways)-1][len(ways[0])-1]
}

func uniquePathsWithObstaclesDFS(obstacleGrid [][]int, curRow,curCol int, ways *int) {
	if curRow == len(obstacleGrid)-1 && curCol == len(obstacleGrid[0])-1 {
		//走完了 更新方法+1
		*ways = *ways+1
		return
	}
	if obstacleGrid[curRow][curCol] == 1 {
		//遇到路障
		return
	}
	if curRow+1 < len(obstacleGrid) {
		uniquePathsWithObstaclesDFS(obstacleGrid,curRow+1,curCol,ways)
	}
	if curCol+1 < len(obstacleGrid[0]) {
		uniquePathsWithObstaclesDFS(obstacleGrid,curRow,curCol+1,ways)
	}
}

//最长连续序列，不要求在数组中连续
// 可以先进行排序再搜索连续序列
// O(n)的算法
func longestConsecutive(nums []int) int {
	setMap := make(map[int]bool)
	max := 0
	for i:=0;i<len(nums);i++ {
		setMap[nums[i]] = true
	}

	for i:=0;i<len(nums);i++ {
		if _,ok := setMap[nums[i]-1];!ok {
			currentNum := nums[i]
			count := 1
			for {
				if _,ok:=setMap[currentNum+1];!ok {
					break
				}
				currentNum++
				count++
			}
			if count>max{
				max = count
			}
		}
	}
	return max
}










