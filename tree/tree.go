package tree

import (
	"fmt"
	"math"
)


type Stack struct {
	arr []interface{}
}

func (s *Stack) Push(val interface{}) {
	s.arr = append(s.arr, val)
}

func (s *Stack) Pop() (val interface{}) {
	val = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return
}

func (s *Stack) Peek() interface{} {
	return s.arr[len(s.arr)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

type Queue struct {
	arr []interface{}
}

func (q *Queue) in (val interface{}) {
	q.arr = append(q.arr,val)
}
func (q *Queue) In (val interface{}) {
	q.in(val)
}

func (q *Queue) IsEmpty () bool {
	return len(q.arr) == 0
}

func (q *Queue) out () (val interface{})  {
	val = q.arr[0]
	if len(q.arr) > 1 {
		q.arr = q.arr[1:]
	} else {
		q.arr = []interface{}{}
	}
	return
}
func (q *Queue) Out() (val interface{}) {
	return q.out()
}


type TreeNode struct {
	Value int64
	Left *TreeNode
	Right *TreeNode
	Parent *TreeNode
}

func CreateTreeOrderedArr (arr []int64) *TreeNode {

	if len(arr) == 0 {
		return nil
	}

	if len(arr) == 1 {
		return &TreeNode{
			arr[0],
			nil,
			nil,
			nil,
		}
	}
	low :=0
	high := len(arr) -1
	mid := (low+high)/2
	root := &TreeNode{
		arr[mid],
		nil,
		nil,
		nil,
	}

	root.Right = CreateTreeOrderedArr(arr[(mid +1):])
	if root.Right != nil {
		root.Right.Parent = root
	}
	if mid != 0 {
		root.Left = CreateTreeOrderedArr(arr[:mid])
		if root.Left != nil {
			root.Left.Parent = root
		}
	}


	return root
}

func PreorderTree (root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Print(root.Value," ")
	PreorderTree(root.Left)
	PreorderTree(root.Right)
}

func PreOrderStack (root *TreeNode) {
	stack := &Stack{}
	tmp := root
	for !stack.IsEmpty() || tmp != nil {

		for tmp != nil  {
			fmt.Print(tmp.Value," ")
			stack.Push(tmp)
			tmp = tmp.Left
		}

		if !stack.IsEmpty() {
			tmpinter := stack.Pop()

			tmp,_ = tmpinter.(*TreeNode)

			tmp = tmp.Right

		}
	}
}


func MidOrderTree(root *TreeNode) {
	if root == nil {
		return
	}

	MidOrderTree(root.Left)
	fmt.Print(root.Value," ")
	MidOrderTree(root.Right)
}

func MidOrderTreeStack(root *TreeNode) {
	stack := &Stack{}
	tmp := root
	for !stack.IsEmpty() || tmp != nil  {
		for tmp != nil {
			stack.Push(tmp)
			tmp = tmp.Left
		}

		if !stack.IsEmpty() {
			val := stack.Pop()
			tmp = val.(*TreeNode)
			print(tmp.Value," ")
			tmp = tmp.Right
		}
	}
}

func LastOrderTree (root *TreeNode) {
	if root == nil {
		return
	}
	LastOrderTree(root.Left)
	LastOrderTree(root.Right)
	fmt.Print(root.Value, " ")
}

func LastOrderTreeStack (root *TreeNode) {
	stack := &Stack{}
	tmp := root
	var lastVisitNode *TreeNode = nil

	for tmp != nil {
		stack.Push(tmp)
		tmp = tmp.Left
	}

	for !stack.IsEmpty() {
		val := stack.Pop()
		tmp = val.(*TreeNode)

		if tmp.Right == nil || lastVisitNode == tmp.Right {
			fmt.Print(tmp.Value, " ")
			lastVisitNode = tmp
		} else {
			stack.Push(tmp)
			tmp = tmp.Right
			for tmp != nil {
				stack.Push(tmp)
				tmp = tmp.Left
			}
		}
	}
}

func WideOrderTree (root *TreeNode) {
	q := &Queue{}
	q.in(root)
	for !q.IsEmpty() {
		val := q.out()

		tmp := val.(*TreeNode)
		fmt.Println(tmp.Value)
		if tmp.Left != nil {
			q.in(tmp.Left)
		}

		if tmp.Right != nil {
			q.in(tmp.Right)
		}
	}
}
//二叉树右视图，层序遍历
func GetTreeRightView(root *TreeNode) []int64 {
	tmpArr := []*TreeNode{}
	tmpArr = append(tmpArr, root)
	resArr := []int64{}
	for len(tmpArr)>0 {
		resArr = append(resArr,tmpArr[len(tmpArr)-1].Value)
		curTmpArr := []*TreeNode{}
		for i:=0;i<len(tmpArr);i++ {
			if tmpArr[i].Left != nil {
				curTmpArr = append(curTmpArr,tmpArr[i].Left)
			}
			if tmpArr[i].Right != nil {
				curTmpArr = append(curTmpArr,tmpArr[i].Right)
			}
		}
		tmpArr = curTmpArr
	}
	return resArr
}



func GetHeight (root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := GetHeight(root.Left)
	rightHeight := GetHeight(root.Right)
	return Max(leftHeight,rightHeight)+1
}

func IsBalanceTree (root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isLeftBalance , leftHeight := IsBalanceTree(root.Left)

	if !isLeftBalance {
		return false, 0
	}

	isRightBalance , rightHeight := IsBalanceTree(root.Right)
	if !isRightBalance {
		return false, 0
	}

	if abs((leftHeight-rightHeight)) <=1 {
		return true,Max(leftHeight, rightHeight) + 1
	}

	return false,0
}

func GetLeveledNode (root *TreeNode) [][]int64 {
	levelArr := []*TreeNode{}
	result := [][]int64{}

	levelArr = append(levelArr, root)
	for len(levelArr) > 0 {
		tmp_arr := []int64{}
		for i:= 0; i< len(levelArr); i++ {
			tmp_arr = append(tmp_arr, levelArr[i].Value)
		}
		result = append(result, tmp_arr)
		tmp_level_arr := levelArr
		levelArr = []*TreeNode{}
		for _,node :=range tmp_level_arr {
			if node.Left != nil {
				levelArr = append(levelArr, node.Left)
			}

			if node.Right != nil {
				levelArr = append(levelArr, node.Right)
			}
		}
	}

	return result
}

func GetLevelNodeNum (root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	if k == 1 {
		return 1
	}

	leftNodeNum := GetLevelNodeNum(root.Left, k-1)
	rightNodeNum := GetLevelNodeNum(root.Right, k-1)
	return leftNodeNum + rightNodeNum
}

func IsSearchTree (root *TreeNode, min int64, max int64) bool {
	if root == nil {
		return true
	}

	if root.Value < min || root.Value > max {
		return false
	}

	isleft := IsSearchTree(root.Left, min, root.Value)
	isRight := IsSearchTree(root.Right, root.Value, max)

	return (isleft && isRight)
}

//中序遍历方法
func IsSearchTreeMidOrder (root *TreeNode, prev int64) bool {

	if root == nil {
		return true
	}

	isLeft := IsSearchTreeMidOrder(root.Left, prev)
	if isLeft {
		if root.Value < prev {
			return false
		} else {
			prev = root.Value
			return IsSearchTreeMidOrder(root.Right, prev)
		}
	}
	return false
}
//中序非递归
func IsSearchTreeMidOrderStack (root *TreeNode) bool {
	stack := &Stack{}
	tmp :=root
	prev := int64(math.MinInt64)
	for tmp!= nil && !stack.IsEmpty() {
		for tmp != nil  {
			stack.Push(tmp)
			tmp = tmp.Left
		}

		if !stack.IsEmpty() {
			val := stack.Pop()
			tmp = val.(*TreeNode)
			if tmp.Value < prev {
				return false
			} else {
				prev = tmp.Value
			}
			tmp = tmp.Right
		}
	}
	return true
}
// 没有的话中序遍历
func FindTreeMidOrderNextNode (node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		tmp := node.Right
		for tmp.Left!= nil {
			tmp = tmp.Left
		}
		return tmp
	} else {
		parent := node.Parent
		tmp := node
		for parent != nil && tmp != parent.Left {
			tmp = parent
			parent = parent.Parent
		}
		return parent
	}
}
func FindTreeMidOrderNextNodeLoop (root, node, prev *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return nil
	}

	leftNode := FindTreeMidOrderNextNodeLoop(root.Left, node, prev)
	if leftNode != nil {
		return leftNode
	}
	if prev == node {
		return root
	}
	prev = root
	rightNode := FindTreeMidOrderNextNodeLoop(root.Right, node, prev)
	return rightNode
}
//搜索二叉树
func FindCommonAccesstorIBST(root *TreeNode, p *TreeNode, q*TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	tmp := root

	var minNode, maxNode *TreeNode

	if p.Value > q.Value {
		maxNode = p
		minNode = q
	} else {
		maxNode = q
		minNode = p
	}

	for tmp != nil {
		if tmp.Value >= minNode.Value && tmp.Value <= maxNode.Value {
			return tmp
		} else if tmp.Value > maxNode.Value {
			tmp = tmp.Left
		} else if tmp.Value < minNode.Value  {
			tmp = tmp.Right
		}
	}
	return nil
}

func getParentListLen (node *TreeNode) int {
	len := 0

	for node != nil {
		len ++
		node = node.Parent
	}
	return len
}
//有父指针转化为链表的公共节点
func FindCommonAccesstorIBSTParent(root *TreeNode, p *TreeNode, q*TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	lenp := getParentListLen(p)
	lenq := getParentListLen(q)

	var diffNum int
	tmpp,tmpq := p,q

	if lenp >= lenq {
		diffNum = lenp - lenq
		for i:=0; i< diffNum; i++ {
			tmpp = tmpp.Parent
		}
	} else {
		diffNum = lenq - lenp
		for i:=0; i< diffNum; i++ {
			tmpq = tmpq.Parent
		}
	}
	for tmpq != tmpp  {
		tmpq = tmpq.Parent
		tmpp = tmpp.Parent
	}

	return tmpp
}

//没有父指针 非递归思路 先找出路径
func FindCommonAccesstor (root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := FindCommonAccesstor(root.Left, p, q)
	right := FindCommonAccesstor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}

	return nil
}


func GetRootToNode (root *TreeNode, node *TreeNode, num *int) bool {
	if root == nil {
		return false
	}

	if root == node {
		return true
	}

	(*num)++
	isInLeft := GetRootToNode(root.Left, node, num)
	isInRight := GetRootToNode(root.Right, node, num)
	if !isInLeft && !isInRight {
		(*num)--
	}

	return (isInLeft || isInRight)
}

func GetRootToNodePath (root *TreeNode, node *TreeNode, path *[]*TreeNode) bool {
	if root == nil {
		return false
	}

	if root == node {
		*path = append(*path, root)
		for _,v := range *path {
			println(fmt.Sprintf("tree path=%d", v.Value))
		}
		return true
	}

	*path = append(*path, root)
	isInLeft := GetRootToNodePath(root.Left, node, path)
	isInRight := GetRootToNodePath(root.Right, node, path)
	if !isInRight && !isInLeft {
		*path = (*path)[:len(*path)-1]
	}
	return isInLeft || isInRight
}

func GetNodeInKlevel (root *TreeNode, k int) []*TreeNode {
	levelArr := []*TreeNode{}
	levelArr = append(levelArr, root)
	currentLevel := 0
	if currentLevel == k-1 {
		return levelArr
	}
	for len(levelArr) != 0 {
		if currentLevel == k-1 {
			return levelArr
		}
		currentLevel ++
		tmpArr := []*TreeNode{}
		for i:=0; i< len (levelArr); i++ {
			if levelArr[i].Left != nil {
				tmpArr = append(tmpArr, levelArr[i].Left)
			}

			if levelArr[i].Right != nil {
				tmpArr = append(tmpArr, levelArr[i].Right)
			}
		}
		levelArr = tmpArr
	}

	return []*TreeNode{}
}

func GetNodeInKLevelLoop (root *TreeNode, k int, res *[]*TreeNode) {
	if root == nil {
		return
	}

	if k == 1 {
		*res = append(*res, root)
		return
	}

	GetNodeInKLevelLoop(root.Left, k-1, res)
	GetNodeInKLevelLoop(root.Right, k-1, res)
	return
}

func GetRootToNodeRoad (root *TreeNode, node *TreeNode, arr *[]*TreeNode) bool {
	if root == nil {
		return false
	}

	if root == node {
		return true
	}

	*arr = append(*arr, root)

	isInLeft := GetRootToNodeRoad(root.Left, node, arr)
	isInRight := GetRootToNodeRoad(root.Right, node, arr)

	if !isInLeft && !isInRight {
		*arr = (*arr)[:len(*arr)-1]
	}

	return isInLeft||isInRight
}

func IsChildTree (root1 *TreeNode, root2 *TreeNode) bool {
	if isMatch(root1, root2) {
		return true
	}

	return IsChildTree(root1.Left,root2) || IsChildTree(root1.Right,root2)
}

func isMatch (root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 != nil {
		return false
	}

	if root1 == nil && root2 == nil {
		return true
	}

	if root1 != nil && root2 == nil {
		return false
	}

	if root1.Value == root2.Value {
		return true
	}

	return (isMatch(root1.Left, root2.Left)&&isMatch(root1.Right, root2.Right))
}

func GetSumRoad (root *TreeNode, arr []*TreeNode, expectedSum int64)  {
	if root == nil {
		return
	}

	if expectedSum - root.Value < 0 {
		return
	}
	arr = append(arr, root)

	if root.Left == nil && root.Right == nil && expectedSum-root.Value == 0 {
		for _, v := range arr {
			println(v.Value)
		}
		return
	}

	GetSumRoad(root.Left, arr, expectedSum-root.Value)
	GetSumRoad(root.Right, arr, expectedSum-root.Value)
}

func ConvertTree2Link (root *TreeNode) *TreeNode {
	var pLast *TreeNode
	convertTreeNode(root, &pLast)
	for pLast.Left != nil {
		pLast = pLast.Left
	}

	return pLast
}

func convertTreeNode (root *TreeNode, pLast **TreeNode) {
	if root == nil {
		return
	}

	convertTreeNode(root.Left, pLast)
	root.Left = *pLast
	if (*pLast) != nil {
		(*pLast).Right = root
	}
	*pLast = root
	convertTreeNode(root.Right, pLast)
}


func SnakePrintTree (root *TreeNode) {
	if root ==nil {
		return
	}

	levelArr := []*TreeNode{root}
	level := 1
	for len (levelArr) > 0 {
		if level%2 == 1 {
			PrintArrOrder(levelArr,true)
		} else {
			PrintArrOrder(levelArr,false)
		}

		tmp := []*TreeNode{}
		for _,item :=range levelArr {
			if item.Left != nil {
				tmp = append(tmp, item.Left)
			}
			if item.Right != nil {
				tmp = append(tmp, item.Right)
			}
		}
		levelArr = tmp
		level++
	}
}


func PrintArrOrder (arr []*TreeNode, normal bool) {
	if normal {
		for i:=0;i< len(arr);i++ {
			println(arr[i].Value)
		}
	} else {
		for i:=len(arr)-1;i>=0 ;i-- {
			println(arr[i].Value)
		}
	}
}


var leftBoarder = make(map[int]*TreeNode)
var rightBoarder = make(map[int]*TreeNode)
var leaves = make([]*TreeNode, 0)

func SetBoarder (root *TreeNode, h int) {
	if root == nil {
		return
	}

	if _, ok := leftBoarder[h]; !ok {
		leftBoarder[h] = root
		rightBoarder[h] = nil
	}

	if root != leftBoarder[h] {
		rightBoarder[h] = root
	}

	SetBoarder(root.Left, h+1)
	SetBoarder(root.Right, h+1)
}

func SetLeaves (root *TreeNode, h int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil && leftBoarder[h] != root && rightBoarder[h] != root {
		leaves = append(leaves, root)
	}

	SetLeaves(root.Left, h+1)
	SetLeaves(root.Right, h+1)
}

func PrintTreeBoarder (root *TreeNode) {
	SetBoarder(root, 0)
	SetLeaves(root, 0)
	for i:=0; i<len(leftBoarder); i++ {
		if leftBoarder[i] != nil {
			fmt.Println(leftBoarder[i].Value)
		}
	}
	for i:=0; i < len(leaves); i++ {
		if leaves[i] != nil {
			fmt.Println(leaves[i].Value)
		}
	}
	for i:=len(rightBoarder)-1; i>=0; i-- {
		if rightBoarder[i] != nil {
			fmt.Println(rightBoarder[i].Value)
		}
	}
}


func Max(a, b int) int {
	if a >=b {
		return a
	} else {
		return b
	}
}
func Min(a, b int) int {
	if a <=b {
		return a
	} else {
		return b
	}
}
func abs (a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

//任意路径最大和不单单是从root 出发的链路
//这种类型其实都是高度的变种后续遍历返回左右子树的最优解
func MaxGainIntree (root *TreeNode) int64 {
	maxSum := int64(0)
	MaxGainInTreeDFS(root, &maxSum)
	return maxSum
}

//树中的最大和 任意路径
func MaxGainInTreeDFS (root *TreeNode, MaxSum *int64) int64 {
	if root == nil {
		return 0
	}
	leftMaxSum := MaxGainInTreeDFS(root.Left, MaxSum)
	if leftMaxSum < 0 {
		leftMaxSum = 0
	}
	rightMaxSum := MaxGainInTreeDFS(root.Right, MaxSum)
	if rightMaxSum < 0 {
		rightMaxSum = 0
	}

	singleMax := int64(Max(int(rightMaxSum),int(leftMaxSum)))+root.Value
	curMaxSum := rightMaxSum + leftMaxSum + root.Value
	if curMaxSum > *MaxSum {
		*MaxSum = curMaxSum
	}
	return singleMax
}
//二叉树是否存在一个路径和==target的路径
func HasPathSum(root *TreeNode, targetSum int) bool {
	return HasPathSumDFS(root,targetSum,0)
}

func HasPathSumDFS (root *TreeNode, targetSum int, sum int) bool {
	if root == nil {
		return false
	}
	sum += int(root.Value)
	if sum == targetSum {
		return true
	}
	hasLeft := HasPathSumDFS(root.Left, targetSum, sum)
	hasRight := HasPathSumDFS(root.Right, targetSum, sum)
	//sum -= int(root.Value)
	return hasLeft||hasRight
}


// 单条路径 没必要从根节点触发
func HasPathSumAnyWay(root *TreeNode, targetSum int) bool {
	return HasPathSumDFSWay(root,targetSum,0)
}

func HasPathSumDFSWay (root *TreeNode, targetSum int, sum int) bool {
	if root == nil {
		return false
	}

	if sum + int(root.Value) == targetSum {
		return true
	}
	hasLeft1 := HasPathSumDFS(root.Left, targetSum, sum)
	hasRight1 := HasPathSumDFS(root.Right, targetSum, sum)
	hasLeft := HasPathSumDFS(root.Left, targetSum, sum + int(root.Value))
	hasRight := HasPathSumDFS(root.Right, targetSum, sum + int(root.Value))

	//sum -= int(root.Value)
	return hasLeft||hasRight||hasLeft1||hasRight1
}

func CountCompleteNodes(root *TreeNode) int {
	level := 0
	tmpRoot := root
	for tmpRoot!= nil {
		level++
		tmpRoot = tmpRoot.Left
	}
	//level 是二叉树的层数，则整颗二叉树在2^(n-1)~2^n-1
	low := 1<<(level-1)
	high := (1<<level)-1
	var ans = 0
	for low <=high {
		mid := (low+high)/2
		if isExistInCompeleteTree(root,level,mid) {
			low = mid +1
			ans = mid
		} else {
			high = mid-1
		}
	}
	return ans
}
func isExistInCompeleteTree (root *TreeNode, level,k int) bool {
	testNum := 1<<(level-1)
	tmpRoot := root
	for i:=0;i<level-1;i++ {
		testNum = testNum >> 1
		if tmpRoot == nil {
			break
		}
		if (k &testNum) > 0 {
			tmpRoot = tmpRoot.Right
		} else {
			tmpRoot = tmpRoot.Left
		}
	}
	if tmpRoot == nil {
		return false
	}
	return true
}

/*
337. 打家劫舍 III
小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。

除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。

给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。

简化一下这个问题：一棵二叉树，树上的每个点都有对应的权值，每个点有两种状态（选中和不选中），问在不能同时选中有父子关系的点的情况下，能选中的点的最大权值和是多少
 */


func RobMaxTree (root *TreeNode) int64 {
	n := make(map[*TreeNode]int64)
	y := make(map[*TreeNode]int64)

	n[nil] = 0
	y[nil] = 0
	RobMaxDfs(root, n,y)
	return int64(math.Max(float64(n[root]),float64(y[root])))
}

func RobMaxDfs (root *TreeNode, n,y map[*TreeNode]int64) {
	if root == nil {

		return
	}
	RobMaxDfs(root.Left, n,y)
	RobMaxDfs(root.Right, n,y)

	n[root] = int64(math.Max(float64(n[root.Left]), float64(y[root.Left])) + math.Max(float64(n[root.Right]), float64(y[root.Right])))
	y[root] = root.Value + n[root.Left] + n[root.Right]
}


/*
129. 求根节点到叶节点数字之和
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。
 */


func SumByTreePathNumberDfs (root *TreeNode, cur int64) int64 {
	if root == nil {
		return 0
	}
	cur = cur*10 + root.Value
	if root.Left == nil && root.Right == nil {
		//到叶子节点了
		return cur
	} else {
		return SumByTreePathNumberDfs(root.Left, cur) + SumByTreePathNumberDfs(root.Right, cur)
	}
}
/*
二叉树路劲和
*/
func SumAllRoad (root *TreeNode, cur int64) int64 {
	if root == nil {
		return 0
	}
	cur += root.Value
	if root.Left == nil && root.Right == nil {
		return cur
	} else {
		return SumAllRoad(root.Left,cur) + SumAllRoad(root.Right,cur)
	}
}

/*
是否镜像
 */
func IsMirrorTree (p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Value == q.Value && IsMirrorTree(p.Left, q.Right) && IsMirrorTree(p.Right, q.Left)
}
