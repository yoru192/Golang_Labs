package main

import "fmt"

func generate() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int, 10)
	go func() {
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) int {
	total := 0
	for n := range in {
		total += n
	}
	return total
}

func main() {
	nums := generate()
	evens := filterEven(nums)
	squares := square(evens)

	result := sum(squares)

	fmt.Printf("Сума квадратів парних чисел від 1 до 100: %d\n", result)
}
