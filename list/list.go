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

// 判断链表是否有环，快慢指针，如果快指针比慢指针差一步则下一次迭代两者会相遇，当差两步一次迭代后会变成差一步再次迭代相遇 2不当差3步的时候
// 一次迭代两者变成差两步在迭代差一步，在迭代相遇一次类推永远会相遇
/*
2.求有环单链表的环长

 　　在环上相遇后，记录第一次相遇点为Pos，之后指针slow继续每次走1步，fast每次走2步。在下次相遇的时候fast比slow正好又多走了一圈，也就是多走的距离等于环长。

　　设从第一次相遇到第二次相遇，设slow走了len步，则fast走了2*len步，相遇时多走了一圈：

　　　　环长=2*len-len。
	这个也很好理解，从环上相同的点出发，快指针走完一圈 满指针才走完半圈 len/2 满指针在走半圈，正好走完一圈，此时快指针也刚好走完一圈，两者
再次在之前的点相遇，此时走的迭代的次数就是环形的长度
 */

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
/*
第一次相遇时，slow走的长度 S = LenA + x;

　　　　第一次相遇时，fast走的长度 2S = LenA + n*R + x;

　　　　所以可以知道，LenA + x =  n*R;　　LenA = n*R -x;
      所以 在第一次相遇后，两个指针分别从pos 和head 每次走一步 两者会在n*R -x 处再次相遇就是入口点
 */

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

//两个倒序组成链表的
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


//旋转链表

