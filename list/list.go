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
	hair:= &ListNode{Next: head}
	quick := hair
	for i:=0; i<k; i++ {
		if quick == nil {
			break
		}
		quick = quick.Next
	}

	if quick == nil {
		return nil
	}

	for quick != nil  {
		quick = quick.Next
		hair = hair.Next
	}
	return hair
}

func ReverseList (head *ListNode) *ListNode {
	var prev *ListNode
	next := head.Next

	for next != nil {
		head.Next = prev
		prev = head
		head = next
		next = next.Next
	}
	head.Next = prev
	return head
}

func MergeSortedList (head1 *ListNode, head2 *ListNode) *ListNode {
	hair := &ListNode{Next: nil}
	tail := hair
	for head1 != nil && head2 != nil {
		if head1.Value < head2.Value {
			tail.Next = head1
			tail = head1
			head1 = head1.Next
			tail.Next = nil
		} else {
			tail.Next = head2
			tail = head2
			head2 = head2.Next
			tail.Next = nil
		}
	}

	if head1 != nil {
		tail.Next = head1
	}
	if head2 != nil {
		tail.Next = head2
	}
	return hair.Next
}

//有序链表中位数
func GetMidInList (head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	hair := &ListNode{Next: head}
	slow := hair
	fast := hair
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
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
	if head == nil {
		return nil
	}

	hair := &ListNode{Next: head}
	first := hair
	prev :=  hair
	second := hair
	for second != nil && second.Next != nil {
		prev = first
		first = first.Next
		second = second.Next.Next
	}

	prev.Next = first.Next
	first.Next = nil
	return hair.Next
}

// 翻转部分链表
func ReversePartedList (head *ListNode, from,to int) *ListNode {
	if to < from {
		return nil
	}
	hair := &ListNode{Next: head}
	pre := hair
	endNode := hair

	for i:= 0; i < to && endNode != nil; i++ {
		if i == from-1 {
			pre = endNode
		}

		endNode = endNode.Next
	}
	if endNode == nil {
		return nil
	}

	startNode := pre.Next
	suffix := endNode.Next
	pre.Next = nil
	endNode.Next = nil

	reverseList(startNode)
	pre.Next = endNode
	startNode.Next = suffix
	return hair.Next
}

/*
25. K 个一组翻转链表
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换
 */
func ReverseListByGroup (head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	prev := hair
	for prev != nil {
		tail := prev
		for i:=0; i < k && tail != nil; i++ {
			tail = tail.Next
		}
		if tail == nil {
			break
		}
		var suffix *ListNode
		suffix = tail.Next
		startNode := prev.Next
		prev.Next = nil
		reverseList(startNode)
		prev.Next = tail
		if startNode != nil {
			startNode.Next = suffix
		}
		prev =  startNode
	}
	return hair.Next
}
//头结点变为尾结点
func reverseList (head *ListNode)  {
	var prev *ListNode
	var pHead = head
	if head == nil || head.Next == nil {
		return
	}
	var next = head.Next
	for next != nil {
		pHead.Next = prev
		prev = pHead
		pHead = next
		next = next.Next
	}
	pHead.Next = prev
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

	return MergeSortedList(left, right)
}


func SortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	head1, head2 := curList(head)

	left := SortList(head1)
	right := SortList(head2)

	return MergeSortedList(left, right)
}

// 利用快慢指针 找到链表中点
func curList(head *ListNode)(*ListNode, *ListNode){
	hair :=  &ListNode{Next: head}
	slow := hair
	fast := hair
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	newHead := slow.Next
	slow.Next = nil
	return hair.Next, newHead
}

func DeleteSameNodeInList (head *ListNode) *ListNode {
	slow := head
	quick := head.Next
	for quick != nil {
		if slow.Value != quick.Value {
			slow.Next = quick
			slow = quick
		}
		quick = quick.Next
	}
	slow.Next = quick
	return head
}