package show_half_in_arr

func FindShowHafNumByPartition (arr []int64, low,high int) int64  {
	index := Partition(arr, low,high)
	mid := len(arr)/2
	for index != mid {
		if index < mid {
			index = Partition(arr, index+1, high)
		} else {
			index = Partition(arr, low, index-1)
		}
	}

	return arr[index]
}

func Partition(arr []int64, low, high int) int {
	base := arr[low]
	i,j:= low+1,high
	for i < j {
		for j > i && arr[j] >= base  {
			j--
		}

		for i < j && arr[i] <= base {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j],arr[i]
		}
	}

	arr[low] = arr[i]
	arr[i] = base

	return i
}
