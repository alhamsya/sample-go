package number

func Fibonacci(n int) int {

	//stop condition
	if n == 0 {
		return 0
	}

	//stop condition if request input 1
	if n == 1 {
		return 1
	}

	result := Fibonacci(n-1) + Fibonacci(n-2)
	return result
}
