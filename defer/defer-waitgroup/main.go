package main

import "sync"

func main() {
	// wait group helps to notify the user when all the goroutine execution has finished
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				println(err.(string))
			}
		}()

		// cause deadlock
		// justPanic()

		// solution
		defer justPanic()
		wg.Done()
	}()

	// this will block until all the number specified in go routine has called wg.Done()
	wg.Wait()

}

func justPanic() {
	panic("I am panicking")
}
