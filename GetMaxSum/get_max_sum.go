package GetMaxSum
// 连续子数组最大和
func GetMaxSum (arr []int64) int64 {

	if len(arr) == 1 {
		return arr[0]
	}

	maxSum, maxhere := arr[0], arr[0]

	for i := 1; i< len(arr); i++ {
		maxhere += arr[i]
		if maxhere > maxSum {
			maxSum = maxhere
		} else if maxhere < 0{
			maxhere = 0
		}
	}
	return maxSum
}
