package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func task1() {
	var counter int
	var mu sync.Mutex

	evenCh := make(chan int, 10)
	oddCh := make(chan int, 10)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range evenCh {
			if n%3 == 0 {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range oddCh {
			if n%33 == 0 {
				mu.Lock()
				counter--
				mu.Unlock()
			}
		}
	}()

	for i := 1; i <= 1000; i++ {
		if i%2 == 0 {
			select {
			case evenCh <- i:
			}
		} else {
			select {
			case oddCh <- i:
			}
		}
	}

	close(evenCh)
	close(oddCh)
	wg.Wait()

	fmt.Printf("[Завдання 1 — mutex]  counter = %d\n", counter)
}

func task2() {
	var counter int64

	evenCh := make(chan int, 10)
	oddCh := make(chan int, 10)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range evenCh {
			if n%3 == 0 {
				atomic.AddInt64(&counter, 1)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range oddCh {
			if n%33 == 0 {
				atomic.AddInt64(&counter, -1)
			}
		}
	}()

	for i := 1; i <= 1000; i++ {
		if i%2 == 0 {
			select {
			case evenCh <- i:
			}
		} else {
			select {
			case oddCh <- i:
			}
		}
	}

	close(evenCh)
	close(oddCh)
	wg.Wait()

	fmt.Printf("[Завдання 2 — atomic] counter = %d\n", atomic.LoadInt64(&counter))
}

func main() {
	task1()
	task2()
}
