package main

import (
	"fmt"
	"renqiyang/interview/begpack"
	"renqiyang/interview/list"
	"renqiyang/interview/other"
	"renqiyang/interview/tree"
)

func main() {
	/*arr := []int64{1,5,7,4,3,11,15,12,13,14,20,18,17,9,2,8}


	quickSort(arr,0,15)
	fmt.Println(arr)
	arr = []int64{1,5,7,4,3,11,15,12,13,14,20,18,17,9,2,8}
	quickSortTopK(arr,0,len(arr)-1,3)
	fmt.Println(arr)
	arr = []int64{1,5,7,4,3,11,15,12,13,14,20,18,17,9,2,8}
	arr = mergeSort(arr)
	fmt.Println(arr)
	arr2Dimantion := [][]int64{
		{2,4,6,8},
		{1,3,5,7,9},
	}

	arr = merge2TwoDimationArr(arr2Dimantion)
	fmt.Println(arr)
	arr = []int64{1,5,7,4,3,11,15,12,13,14,20,18,17,9,2,8}
	heap_sort.HeapSort(arr)
	fmt.Println(arr)
	arr = []int64{1,5,7,4,3,11,15,12,13,14,20,18,17,9,2,8}
	arr = heap_sort.GetTopKByHeap(arr,10)
	fmt.Println(arr)
	arr = []int64{1,5,7,4,3,11,15,12,13,14,20,18,17,9,2,8}
	bublesort.Bublesort(arr)
	fmt.Println(arr)
	arr = []int64{1,5,7,4,3,11,15,12,13,14,20,18,17,9,2,8}
	selected_sort.SelectedSort(arr)
	fmt.Println(arr)
	arr = []int64{1,5,7,4,3,11,15,12,13,14,20,18,17,9,2,8}
	insert_sort.InsertSort(arr)
	fmt.Println(arr)
	arr = []int64{1,5,5,4,3,5,15,5,13,14,5,5,17,9,5,5}
	showMaxNum:= show_half_in_arr.FindShowHafNumByPartition(arr, 0, 15)
	fmt.Println(showMaxNum)
	showMaxNum=show_half_in_arr.GetShowHalfNumbyShowCount(arr)
	fmt.Println(showMaxNum)
	arr = []int64{1,5,7,4,3,11,15,12,13,14,20}
	mid_num_in_stream.GetMidNumInStream(arr)
	arr = []int64{1,3,5,7}
	index :=binary_search.BinarySearch(arr,2)
	fmt.Println(index)
	arr = []int64{7,7,1,3,5,6}
	index =binary_search.BinarySearchRotate(arr,7)
	fmt.Println(index)
	index = binary_search.BinarySearchMinInRotate(arr)
	fmt.Println(index)
	arr = []int64{1,3,5,7,9}
	root := tree.CreateTreeOrderedArr(arr)
	tree.PreorderTree(root)
	fmt.Println("------------")
	tree.PreOrderStack(root)
	fmt.Println("------------")
	tree.MidOrderTree(root)
	fmt.Println("------------")
	tree.MidOrderTreeStack(root)
	fmt.Println("------------")
	tree.LastOrderTree(root)
	fmt.Println("------------")
	tree.LastOrderTreeStack(root)
	fmt.Println("------------")
	tree.WideOrderTree(root)
	fmt.Println("------------")
	height := tree.GetHeight(root)
	fmt.Println(height)
	fmt.Println("------------")
	isBalance,_ := tree.IsBalanceTree(root)
	fmt.Println(isBalance)
	res := tree.GetLeveledNode(root)
	for _,arr := range res {
		fmt.Println(arr)
	}
	num := tree.GetLevelNodeNum(root, 2)
	fmt.Println(num)
	fmt.Println(tree.IsSearchTree(root,math.MinInt64, math.MaxInt64))
	fmt.Println(tree.IsSearchTreeMidOrder(root,math.MinInt64))
	fmt.Println(tree.IsSearchTreeMidOrderStack(root))
	fmt.Println(tree.FindTreeMidOrderNextNode(root.Left.Right))
	fmt.Println(tree.FindCommonAccesstorIBST(root, root.Left.Right, root.Right.Right))
	fmt.Println(tree.FindCommonAccesstorIBSTParent(root, root.Left.Right, root.Right.Right))
	fmt.Println(tree.FindCommonAccesstor(root, root.Right, root.Right.Right))
	num = 0
	tree.GetRootToNode(root, root.Right.Right, &num)
	fmt.Println((num))
	println("----------")
	arr1 := []*tree.TreeNode{}
	tree.GetRootToNodeRoad(root, root.Right.Right, &arr1)
	for i:= 0; i< len(arr1); i++ {
		fmt.Println(arr1[i].Value)
	}
	println("----------")
	arr1 = tree.GetNodeInKlevel(root, 3)
	for i:= 0; i< len(arr1); i++ {
		fmt.Println(arr1[i].Value)
	}
	println("----------")
	tree.GetSumRoad(root,[]*tree.TreeNode{}, 0, 9)
	node := binary_search.HashBinarySearch([]int64{1,3,5,7}, 0)
	phead := tree.ConvertTree2Link(root)
	fmt.Println(phead)
	fmt.Println(node)

	plhead := &list.ListNode{}
	plhead.Value = 0
	tmp := plhead
	for i:=1; i<6;i++ {
		tmp.Next = &list.ListNode{
			Value: int64(i),
			Next:  nil,
		}
		tmp = tmp.Next
	}
	listNode := list.GetLastKNode(plhead, 3)
	fmt.Println(listNode.Value)
	newHead := list.ReverseList(plhead)
	node0 := &list.ListNode{
		0,
		nil,
	}
	node1 := &list.ListNode{
		1,
		nil,
	}
	node2 := &list.ListNode{
		2,
		nil,
	}
	node3 := &list.ListNode{
		3,
		nil,
	}
	node4 := &list.ListNode{
		4,
		nil,
	}
	node0.Next = node2
	node2.Next = node4
	node1.Next = node3
	newHead = list.MergeSortList(node0,node1)
	fmt.Println(newHead)
	midNode := list.GetMidInList(newHead)
	fmt.Println(midNode)
	node4.Next = node1
	//ringLen := list.GetRingLen(newHead)
	//             hhhhhhd fmt.Println(ringLen)
	enterNode := list.GetEnterNodeInRing(newHead)
	fmt.Println(enterNode)

	arrTest := []int64{3,-11,2,-1,-10,5,6,-20}
	resArrTest := binary_search.GetLongestUpSubArr(arrTest)
	fmt.Println(resArrTest)
	maxSum := GetMaxSum.GetMaxSum(arrTest)
	fmt.Println(maxSum)

	arr = []int64{1,3,5,7,9,11,13,15,17,19,20}
	root = tree.CreateTreeOrderedArr(arr)
	tree.PrintBorder(root)

	fmt.Println(fmt.Sprintf("value=%v", map[string]interface{}{"test":"test"}))*/
	/*head := &list.ListNode{}
	head.Value = 0
	head.Next = nil
	tail := head
	for i:=1;i<7;i++{
		tmp := &list.ListNode{}
		tmp.Value = int64(i)
		tail.Next = tmp
		tail = tmp
	}
	newList := list.ReversePartList(head,2,5)
	p:=newList
	for ;p!=nil;p=p.Next{
		fmt.Println(p.Value)
	}*/

	/*s:="aaabcefsdfedddabcefdklcdsa"
	noduplen := other.GetNodupSubStrLen(s)
	fmt.Println(fmt.Sprintf("no dup num=%d", noduplen))*/
	/*num1_1 := list.ListNode{
		3,
		nil,
	}
	num1_2 := list.ListNode{
		6,
		nil,
	}
	num1_3 := list.ListNode{
		5,
		nil,
	}
	num2_1 := list.ListNode{
		8,
		nil,
	}
	num2_2 := list.ListNode{
		8,
		nil,
	}
	num1_1.Next = &num1_2
	num1_2.Next = &num1_3
	num2_1.Next = &num2_2
	sum := list.SumTwoList(&num1_1, &num2_1)
	//fmt.Println(fmt.Sprintf("list sum=%d",sum))
	list.PrintList(sum)

	res := binary_search.GetMidIntwosortedArr([]int64{1,3,5,7,9},[]int64{2,4,6,8,10,11})
	fmt.Println("GetMidIntwosortedArr=",res)

	str := other.DecodeStr("a2[a2[ab]d]c")
	fmt.Println(str)
	dupArr := []int{1,1,1,2,2,2,3,3,4,4,5}
	other.DeleteDupInArr(dupArr)
	fmt.Println(fmt.Sprintf("arr=%v", dupArr))
	other.Permute([]int{1,2,3,4,5})
	res1 := binary_search.FindPeek([]int{1,2,3,4,5,6})
	res2 := binary_search.FindPeek([]int{6,5,4,3,2,1})
	res3 := binary_search.FindPeek([]int{1,2,3,2,1})
	fmt.Println(fmt.Sprintf("res1=%d, res2=%d, res3=%d", res1, res2, res3))
	maxRain := other.CatchRainNormal([]int{4,2,0,3,2,5})
	fmt.Println(maxRain)
	coins := []int{1, 2, 5}
	fmt.Println(fmt.Sprintf("min change num=%d", other.CoinChangeNormal(coins,11)))
	fmt.Println(fmt.Sprintf("min change num=%d", other.CoinChangeDp(coins,11)))
	//fmt.Println(fmt.Sprintf("longest upnum=%d", other.GetLongestUpNum([]int{10,9,2,5,3,7,101,18})))
	a,b,c,sum1 := binary_search.GetClosestThreeNumSum([]int64{-1,2,1,-4},1)
	fmt.Println(fmt.Sprintf("closest a b c sum=%d,%d,%d,%d",a,b,c,sum1))
	other.GetAllSubset([]int64{1,2,3})
	other.GenerateParenthesis(3)
	intMap := [][]int{
		[]int{1,1,0,0,0},
		[]int{1,1,0,0,0},
		[]int{0,0,1,0,0},
		[]int{0,0,0,1,1},
	}
	islandNum := other.GetIsLandNum(intMap)
	fmt.Println(fmt.Sprintf("islandNum=%d",islandNum))
	singleDup := binary_search.GetOnceDupNumInArr([]int{1,3,4,2,2})
	fmt.Println(fmt.Sprintf("single dup=%d",singleDup))
	//other.FindCombineSumEqTarget([]int{2,3,6,7}, 7)
	resDFS := other.FindCombineSumEqTarget([]int{2,3,6,7}, 7)
	fmt.Println("组合综合 dfs=", resDFS)
	resDp := begpack.GetCombineSumEqTargetDp([]int{2,3,6,7}, 7)
	fmt.Println("组合综合 dp=", resDp)
	minSum:=other.GetMinRoadSumGrid([][]int{
		[]int{1,3,1},
		[]int{1,1,5},
		[]int{4,2,1},
	},3,3)
	fmt.Println(fmt.Sprintf("minRoadSum=%d",minSum))
	resrow,resCol := binary_search.FindInTwoDimensionSorted([][]int{
		[]int{1,4,7,11,15},
		[]int{2,5,8,12,19},
		[]int{3,6,9,16,22},
		[]int{10,13,14,17,24},
		[]int{18,21,23,26,30},
	},5)
	fmt.Println(fmt.Sprintf("resrow=%d resCol=%d",resrow,resCol))
	resMax := recursion_dynamic.GetMaxMulti([]int{-2,0,-1})
	fmt.Println(fmt.Sprintf("resMax=%d",resMax))
	newstr := other.MinWindow("ADOBECODEBANC", "ABC")
	fmt.Println(fmt.Sprintf("newstr=%s",newstr))
	recursion_dynamic.GetNumberofLIS([]int{1,3,5,4,7,6})
	recursion_dynamic.GetLongetLIS([]int{1,3,5,4,7,6})
	rowIndex,colIndex := binary_search.FindInTwoDimensionStrictSorted([][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	},5)
	fmt.Println(fmt.Sprintf("rowIndex=%d,colIndex=%d",rowIndex,colIndex))
	index1 := binary_search.BinarySearchMinInRotate([]int64{5,5,7,8,1,2,5,5})
	fmt.Println(fmt.Sprintf("BinarySearchMinInRotate index=%d",index1))
	recursion_dynamic.MultiByStokMaxProfit([]int{7,1,5,3,6,4})
	testGrid := [][]int{
		[]int{3,2,3,4,5},
		[]int{5,4,3,2,1},
		[]int{1,2,3,3,3},
		[]int{2,3,3,2,1},
	}
	other.SpreadGrid(testGrid, 2,3,2,2)
	println(fmt.Sprintf("testGrid = %+v", testGrid))



	other.JustifyWord([]string{"This", "is", "an", "example", "of", "text", "justification."},16)
	resMulti := other.MultiExceptself([]int{1,2,3,4,5})
	fmt.Println(resMulti)
	testGrid1 :=[][]int{
		[]int{3,2,3,4},
		[]int{5,4,3,2},
		[]int{1,2,3,3},
		[]int{2,3,3,2},
	}
	other.RotateMatrix(testGrid1)
	fmt.Println(fmt.Sprintf("testGrid1=%+v", testGrid1))
	tmp1 := &list.ListNode {
		3,
		nil,
	}
	tmp2 := &list.ListNode{
		2,
		nil,
	}
	tmp3 := &list.ListNode {
		5,
		nil,
	}
	tmp4 := &list.ListNode {
		1,
		nil,
	}
	tmp1.Next = tmp2
	tmp2.Next = tmp3
	tmp3.Next = tmp4
	newhead := list.SortList(tmp1)
	printTmp := newhead
	for printTmp != nil {
		fmt.Println(printTmp.Value)
		printTmp = printTmp.Next
	}
	resNew := binary_search.GetNoDupNumInSortedDoubleArr([]int{1,2,2,3,3,4,4,5,5,6,6,7,7,8,8})
	fmt.Println("resNew = ", resNew)*/
	num1_1 := list.ListNode{
		3,
		nil,
	}
	num1_2 := list.ListNode{
		5,
		nil,
	}
	num1_3 := list.ListNode{
		4,
		nil,
	}
	num1_4 := list.ListNode{
		2,
		nil,
	}
	num1_5 := list.ListNode{
		6,
		nil,
	}
	num1_6 := list.ListNode{
		1,
		nil,
	}
	num1_1.Next = &num1_2
	num1_2.Next = &num1_3
	num1_3.Next = &num1_4
	num1_4.Next = &num1_5
	num1_5.Next = &num1_6

	newHead := list.SortList(&num1_1)
	printList(newHead)
	other.GetLongestPalindrome("fdabadcedasdada")
	arr1 := []int{1,2,3,3,3,3,4,5}
	other.DeleteDupInSortedArr(&arr1)
	println(fmt.Sprintf("arr=%v",arr1))
	/*newHead := list.DeleteSameNodeInList(&num1_1)
	printList(newHead)
	testDupIntArr := []int{0,3,2,3,3,3}
	res := binary_search.GetDupInArr(testDupIntArr)
	println(fmt.Sprintf("GetDupInArr res=%d",res))

	num2_1 := list.ListNode{
		1,
		nil,
	}
	num2_2 := list.ListNode{
		2,
		nil,
	}
	num2_3 := list.ListNode{
		3,
		nil,
	}
	num2_4 := list.ListNode{
		4,
		nil,
	}
	num2_5 := list.ListNode{
		5,
		nil,
	}
	num2_6 := list.ListNode{
		6,
		nil,
	}
	num2_1.Next = &num2_2
	num2_2.Next = &num2_3
	num2_3.Next = &num2_4
	num2_4.Next = &num2_5
	num2_5.Next = &num2_6

	//reversed := list.ReversePartedList(&num2_1,2,4)
	reversed := list.ReverseListByGroup(&num2_1, 2)
	printList(reversed)

	println(other.DecodeStr("a2[a2[ab]]c"))
	resCombine := other.Subsets([]int{1,2,3})
	println(fmt.Sprintf("resCombine=%v",resCombine))

	println(other.GetMinRoadSumGridByDfs([][]int{{1,2,3},{4,5,6},{7,8,9}}))

	resMaxArr := other.GetFirstMaxFromRight([]int{1,5,3,6,4,8,9,10})
	fmt.Println(resMaxArr)
	other.CoverGrid(4, 0,1)

	upSet := other.GetAllSubUpSet([]int{4,6,7,7})
	println(fmt.Sprintf("upSet=%v", upSet))

	/*resCommon := recursion_dynamic.GetLongestCommonSubsequence("abcde", "bcde")
	println(resCommon)
	resArrCommon := recursion_dynamic.GetLongestSubArr([]int{1,4,5,6}, []int{4,5,6})
	println(resArrCommon)*/

	arr := []int64{1,3,5,7,9}
	root := tree.CreateTreeOrderedArr(arr)
	tree.GetSumRoad(root, []*tree.TreeNode{}, 9)
	tree.PrintTreeBoarder(root)
	println(tree.HasPathSumAnyWay(root,25))

	coins := []int{10,2,7,6,5}
	begpack.FindCombineSumEqTargetNoDup1(coins, 8)
	begpack.FindCombineSumEqTargetNoDupInRArr(coins, 8)
	TraverseMatrixArr := other.TraverseMatrix([][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}})
	println(TraverseMatrixArr)
	/*other.GetAllSubUpSet([]int{4,6,7,3,5})
	res := recursion_dynamic.GetMaxLengthOfLIS([]int{4,6,7,3,5})
	println(fmt.Sprintf("GetNumberofLIS=%v", res))
	resD := recursion_dynamic.GetLongestCommonSubsequence("abce","afce")
	fmt.Println(fmt.Sprintf("DecodeAZ1 count=%d", resD))*/
	rootSum := &tree.TreeNode{Value: 1}
	rootSum.Left = &tree.TreeNode{Value: 2}
	rootSum.Right = &tree.TreeNode{Value: 3}
	println(fmt.Sprintf("SumByTreePathNumberDfs=%d",tree.SumByTreePathNumberDfs(rootSum, 0)))
	sortArr := []int64{5,3,6,7}
	quickSort(sortArr, 0, 3)
	println(fmt.Sprintf("quickSort arr=%v", sortArr))
}


//快排 最重要的特性在partition，左边小又变大
func quickSort (arr []int64, low, high int) {

	if low < high {
		index := partition(arr, low, high)
		quickSort(arr, index+1, high)
		quickSort(arr, low, index-1)
	}
}

func partition(arr []int64, low,high int) int {
	standard := arr[low]
	i,j := low,high

	for i<j {
		for j > i && arr[j] >= standard {
			j--
		}

		for i < j && arr[i] <= standard {
			i++
		}

		if i<j {
			arr[i],arr[j] = arr[j],arr[i]
		}
	}

	arr[low] = arr[i]
	arr[i] = standard
	return i
}
//查找前k的数字
func quickSortTopK (arr []int64, low, high, k int) {
	index := partition(arr, low, high)
	for index != k -1 {
		if index < k-1 {
			index = partition(arr, index +1, high)
		}
		if index > k-1 {
			index = partition(arr, low, index-1)
		}
	}
}

func merge2TwoDimationArr (arr[][]int64) []int64 {
	if len(arr) == 1 {
		return arr[0]
	}

	if len(arr) == 0 {
		return []int64{}
	}
	mid := len(arr)/2
	left := merge2TwoDimationArr(arr[:mid])
	right := merge2TwoDimationArr(arr[mid:])

	return merge(left, right)
}
//归并的过程
func mergeSort(data []int64) []int64 {
	if len(data) <= 1 {
		return data
	}
	//递[归]
	middle := len(data) / 2
	//不断地进行左右对半划分
	left := mergeSort(data[:middle])
	right := mergeSort(data[middle:])
	//合[并]
	return merge(left, right)
}

func merge(left, right []int64) (result []int64) {
	l, r := 0, 0
	// 注意：[左右]对比，是指左的第一个元素，与右边的第一个元素进行对比，哪个小，就先放到结果的第一位，然后左或右取出了元素的那边的索引进行++
	for l < len(left) && r < len(right) {
		//从小到大排序.
		if left[l] > right[r] {
			result = append(result, right[r])
			//因为处理了右边的第r个元素，所以r的指针要向前移动一个单位
			r++
		} else {
			result = append(result, left[l])
			//因为处理了左边的第r个元素，所以r的指针要向前移动一个单位
			l++
		}
	}
	// 比较完后，还要分别将左，右的剩余的元素，追加到结果列的后面(不然就漏咯）。
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

func printList (head *list.ListNode) {
	for p:=head; p != nil; p = p.Next {
		println(p.Value)
	}
}



