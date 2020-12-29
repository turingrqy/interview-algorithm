package selected_sort

func SelectedSort (arr []int64) {
	length := len(arr)
	for i:=0; i< length; i++ {
		maxIndex := i

		for j:=i; j< length; j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}
		arr[i],arr[maxIndex] = arr[maxIndex],arr[i]
	}
}
