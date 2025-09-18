package main

import (
	"fmt"
	"sync"
)

func main_1() {
	//goroutine is a lightweight thread managed by the Go runtime
	//thread is managed by OS

	// race condition
	// access sum -> get value sum -> value + 1 -> save

	//Mutex block 1 variable khong cho thg nao access cho den khi unclock
	//waitGroup.Add(1) add them 1 goroutine muon chay, +1
	//waitGroup.Done() done so luong goroutine muon chay, -1
	//waitGroup.Wait() wait so luong goroutine muon chay, check done == 0 thi ket thuc
	var mu sync.Mutex
	var wg sync.WaitGroup

	sum := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			mu.Lock()
			sum += 1
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()
			mu.Lock()
			sum += 1
			mu.Unlock()
		}()
		// go1
		// go2
		// go3
	}

	// fmt.Println("Hello World", sum)
	wg.Wait()
	fmt.Println("Hello World", sum)
}
