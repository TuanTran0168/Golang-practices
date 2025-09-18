package main

import (
	"fmt"
)

func main1() {
	//channel la de giao tiep giua cac goroutine
	fmt.Println("Hello world")
	// create channel (type)
	channel := make(chan string) // channel unbuffered
	// => fatal error: all goroutines are asleep - deadlock!

	// Always use channel unbuffered with goroutine
	go func() {
		fmt.Println(<-channel) // vao sau nen block, nhan du lieu tu channel
		fmt.Println("nhan xong")
	}()

	channel <- "truyen ne" // vao truoc block, truyen du lieu vao channel

	fmt.Print("____________________________________________________\n")

	// create channel (type)
	channel2 := make(chan string, 1) // channel buffered
	channel2 <- "hello"
	fmt.Println(<-channel2)
	channel2 <- "hello123"
	fmt.Println(<-channel2)

	fmt.Print("____________________________________________________\n")
	sum := 0
	channel3 := make(chan int)
	for i := 0; i < 1000; i++ {

		go func() {
			sum += <-channel3
		}()
		channel3 <- 1
	}

	fmt.Println(sum)
	fmt.Print("____________________________________________________\n")
}
