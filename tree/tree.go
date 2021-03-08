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

func GetSumRoad (root *TreeNode, arr []*TreeNode, currentSum int64, expectedSum int64)  {
	if root == nil {
		return
	}

	currentSum += root.Value
	arr = append(arr, root)
	if currentSum == expectedSum && root.Right == nil && root.Left == nil {
		fmt.Println(arr)
	}

	if root.Left != nil {
		GetSumRoad(root.Left, arr, currentSum, expectedSum)
	}

	if root.Right != nil {
		GetSumRoad(root.Right, arr, currentSum, expectedSum)
	}

	if len(arr) > 1 {
		arr = arr[:len(arr)-1]
	} else {
		arr = []*TreeNode{}
	}
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

func PrintBorder (root *TreeNode) {

	if root == nil {
		fmt.Println("null")
	}

	levelArr := []*TreeNode{}
	levelArr = append(levelArr, root)
	for len(levelArr) > 0 {
		len := len(levelArr)
		fmt.Println(levelArr[0].Value, levelArr[len-1].Value)

		tmpArr := levelArr
		levelArr = []*TreeNode{}
		if tmpArr[0].Left != nil {
			levelArr = append(levelArr, tmpArr[0].Left)
		} else if tmpArr[0].Right != nil {
			levelArr = append(levelArr, tmpArr[0].Right)
		}

		if tmpArr[len-1].Right != nil {
			levelArr = append(levelArr, tmpArr[len-1].Right)
		} else if tmpArr[len-1].Left != nil  {
			levelArr = append(levelArr, tmpArr[len-1].Left)
		}
	}

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

func PrintLeftBorder (root *TreeNode) {
	levelArr := []*TreeNode{root}

	for len (levelArr) > 0 {
		tmp := []*TreeNode{}
		if len(levelArr) == 1 {
			println(levelArr[0].Value)
			if levelArr[0].Left != nil {
				tmp = append(tmp, levelArr[0].Left)
			}
			if levelArr[0].Right != nil {
				tmp = append(tmp, levelArr[0].Right)
			}
		} else {
			needPrint:=true
			for key,node :=range levelArr {
				if node.Left != nil {
					tmp = append(tmp, node.Left)
				}
				if node.Right != nil {
					tmp = append(tmp, node.Right)
				}
				if key != len(levelArr) -1 && needPrint {
					if node.Left == nil && node.Right == nil {
						print(node.Value)
					} else {
						print(node.Value)
						needPrint = false
					}
				}

			}
		}
		levelArr = tmp
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

var edgeRes [][2]*TreeNode
var leaves []*TreeNode

func SetEdge (root *TreeNode,h int) {
	if root == nil {
		return
	}

	if edgeRes[h][0] == nil {
		edgeRes[h][0] = root
	}

	if edgeRes[h][0] != nil && root != edgeRes[h][0] {
		edgeRes[h][1] = root
	}
	SetEdge(root.Left,h+1)
	SetEdge(root.Right,h+1)
}

func SetLeaves (root *TreeNode,h int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil && root != edgeRes[h][0] && edgeRes[h][1] != root {
		leaves = append(leaves, root)
	}
}

func PrintTreeBoader (root *TreeNode) {
	h := GetHeight(root)
	edgeRes = make([][2]*TreeNode, 0, h)
	SetEdge(root,0)
	SetLeaves(root,0)
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
//二叉树是否存在一个路径和==target的路径s
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



