package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibMemo(n int) int {
	if n < 2 {
		return n
	}

	// initialize memo
	m := map[int]int{}
	return fibHelper(n, m)
}

func fibHelper(n int, m map[int]int) int {
	if n < 2 {
		return n
	}

	if mem, ok := m[n]; ok {
		return mem
	}
	result := fibHelper(n-2, m) + fibHelper(n-1, m)

	// add to memo
	m[n] = result
	return result
}
