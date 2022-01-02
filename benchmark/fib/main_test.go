package main

import "testing"

func TestFib(t *testing.T) {
	println(fib(30))
}

func TestFibMemo(t *testing.T) {
	println(fibMemo(30))
}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(i)
	}
}

func benchmarkFibMemo(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibMemo(i)
	}
}

func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }

func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }

func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

func BenchmarkFibMemo10(b *testing.B) { benchmarkFibMemo(10, b) }

func BenchmarkFibMemo20(b *testing.B) { benchmarkFibMemo(20, b) }

func BenchmarkFibMemo40(b *testing.B) { benchmarkFibMemo(40, b) }
