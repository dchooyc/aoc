package m

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ArrMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		max = Max(max, arr[i])
	}
	return max
}

func ArrMin(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		min = Min(min, arr[i])
	}
	return min
}

func ArrSum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
