package show_half_in_arr

func GetShowHalfNumbyShowCount(arr[] int64) int64 {
	current := arr[0]
	showCount :=1
	for i:=1; i< len(arr); i++ {
		if arr[i] == current {
			showCount++
		}

		if arr[i] != current {
			showCount--
			if showCount == 0{
				current =  arr[i]
				showCount = 1
			}
		}
	}

	if showCount >= 1 {
		return current
	}

	return 0
}
