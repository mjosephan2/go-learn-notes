package main

import "sync"

func main() {
	r := runOnce{}

	r.start()
	r.start()
}

type runOnce struct {
	sync.Once
	string
}

func (r *runOnce) start() {
	r.Do(func() {
		println("This function will only run once")
	})
	println("Run multiple times")
}
