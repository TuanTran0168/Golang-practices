package main

import (
	"fmt"
	"time"
)

func sender(name string, ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s gửi: %d\n", name, i)
		ch <- i // gửi dữ liệu vào channel
		time.Sleep(time.Millisecond * 500)
	}
}

func mai1() {
	ch := make(chan int)

	// 2 goroutine gửi dữ liệu
	go sender("Goroutine 1", ch)
	go sender("Goroutine 2", ch)

	// Goroutine nhận dữ liệu
	go func() {
		total := 0
		for i := 0; i < 10; i++ { // tổng số dữ liệu là 5 + 5 = 10
			total += <-ch
			fmt.Println("Nhận:", total)
		}
	}()

	// Chờ các goroutine gửi xong
	time.Sleep(6 * time.Second)
}
