package bublesort

func Bublesort (arr []int64) {
	len := len(arr)
	for i:=0;i<len;i++ {
		for j:= 1; j< len-i;j++ {
			if arr[j] < arr[j-1] {
				arr[j],arr[j-1] = arr[j-1],arr[j]
			}
		}
	}
}
