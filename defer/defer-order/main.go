package main

// go defer execution uses LIFO execution order
func main() {
	defer println("defer number 1")

	level1()
}

func level1() {
	defer println("defer number 2")

	// recovery
	defer func() {
		if err := recover(); err != nil {
			println("recovering in progress")
		}
	}()
	defer println("defer number 3")

	level2()
}

func level2() {
	defer println("defer number 4")
	panic("panics here")
}
