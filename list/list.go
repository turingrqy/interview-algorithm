package list

import (
	"fmt"
)

type ListNode2 struct {
	Key int
	Value int
	Next *ListNode2
	Prev *ListNode2
}

type ListNode struct {
	Value int64
	Next *ListNode
}

// 链表中倒数不使用额外空间 快慢指针
func GetLastKNode (head *ListNode, k int) *ListNode {
	quick := head
	for i := 1; i< k && quick != nil ; i++ {
		quick = quick.Next
	}

	if quick == nil {
		return nil
	}

	for quick.Next != nil {
		head = head.Next
		quick = quick.Next
	}

	return head
}

func ReverseList (head *ListNode) *ListNode {
	var prev *ListNode
	var pHead = head
	if head == nil || head.Next == nil {
		return head
	}
	var next = head.Next
	for next != nil {
		pHead.Next = prev
		prev = pHead
		pHead = next
		next = next.Next
	}
	pHead.Next = prev
	return pHead
}

func MergeSortList (head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	p1 := head1
	p2 := head2
	var tail *ListNode
	var newHead *ListNode
	if p1.Value <= p2.Value {
		newHead = p1
		tail = newHead
		p1 = p1.Next
		tail.Next = nil

	} else {
		newHead = p2
		tail = newHead
		p2 = p2.Next
		tail.Next = nil
	}
	for p1 != nil && p2 != nil {

		if p1.Value <= p2.Value {
			tail.Next = p1
			tail = p1
			p1 = p1.Next
			tail.Next = nil

		} else {
			tail.Next = p2
			tail = p2
			p2 = p2.Next
			tail.Next = nil
		}
	}

	if p1 != nil {
		tail.Next = p1
	}
	if p2 != nil {
		tail.Next = p2
	}
	return newHead
}
//有序链表中位数
func GetMidInList (head *ListNode) *ListNode {
	first := head
	second := head

	for second != nil && second.Next != nil  {
		first = first.Next
		second = second.Next.Next
	}

	return first
}

func GetRingLen (head *ListNode) int {
	first := head
	second := head
	len := 0
	for second != nil && second.Next != nil  {
		first = first.Next
		second = second.Next.Next
		if first == second {
			break
		}
	}
	if second == nil || second.Next == nil {
		return 0
	}
	for second != nil && second.Next != nil  {
		first = first.Next
		second = second.Next.Next
		len++
		if first == second {
			break
		}
	}
	return len
}

func GetEnterNodeInRing (head *ListNode) *ListNode {
	first := head
	second := head

	for second != nil && second.Next != nil  {
		first = first.Next
		second = second.Next.Next

		if first == second {
			tmp := head
			for tmp != second {
				tmp = tmp.Next
				second = second.Next
			}
			return tmp
		}
	}
	return nil
}

func SortList (head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	midNode,midNext := GetMidNodeInList(head)
	midNode.Next = nil
	leftList := SortList(head)
	rightList := SortList(midNext)

	return merge(leftList, rightList)
}

func GetMidNodeInList(head *ListNode) (*ListNode,*ListNode) {
	if head == nil {
		return nil, nil
	}

	if head.Next == nil {
		return head,nil
	}

	first := head
	second := head

	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next
	}

	return first, first.Next
}

func merge (head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	var newHead *ListNode
	var tail *ListNode
	p1,p2 := head1, head2
	if head1.Value > head2.Value {
		newHead = head2
		tail = newHead
		p2 = head2.Next
		tail.Next = nil
	} else {
		newHead = head1
		tail = newHead
		p1 = head1.Next
		tail.Next = nil
	}

	for p1 != nil && p2 != nil {
		if p1.Value < p2.Value {
			tail.Next = p1
			tail = p1
			p1 = p1.Next
			tail.Next = nil
		} else {
			tail.Next = p2
			p2 = p2.Next
			tail.Next = nil
		}
	}

	if p1 != nil {
		tail.Next = p1
	}
	if p2 != nil {
		tail.Next = p2
	}

	return newHead
}

//给定两个有序链表的头指针head1和head2，打印两个链表的公共部分
func PrintOrderedCommonList(head1 *ListNode, head2 *ListNode) {
	var p1, p2 = head1,head2
	for {
		if p1 == nil||p2==nil {
			return
		}
		if p1.Value == p2.Value {
			p1 = p1.Next
			p2 = p2.Next
			fmt.Println(p1.Value)
		}
		if p1.Value < p2.Value{
			p1 = p1.Next
		} else {
			p2=p2.Next
		}
	}
}

// 删除中间节点还是快慢指针
func RemoveMidNodeInList (head *ListNode) *ListNode {
	if head.Next==nil{
		return nil
	}
	if head.Next.Next == nil {
		return head.Next
	}
	curr := head
	quick := head.Next.Next
	for ;curr.Next!=nil && quick.Next.Next!=nil; {
		curr = curr.Next
		quick = quick.Next.Next
	}
	curr.Next = curr.Next.Next
	return head
}
// 翻转部分链表
func ReversePartList (head *ListNode, from,to int64) *ListNode {
	var fpre,tPos *ListNode
	p := head
	len := int64(1)
	for ;p!=nil;p=p.Next{
		if len == from-1 {
			fpre = p
		}
		if len == to+1 {
			tPos = p
		}
		len++
	}

	if from <1 || to > len {
		return head
	}
	var pre,curr *ListNode=tPos,nil
	if fpre == nil {
		curr = head
	} else {
		curr = fpre.Next
	}
	if curr.Next == nil {
		return head
	}

	quick := curr.Next
	for curr != tPos {
		curr.Next = pre
		pre=curr
		curr = quick
		quick = quick.Next
	}
	if fpre == nil {
		return pre
	}
	fpre.Next = pre
	return head
}

//两数相加
func SumTwoList (l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	var carry int64
	for l1 != nil || l2 != nil {
		v1,v2 :=int64(0),int64(0)
		if l1 != nil {
			v1 = l1.Value
		}
		if l2 != nil {
			v2 = l2.Value
		}
		sumTmp := v1+v2+carry
		carry = sumTmp/10
		realNum := sumTmp%10
		tmp := &ListNode{}
		tmp.Value =realNum
		tmp.Next = nil
		if head == nil {
			head = tmp
		}
		if tail == nil {
			tail = tmp
		} else {
			tail.Next = tmp
			tail = tmp
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return head
}

func PrintList (head *ListNode) {
	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}


//合并k个有序链表
//归并
func mergeKSortedLists(lists []*ListNode) *ListNode {
	if len(lists) == 0{
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	half := len(lists)/2
	left := mergeKSortedLists(lists[:half])
	right:= mergeKSortedLists(lists[half:])

	return mergeTwoSortedList(left, right)
}

func mergeTwoSortedList(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	var newHead *ListNode
	var tail *ListNode
	var p1 *ListNode
	var p2 *ListNode
	if head1.Value < head2.Value {
		newHead = head1
		tail = newHead
		p1 = head1.Next
		tail.Next = nil
	} else {
		newHead = head2
		tail = newHead
		p2 = head2.Next
		tail.Next = nil
	}

	for ;p1 != nil && p2 != nil; {
		if p1.Value < p2.Value {
			tail.Next = p1
			tail = p1
			p1 = p1.Next
			tail.Next = nil
		} else {
			tail.Next = p2
			tail = p2
			p2 = p2.Next
			tail.Next = nil
		}
	}
	if p1 != nil {
		tail.Next = p1
	}
	if p2 != nil {
		tail.Next = p2
	}
	return newHead
}


