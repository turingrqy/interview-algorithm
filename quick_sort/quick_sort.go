package quick_sort

//无序数组中的topk
//数组中取中位数 数组中大多数元素
func GetTopkbyPartition(arr []int64, k int) int {
	if k > len(arr) {
		return -1
	}
	low := 0
	high := len(arr)-1
	index := partitionMax(arr,low, high)
	for index != k-1 {
		if index > k-1 {
			high = index-1
			index = partitionMax(arr,low,high)
		} else {
			low = index+1
			index = partitionMax(arr,low,high)
		}
	}
	return 0
}

func partitionMax(arr []int64, high, low int) int {
	i,j := low,high
	for i<j {
		for ;j>i;j-- {
			if arr[j] > arr[low] {
				break
			}
		}

		for ;i<j;i++ {
			if arr[i] < arr[low] {
				break
			}
		}
		if i<j {
			arr[i],arr[j] = arr[j],arr[i]
		}
	}
	arr[i],arr[low] = arr[low], arr[i]
	return i
}
func partitionMin(arr []int64, high, low int) int {
	i,j := low,high
	for i<j {
		for ;j>i;j-- {
			if arr[j] < arr[low] {
				break
			}
		}

		for ;i<j;i++ {
			if arr[i] > arr[low] {
				break
			}
		}
		if i<j {
			arr[i],arr[j] = arr[j],arr[i]
		}
	}
	arr[i],arr[low] = arr[low], arr[i]
	return i
}