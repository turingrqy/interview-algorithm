package insert_sort

func InsertSort(values []int64) {
	length := len(values)
	if length <= 1 {
		return
	}

	for i:=1; i<length;i++ {
		base := values[i]
		j:=i-1

		for ;j>=0&& values[j] < base;j-- {
			values[j+1] = values[j]
		}

		if j+1 != i {
			values[j+1] = base
		}
	}
}
