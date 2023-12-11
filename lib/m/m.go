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

func Gcd(num, den int) int {
	if num < den {
		num, den = den, num
	}
	rem := num % den
	for rem != 0 {
		num, den = den, rem
		rem = num % den
	}
	return den
}

func Lcm(a, b int) int {
	return (a * b) / Gcd(a, b)
}

func Max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func Gcd64(num, den int64) int64 {
	if num < den {
		num, den = den, num
	}
	rem := num % den
	for rem != 0 {
		num, den = den, rem
		rem = num % den
	}
	return den
}

func Lcm64(a, b int64) int64 {
	return (a * b) / Gcd64(a, b)
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

func ArrMax64(arr []int64) int64 {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		max = Max64(max, arr[i])
	}
	return max
}

func ArrMin64(arr []int64) int64 {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		min = Min64(min, arr[i])
	}
	return min
}

func ArrSum64(arr []int64) int64 {
	var sum int64
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
