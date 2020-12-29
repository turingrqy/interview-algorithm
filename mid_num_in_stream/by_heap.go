package mid_num_in_stream

import (
	"fmt"
	"renqiyang/interview/bublesort"
)
//cong数值构建的时候pas 从后万千传 压入排序传0
var maxHeap []int64
var minHeap []int64
var currentCount = 0



func GetMidNumInStream (input []int64)  {
	currentArr := []int64{}
	for _,val :=range input {
		currentArr = append(currentArr, val)
		PushToHeap(val)
		if currentCount%2 == 0 {
			bublesort.Bublesort(currentArr)
			fmt.Println(currentArr)
			fmt.Println((float64(minHeap[0])+float64(maxHeap[0]))/2)
		} else {
			bublesort.Bublesort(currentArr)
			fmt.Println(currentArr)
			fmt.Println(minHeap[0])
		}
	}
}

func PushToHeap (n int64) {

	if currentCount%2 == 0 {
		//push 到最小堆 （先去最大堆，在压入最小堆）
		maxHeap = append([]int64{n}, maxHeap...)
		adjustMaxHeap(maxHeap,0)
		Max := maxHeap[0]
		if len(maxHeap) == 1 {
			maxHeap = []int64{}
		} else {
			maxHeap = maxHeap[1:]
			adjustMaxHeap(maxHeap, 0)
		}
		minHeap = append([]int64{Max}, minHeap...)
		adjustMinHeap(minHeap,0)

	} else {
		minHeap = append([]int64{n}, minHeap...)
		adjustMinHeap(minHeap,0)

		Min := minHeap[0]
		if len(minHeap) == 1 {
			minHeap = []int64{}
		} else {
			minHeap = minHeap[1:]
			adjustMinHeap(minHeap, 0)
		}
		maxHeap = append([]int64{Min}, maxHeap...)
		adjustMaxHeap(maxHeap,0)
	}
	currentCount++
}


func adjustMinHeap (arr []int64, pos int) {
	node := pos
	len := len(arr)
	for node < len {
		child := 0
		if 2*node+2 < len {
			if arr[2*node+2] < arr[2*node+1] {
				child = 2*node+2
			} else {
				child = 2*node+1
			}
		} else if 2*node+1 < len {
			child = 2*node+1
		}

		if child > 0 && arr[child] < arr[node] {
			arr[child],arr[node] = arr[node],arr[child]
			node = child
		} else {
			break
		}
	}
}

func adjustMaxHeap (arr []int64, pos int) {
	node := pos
	len := len(arr)
	for node < len {
		child := 0
		if 2*node+2 < len {
			if arr[2*node+2] > arr[2*node+1] {
				child = 2*node+2
			} else {
				child = 2*node+1
			}
		} else if 2*node+1 < len {
			child = 2*node+1
		}

		if child > 0 && arr[child] > arr[node] {
			arr[child],arr[node] = arr[node],arr[child]
			node = child
		} else {
			break
		}
	}
}
