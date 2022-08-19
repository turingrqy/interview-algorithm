package other

import (
	"fmt"
	"math"
	"reflect"
	"renqiyang/interview/tree"
	"strconv"
	"unsafe"
)

// 给定日志 user_id login_time logoutTime

type Log struct {
	LNTime int
	LOTime int
}

type PhaseRes struct {
	Start int
	End int
}

func GetPeekRange (data []Log) []PhaseRes{
	countArr := make([]int, 86400)
	for i:=0; i<len(data); i++ {
		countArr[data[i].LNTime]+=1
		if data[i].LOTime + 1 < len(countArr) {
			countArr[data[i].LOTime + 1]-=1
		}
	}

	for i:=1; i<len(countArr); i++ {
		countArr[i] = countArr[i]+countArr[i-1]
	}

	max := 0
	for i:=0; i<len(countArr); i++ {
		if countArr[i]>max {
			max = countArr[i]
		}
	}
	start := -1
	end := -1
	res := []PhaseRes{}
	for i:=0; i<len(countArr); i++ {
		if countArr[i] == max   {
			if start != -1 {
				start = i
			}
			end = i
		} else {
			res = append(res, PhaseRes{Start: start, End: end})
			start = -1
			end = -1
		}
	}
	if start != -1 {
		res = append(res, PhaseRes{Start: start, End: end})
	}
	return res
}

// 数组有重复的两数之和 找出所有组合
func GetTwoNumSumEqualTargetNotUniqArr (arr []int, target int)  {
	countMap := make(map[int]int)
	for _,v :=range arr {
		if _, ok := countMap[v]; !ok {
			countMap[v] = 0
		}
		countMap[v]++
	}

	for k, _ := range countMap {
		countMap[k] -=1
		coTarget := target-k
		if v, ok := countMap[coTarget]; ok && v > 0 {
			countMap[coTarget] -=1
			println(fmt.Sprintf("v1=%d, v2=%d", k, coTarget))
		}
	}

}

//s := arr[startIndex:endIndex]
//将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
//最长回文子串

//动态规划 如果aba是 cabac也是
//s[i+1,j-1]是回文的话 如果 s[i]==s[j] name s[i,j]也是
func GetLongestPalindrome (s string) {
	res := make([][]bool,len(s))
	for i:=0; i<len(s);i++ {
		res[i] = make([]bool,len(s))
	}
	ans := ""

	for l:=1; l<=len(s); l++ {
		for i:=0; i<=len(s)-l; i++ {
			j := i+l-1

			if l == 1  {
				res[i][j] = true
				ans = string(s[i])
				continue
			}
			if l==2 && s[i]==s[j] {
				res[i][j] = true
				ans = s[i:i+l]
			}

			if s[i]==s[j] && res[i+1][j-1] {
				res[i][j] = true
				ans = s[i:i+l]
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
	var sublen int
	for ;right < len(s);right++ {
		if index,ok := indexMap[s[right]];!ok {
			indexMap[s[right]] = right
			sublen = right-left+1
			if sublen > max {
				max = sublen
			}
		} else {
			for i:= left;i<=index;i++ {
				delete(indexMap,s[i])
			}
			left = index+1
			indexMap[s[right]] = right
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
			height = arr[right]
			right--
		} else {
			height = arr[left]
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



//解码a2[a2[ab]]c
func DecodeStr (text string) string {
	stack := tree.Stack{}

	for i:=0; i<len(text); i++ {
		if DecodeStrIsNumber(text[i]) {
			stack.Push(text[i])
		} else if text[i] != ']' {
			stack.Push(string(text[i]))
		} else {
			tmpStr := ""
			for !stack.IsEmpty() {
				sub := stack.Pop()
				subStr := sub.(string)
				if subStr == "[" {
					break
				}
				tmpStr = subStr+tmpStr
			}
			// 弹出数字
			dupNum := 0
			scale := 0
			for !stack.IsEmpty() {
				subStr := stack.Peek()
				subByte, ok := subStr.(byte)
				if !ok || !DecodeStrIsNumber(subByte) {
					break
				}
				stack.Pop()
				tmpInt := int(subByte-'0')
				dupNum = tmpInt * int(math.Pow(10,float64(scale))) + dupNum
			}
			finalStr := ""
			for i:=0; i < dupNum; i++ {
				finalStr += tmpStr
			}
			stack.Push(finalStr)
		}
	}
	res := ""
	for !stack.IsEmpty() {
		sub := stack.Pop()
		subStr := sub.(string)
		res = subStr + res
	}
	return res
}
func DecodeStrIsNumber (v byte) bool {
	if v - '0'<= 9 {
		return true
	}
	return false
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
// 删除有序数组中的重复项 不能使用额外空间 如果不是排序的可以先排序
func DeleteDupInSortedArr (arr *[]int) {
	if len(*arr) <= 1 {
		return
	}
	//慢指针
	i:=0
	//快指针
	j:=1
	for ;j<len(*arr);j++ {
		if (*arr)[i] != (*arr)[j] {
			i++
			(*arr)[i] = (*arr)[j]
		}
	}
	originSliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(arr))
	originSliceHeader.Len = i+1
	arr = (*[]int)(unsafe.Pointer(originSliceHeader))
}

// b^0=b
// a^b^a = b^(a^a)=b^0=b
//数组中只出现一次的数字 其他数字成对出现
//相同取0，相异取1。（二进制）
/*不重复的数字有两个，剑指offer上面也有“找出数组中只出现了一次的两个数”，解决办法是先遍历一遍数组，异或得到一个数N，这个数就是只出现一次的那两个数异或的结果，然后找到N最低为1的位（假设是m位），再次遍历数组，按m位为1和为0将数组分为两个数组，此时只出现一次的两个数就被分到了不同的组，然后对每个组按照（1）的方法找出来就可以了，代码如下：
————————————————
版权声明：本文为CSDN博主「ddxu」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/d12345678a/article/details/54233795*/
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

// 子集
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

//递增子序列
func GetAllSubUpSet (nums []int) {
	cur := []int{}
	GetAllSubUpSetDfs(nums, 0, cur)
}
// 有重复的元素，可以选择，但是组合不能重复
func GetAllSubUpSetDfs (nums []int, depth int, cur []int) {
	if depth == len(nums) {
		if len(cur) >= 2 {
			println(fmt.Sprintf("GetAllSubUpSetDfs=%v", cur))
		}
		return
	}

	if len(cur) == 0 || cur[len(cur)-1] != nums[depth] {
		GetAllSubUpSetDfs(nums, depth+1, cur)
	}

	if len(cur) == 0 || nums[depth] >= cur[len(cur)-1] {
		cur = append(cur, nums[depth])
		GetAllSubUpSetDfs(nums, depth+1, cur)
	}
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

func GetMinRoadSumGridByDfs (arr[][]int) int {
	row := 0
	col := 0
	min := math.MaxInt32
	GetMinRoadSumGridDfs(arr, row, col,0, &min)
	return min
}

func GetMinRoadSumGridDfs (arr [][]int, row, col, sum int, min *int) {
	if row == len(arr)-1 && col == len(arr[0])-1 {
		sum += arr[row][col]
		if sum < *min {
			*min = sum
		}
		return
	}
	sum += arr[row][col]
	if row + 1 <= len(arr)-1 {
		GetMinRoadSumGridDfs(arr, row+1, col, sum, min)
	}
	if col + 1 <= len(arr[0])-1 {
		GetMinRoadSumGridDfs(arr, row, col+1, sum, min)
	}

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
	maxLeft[0] = arr[0]
	maxRight[len(arr)-1] = arr[len(arr)-1]
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

		for left < len(heights) {
			if heights[left] >= heights[i] {
				left--
			} else{
				break
			}
		}

		for right < len(heights) {
			if heights[right] >= heights[i] {
				right++
			} else {
				break
			}
		}
		area := heights[i] * (right-left-1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
//单调栈解法
//有一群牛站成一排，每头牛都是面朝右的，每头牛可以看到他右边身高比他小的牛。给出每头牛的身高，要求每头牛能看到的牛的总数。右边第一个比他大的数
//给出一个序列，求出一个子序列，使得这个序列中的最小值乘以这个序列的和的值最大 左边第一个比他小的数， 右边第一个比他小数

//On 每个元素只会入栈出站一次 2n on
func largestRectangleAreaWithStack(heights []int) int {
	leftBoarder := make([]int, len(heights))
	rightBoarder := make([]int, len(heights))

	stack := tree.Stack{}
	for i := 0; i < len(heights); i++ {
		for !stack.IsEmpty() {
			tmp := stack.Peek()
			lastIndex :=  tmp.(int)
			if heights[lastIndex] >= heights[i] {
				stack.Pop()
			} else {
				break
			}
		}
		if stack.IsEmpty() {
			leftBoarder = append(leftBoarder, -1)
		} else {
			tmp := stack.Peek()
			lastIndex :=  tmp.(int)
			leftBoarder = append(leftBoarder, lastIndex)
			stack.Push(i)
		}
	}
	stack = tree.Stack{}
	for j := len(heights)-1; j >=0; j-- {
		for !stack.IsEmpty() {
			tmp := stack.Peek()
			lastIndex :=  tmp.(int)
			if heights[lastIndex] >= heights[j] {
				stack.Pop()
			} else {
				break
			}
		}
		if stack.IsEmpty() {
			rightBoarder = append(rightBoarder, len(heights))
		} else {
			tmp := stack.Peek()
			lastIndex :=  tmp.(int)
			rightBoarder = append(rightBoarder, lastIndex)
			stack.Push(j)
		}
	}
	maxArea := 0
	for k:=0; k<len(heights); k++ {
		area := (rightBoarder[k]-leftBoarder[k]) -1 * heights[k]
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
//无序数组输出每个元素右边第一个比他大的元素
func GetFirstMaxFromRight (arr []int) []int {
	rightBoarder := make([]int, len(arr))
	stack := tree.Stack{}

	for i := len(arr)-1; i>=0; i-- {
		for !stack.IsEmpty() {
			tmp := stack.Peek()
			tmpInt :=  tmp.(int)
			if arr[tmpInt] <= arr[i] {
				stack.Pop()
			} else {
				break
			}
		}
		if stack.IsEmpty() {
			rightBoarder[i] = -1
		} else {
			tmp := stack.Peek()
			tmpInt := tmp.(int)
			rightBoarder[i] = tmpInt
			stack.Push(i)
		}
	}
	return rightBoarder
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
		multiLeft[i] =  multiLeft[i-1] * arr[i-1]
	}

	for j:= len(arr)-2;j>=0;j-- {
		multiRight[j] = multiRight[j+1] * arr[j+1]
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

func GenerateParenthesis (target int) []string {
	path := make([]byte, 2*target)
	res := make([]string,0,0)
	GenerateParenthesisDfs(path, 0, 0, target, &res)
	return res
}
func GenerateParenthesisDfs (path []byte, open, close, target int, res *[]string) {
	if open == target && close == target {
		*res = append(*res, string(path))
		return
	}

	if open < target {
		path =  append(path, '(')
		GenerateParenthesisDfs(path, open+1, close, target, res)
		path = path[:len(path)-1]
	}

	if close < open {
		path =  append(path, ')')
		GenerateParenthesisDfs(path, open, close+1, target, res)
		path = path[:len(path)-1]
	}
}

func checkParenthesisValid (tmpArr []byte) bool {
	balance := 0
	for i:=0; i < len(tmpArr); i++ {
		if tmpArr[i] == '(' {
			balance++
		} else if tmpArr[i] == ')' {
			balance--
		}

		if balance < 0 {
			return false
		}
	}
	return balance == 0
}
/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
保证最后一个无效的右括号在栈底
 */

func GetLongestValidParenthesis (tmpStr string) int {
	stack := tree.Stack{}
	maxLen := 0
	stack.Push(-1)
	for k, v := range tmpStr {
		if v == '(' {
			stack.Push(k)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(k)
			} else {
				tmpVal := stack.Peek()
				tmpInt := tmpVal.(int)
				maxLen = int(math.Max(float64(maxLen), float64(k-tmpInt)))
			}
		}
	}
	return maxLen
}


/*
最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
滑动窗口解法
*/
func MinWindow(s string, t string) string {
	ansLen := math.MaxInt32
	ansL, ansR := -1, -1
	oriCnt := make(map[byte]int, 0)

	for k:=0; k < len(t); k++ {
		oriCnt[t[k]]++
	}

	tmpCnt := make(map[byte]int, 0)
	check := func () bool {
		for k, v := range oriCnt {
			tmp, ok := tmpCnt[k]
			if !ok {
				return false
			}
			if tmp < v {
				return false
			}
		}
		return true
	}

	for l,r := 0,0; r < len(s); r++ {
		if _,ok := oriCnt[s[r]]; ok {
			tmpCnt[s[r]]++
		}

		for check() && l <= r {
			tmpLen :=  r-l+1
			if tmpLen < ansLen {
				ansLen = tmpLen
				ansL, ansR = l, r
			}

			if _,ok := tmpCnt[s[l]]; ok {
				tmpCnt[s[l]]--
			}
			l++
		}
	}

	return s[ansL:ansR+1]
}

/*给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
先上下交换，再对角线交换
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
		for j:=0;j<i;j++ {
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

	findTargetSumWaysDfs(nums, sum+nums[depth], target,depth+1,ways)

	findTargetSumWaysDfs(nums, sum-nums[depth], target,depth+1,ways)

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

//最长连续序列，不要求在数组中连续 不要求相对位置
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


//矩阵传染 target =2 start (i,j) source=3
func SpreadGrid (arr[][] int, target int, source int, startI,startJ int ) {
	dp := [][]int{}
	startPoint := []int{startI,startJ}
	dp = append(dp, startPoint)
	for i:=0; i< len(dp);i++ {
		if arr[dp[i][0]][dp[i][1]] == source {
			arr[dp[i][0]][dp[i][1]] = target
			if dp[i][1]-1 >= 0 {
				startPoint = []int{dp[i][0], dp[i][1]-1}
				dp = append(dp,startPoint)
			}
			if dp[i][1]+1 < len(arr[0]) {
				startPoint = []int{dp[i][0], dp[i][1]+1}
				dp = append(dp,startPoint)
			}
			if dp[i][0]-1 >= 0 {
				startPoint = []int{dp[i][0]-1, dp[i][1]}
				dp = append(dp,startPoint)
			}
			if dp[i][0]+1 < len(arr) {
				startPoint = []int{dp[i][0]+1, dp[i][1]}
				dp = append(dp,startPoint)
			}
		}
	}
}

//棋盘覆盖问题
func CoverGrid (size int, bc,br int) {
	Grid := make ([][]int, size)
	for i:= 0; i< size;i++ {
		Grid[i] = make([]int, size)
	}
	num := 0
	CoverGridPartition(Grid, 0,0,bc,br,size,&num)
	for i:= 0; i< size;i++ {
		println(fmt.Sprintf("res=%+v", Grid[i]))
	}
}

func CoverGridPartition (Grid [][]int,sc, sr, bc, br,size int, num *int) {
	if size == 1 {
		return
	}
	newSize := size/2
	*num++
	t := *num

	if bc < sc + newSize && br < sr + newSize {
		//block 点在当前的分区内
		CoverGridPartition(Grid, sc, sr, bc, br, newSize, num)
	} else {
		Grid[sr + newSize-1][sc + newSize-1] = t
		CoverGridPartition(Grid, sc, sr, sc + newSize-1, sr + newSize-1, newSize, num)
	}

	if bc >= sc + newSize && br < sr + newSize {
		CoverGridPartition(Grid, sc + newSize, sr, bc, br, newSize, num)
	} else {
		//左下角覆盖
		Grid[sr + newSize-1][sc + newSize] = t
		CoverGridPartition(Grid, sc + newSize, sr, sc + newSize, sr + newSize-1, newSize, num)
	}

	if br >= sr + newSize && bc < sc + newSize {
		CoverGridPartition(Grid, sc, sr + newSize, bc, br, newSize, num)
	} else {
		//右上角覆盖
		Grid[sr + newSize][sc + newSize-1] = t
		CoverGridPartition(Grid, sc, sr + newSize, sc + newSize-1, sr + newSize, newSize, num)
	}

	if br >=sr + newSize && bc >= sc +newSize {
		CoverGridPartition(Grid, sc + newSize, sr + newSize, bc, br, newSize, num)
	} else {
		Grid[sr + newSize][sc +newSize] = t
		CoverGridPartition(Grid, sc + newSize, sr + newSize, sc +newSize, sr + newSize, newSize, num)
	}
}
// 调整文本左右对齐
func JustifyWord (words []string, cap int)  {
	start := 0
	end := 0
	//curStringArr := []string{}
	curLen := 0
	curWordLen := 0
	curWordNum := 0
	for ;end < len(words); {

		if curLen == 0 {
			if len(words[end]) > cap {
				return
			}
			//curStringArr = append(curStringArr, words[end])
			curLen += len(words[end])
			curWordLen += len(words[end])
			curWordNum++
			end ++
			continue
		}
		if len(words[end]) + 1 + curLen <= cap {
			curLen += len(words[end]) + 1
			curWordLen += len(words[end])
			curWordNum++
			end ++
		} else {
			//放不下了
			//计算 需要多少空格填充
			fillBlankNum := cap-curWordLen
			divide := fillBlankNum/(curWordNum-1)
			divideleft := fillBlankNum%(curWordNum-1)
			curString := ""
			for i:= start;i<end;i++ {
				curString += words[i]
				if i != end-1 {
					for i:=0;i<divide;i++ {
						curString += " "
					}
					if divideleft > 0 {
						curString += " "
						divideleft --
					}
				}
			}
			println(curString)
			curLen = 0
			curWordLen = 0
			curWordNum = 0
			start = end
		}
	}
	if start < len(words) {
		str := ""
		for i:=start;i<len(words);i++ {
			if str == "" {
				str += words[i]
			} else {
				str += " "
				str +=  words[i]
			}
		}
		println(str)
	}
}

/*
给你一个整数数组 nums 。「数组值」定义为所有满足 0 <= i < nums.length-1 的 |nums[i]-nums[i+1]| 的和。

你可以选择给定数组的任意子数组，并将该子数组翻转。但你只能执行这个操作 一次 。

请你找到可行的最大 数组值 。
变化的值 abs(a[l-1]-a[r]) + abs(a[l]-a[r+1]) - abs(a[l-1]-a[l]) + abs(a[r]-a[r+1])

(a[l−1]+a[l]−abs(a[l]−a[l−1]))−(a[r]+a[r+1]+abs(a[r]−a[r+1])),
(a[l−1]−a[l]−abs(a[l]−a[l−1]))−(a[r]−a[r+1]+abs(a[r]−a[r+1])),
(−a[l−1]+a[l]−abs(a[l]−a[l−1]))−(−a[r]+a[r+1]+abs(a[r]−a[r+1])),
(−a[l−1]−a[l]−abs(a[l]−a[l−1]))−(−a[r]−a[r+1]+abs(a[r]−a[r+1]))
分成前后两部分看

作者：_hututu
链接：https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value/solution/onzuo-fa-jie-jue-ci-wen-ti-by-hu-tu-tu-7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

边界值
l 为0  abs(a[0]-arr[r+1])-abs(arr[r]-arr[r+1])
r = 末位置 abs(a[l-1]-a[l]) - abs(arr[l-1]-arr[n-1])
 */

func GetAbsDiffAfterReverse (arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[1]
	}

	origRes :=0
	for i:=0; i < len(arr)-1; i++ {
		origRes += arr[i]-arr[i-1]
	}
	var maxDiff = math.MinInt32
	for i:=0; i < len(arr); i++ {
		if i != 0 {
			//l为0
			maxDiff = int(math.Max(float64(maxDiff),math.Abs(float64(arr[0]-arr[i+1])) - math.Abs(float64(arr[i]-arr[i+1]))))

		}
		if i != len(arr)-1 {
			// r 为n-1
			maxDiff = int(math.Max(float64(maxDiff),math.Abs(float64(arr[0]-arr[i+1])) - math.Abs(float64(arr[i]-arr[i+1]))))
		}
	}
	var myX = []int{1,1,-1,-1}
	var myY = []int{1,-1,1,-1}

	for i:=0; i<4;i++ {
		left := make([]int,0,0)
		right := make([]int,0,0)
		for j := 0; j < len(arr)-1; j++ {
			cur := int(math.Abs(float64(arr[j]-arr[j+1])))
			v1 := myX[i]*arr[j]
			v2 := myY[i]*arr[j]
			left = append(left, v1 + v2 - cur)
			right = append(right, v1 + v2 +cur)
		}
		maxV1 := GetAbsDiffAfterReverseGetMaxInArr(left)
		minV2 := GetAbsDiffAfterReverseGetMinInArr(right)
		maxDiff = int(math.Max(float64(maxDiff), float64(maxV1-minV2)))
	}
	return origRes+maxDiff
}

func GetAbsDiffAfterReverseGetMaxInArr (arr []int) int {
	var max = math.MinInt32
	for i:=0; i<len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func GetAbsDiffAfterReverseGetMinInArr (arr []int) int {
	var min = math.MaxInt32
	for i:=0; i<len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}



// 整数反转
/*给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func reverseInt (num int) int {
	if num == 0 {
		return 0
	}
	newNum := 0

	for num%10 > 0 {
		newNum += newNum*10 + num%10
		num /= 10
	}
	return newNum
}


/*
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
数组内的pairs 通过左端点排序，再遍历排序后的pairs 一个个处理看是否有区间重合 重合则拿最大值最小值重新划定区间
*/

//字符串相加
func StrSum (num1 string, num2 string) string {
	tmpByteArr := make([]byte,0,0)

	i := len(num1) -1
	j := len(num2) -1

	carry := uint8(0)
	for ;i>=0 && j >=0; {
		cur1 := num1[i] - '0'
		cur2 := num2[j] - '0'
		curNum := cur1+cur2+carry
		curNum = curNum%10
		carry = curNum/10
		curNumByte := curNum + '0'
		tmpByteArr = append(tmpByteArr, curNumByte)
	}


	if i >= 0 {
		for i >=0 {
			tmpByteArr = append(tmpByteArr, num1[i])
			i--
		}
	}

	if j >=0 {
		for j >=0 {
			tmpByteArr = append(tmpByteArr, num2[j])
			j--
		}
	}
	sumByteArr := make([]byte, 0, len(tmpByteArr))
	for i:=len(tmpByteArr)-1; i>=0; i-- {
		sumByteArr = append(sumByteArr, tmpByteArr[i])
	}
	return string(sumByteArr)
}

// 字符串乘法
/*
  则可以通过模拟「竖式乘法」的方法计算乘积。从右往左遍历乘数，将乘数的每一位与被乘数相乘得到对应的结果
 */

func StrMulti (text1, text2 string) string {
	len1 := len(text1)
	len2 := len(text2)
	ans := "0"
	for i:= len1-1; i >= 0; i-- {
		cur := ""
		for j := len2-1; j > i; j-- {
			cur += "0"
		}
		carry := 0
		for j := len1 - 1; j >= 0; j-- {
			x := text1[j] - '0'
			y := text2[i] - '0'
			tmp := int(x) * int(y) + carry
			cur = strconv.Itoa(tmp%10) + cur
			carry = tmp/10
		}
		for ;carry != 0; carry/=10 {
			cur = strconv.Itoa(carry%10) + cur
		}
		ans = StrSum(ans, cur)
	}
	return ans
}

/*给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。

对于每个单元格，你可以往上，下，左，右四个方向移动。 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
深度优先遍历 遍历过的不再遍历
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/fpTFWP
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func GetLongestUpSeqInGrid (grid [][]int) int {
	maxRow := len(grid)
	maxCol := len(grid[0])
	memo := make([][]int, len(grid))
	for i:=0; i<len(grid); i++ {
		memo[i] = make([]int, maxCol)
	}
	ans := 0
	for i:=0; i<len(grid);i++ {
		for j:=0; i<len(grid[0]); j++ {
			ans = int(math.Max(float64(ans), float64(GetLongestUpSeqInGridDfs(grid, i, j, maxRow, maxCol, memo))))
		}
	}
	return ans
}

func GetLongestUpSeqInGridDfs (grid [][]int, row, col, maxRow, maxCol int, memo [][]int) int {
	if memo[row][col] != 0 {
		return memo[row][col]
	}

	//向上走
	if row - 1 >= 0 && grid[row - 1][col] > grid[row][col] {
		memo[row][col] = int(math.Max(float64(memo[row][col]), float64(GetLongestUpSeqInGridDfs(grid, row-1, col, maxRow, maxCol, memo))+1))
	}

	//向下走
	if row + 1 < maxRow && grid[row + 1][col] > grid[row][col] {
		memo[row][col] = int(math.Max(float64(memo[row][col]), float64(GetLongestUpSeqInGridDfs(grid, row+1, col, maxRow, maxCol, memo))+1))
	}

	// 向左走
	if col - 1 >= 0 && grid[row][col-1] > grid[row][col] {
		memo[row][col] = int(math.Max(float64(memo[row][col]), float64(GetLongestUpSeqInGridDfs(grid, row, col-1, maxRow, maxCol, memo))+1))
	}

	if col + 1 >= 0 && grid[row][col+1] > grid[row][col] {
		memo[row][col] = int(math.Max(float64(memo[row][col]), float64(GetLongestUpSeqInGridDfs(grid, row, col+1, maxRow, maxCol, memo))+1))
	}
	return memo[row][col]
}

func GetIntByStr (str string) int64 {
	//边界条件
	sum := int64(0)
	lastAns := int64(0)
	for i:=0; i<len(str); i++ {

		lastAns = sum
		sum = (sum * 10) + int64(str[i]-'0')
		if sum < 0 {
			return -1
		}
	}
	return lastAns
}





























