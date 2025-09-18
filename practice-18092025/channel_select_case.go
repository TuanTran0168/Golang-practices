package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func Sum(a, b int) (sum int) {
	sum = a + b
	return sum
}

func main() {
	var a int = 1844674407370955161
	b := 5000000000000000000
	fmt.Println("Sum:", Sum(a, b))

	ch1 := make(chan int)
	ch2 := make(chan int)

	// Goroutine 1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 100
	}()

	// Goroutine 2
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()

	// Nhận dữ liệu tuần tự
	// goroutine 1 xong truoc nhung goroutine 2 chua xong, nhận dữ liệu tuần tự

	// val2 := <-ch2
	// fmt.Println("Nhận từ ch2:", val2)

	// val1 := <-ch1
	// fmt.Println("Nhận từ ch1:", val1)

	// Nhận dữ liệu từ bất kỳ channel nào sẵn sàng
	for i := 0; i < 2; i++ {
		select {
		case val := <-ch2:
			fmt.Println("Nhận từ ch2:", val)
		case val := <-ch1:
			fmt.Println("Nhận từ ch1:", val)
		}
	}
}
