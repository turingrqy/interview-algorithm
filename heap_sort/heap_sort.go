package heap_sort

/*完全二叉树 2*i + 1 和 2 * i + 2.   i的父节点下标为 (i-1)/2 非稳定排序 从小到大用最大堆
	topk 最大的用最小堆 堆的特性
大根堆定义
*/
//n*logn
func HeapSort(arr []int64) {
	buildHeap(arr)
	length := len(arr)
	for i:=length-1; i>=0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		adjustHeap(arr[:i], 0)
	}
}

func buildHeap (arr []int64) {
	length := len(arr)

	for i:=length-1; i>=0; i-- {
		adjustHeap(arr, i)
	}
}

func adjustHeap(arr []int64, pos int) {
	node := pos
	length := len(arr)

	for node < length {
		maxChildIndex := 0

		if 2*node+2 < length {
			if arr[2*node+2] < arr[2*node+1] {
				maxChildIndex = 2*node+1
			} else {
				maxChildIndex = 2*node+2
			}
		} else if 2*node+1 < length {
			maxChildIndex = 2*node+1
		}

		if maxChildIndex > 0 && arr[maxChildIndex] > arr[node] {
			arr[maxChildIndex], arr[node] = arr[node], arr[maxChildIndex]
			node = maxChildIndex
		} else {
			break
		}
	}
}

func adjustMinHeap (arr []int64, pos int) {
	node := pos
	len := len(arr)
	for node<len {
		minChild := 0
		if 2*node+2 < len {
			if arr[2*node+2] < arr[2*node+1] {
				minChild = 2*node+2
			} else {
				minChild = 2*node+1
			}
		} else if 2*node+1 < len {
			minChild = 2*node+1
		}

		if minChild > 0 && arr[minChild] < arr[node] {
			arr[minChild],arr[node] =  arr[node], arr[minChild]
			node = minChild
		} else {
			break
		}
	}
}

func buildMinHeap (arr []int64) {
	for i:= len(arr)-1; i>=0; i-- {
		adjustMinHeap(arr, i)
	}
}
//nlogk
func GetTopKByHeap (input []int64, k int) []int64 {
	heapArr := make([]int64, 0, k)
	read := 0
	for read=0;read<k;read++ {
		heapArr = append(heapArr, input[read])
	}

	buildMinHeap(heapArr)

	for ;read<len(input); read++ {
		if input[read] > heapArr[0] {
			heapArr[0] = input[read]
			adjustMinHeap(heapArr, 0)
		}
	}
	for i:=k-1; i>=0; i-- {
		heapArr[0],heapArr[i] = heapArr[i], heapArr[0]
		adjustMinHeap(heapArr[:i], 0)
	}
	return heapArr
}


