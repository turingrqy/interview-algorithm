package greed

import (
	"math"
	"sort"
)

/*
45. 跳跃游戏 II
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。
典型的贪心算法
跳跃最少次数
 */

func GetMinStepJump (nums []int) int {
	curPos := 0
	steps := 0

	for curPos < len(nums) {
		steps++
		if curPos + nums[curPos] >= len(nums) {
			break
		}
		max := 0
		maxPos := 0
		for i := curPos + 1; i <= curPos+nums[curPos]; i++ {
			if i+nums[i] > max {
				max = i+nums[i]
				maxPos = i
			}
		}
		curPos = maxPos
	}
	return steps
}

/*
桌子 容纳的人数
人 人数 金额
求最多能赚多少钱
 */

type TablesType []table

type table struct {
	Cap int
	Busy bool
}
func (t TablesType) Len() int {
	return len(t)
}
func (t TablesType) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t TablesType) Less (i, j int) bool {
	if t[i].Cap < t[j].Cap {
		return true
	}
	return false
}

type PeoplesType []people
type people struct {
	Count int
	Money int
}

func (p PeoplesType) Len() int {
	return len(p)
}
func (p PeoplesType) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PeoplesType) Less (i, j int) bool {
	if p[i].Money < p[j].Money {
		return true
	} else if p[i].Money == p[j].Money && p[i].Count > p[j].Count {
		return true
	}
	return false
}

func DistributeTables (tables []table, peoples []people) int {
	sortedTables := TablesType(tables)
	sortedPeoples := PeoplesType(peoples)
	sort.Sort(sortedTables)
	sort.Sort(sort.Reverse(sortedPeoples))
	maxMoney := 0
	//从人少钱多的开始遍历
	for i := 0; i < len(sortedPeoples); i++ {
		// 从小桌开始遍历
		for j:= 0; j<len(sortedTables); j++ {
			if sortedTables[j].Cap >= sortedPeoples[i].Count && !sortedTables[j].Busy {
				maxMoney += sortedPeoples[i].Money
				sortedTables[j].Busy = true
			}
		}
	}
	return maxMoney
}

/*763. 划分字母区间
字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
*/

func PartitionLabelsStr (text string) []int {
	res := make([]int, 0, 0)
	uniqMap := make(map[byte]int)
	for i:=0; i<len(text); i++ {
		uniqMap[text[i]] = i
	}

	start := 0
	end := 0

	for i:=0; i < len(text); i++ {
		end = int(math.Max(float64(uniqMap[text[i]]), float64(end)))
		if end == i {
			res = append(res, i-start+1)
			start = end+1
			end = 0
		}
	}
	return res
}

// todo 最长递增子数组 滑动窗口









