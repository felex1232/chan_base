package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			fmt.Println("Sub sending..", i+1)
			time.Sleep(time.Second)
		}

		ch <- 1
		fmt.Println("Sub sent")
	}()

	fmt.Println("Main Runing..")
	i := 0
Loop:
	for {
		select {
		default:
			i++
			fmt.Println("Main Waiting..", i)

		case <-ch:
			fmt.Println("Main Received")
			break Loop

		case <-time.After(time.Millisecond * 3):
			fmt.Println("Main Timeout")
			break Loop
		}
		time.Sleep(time.Millisecond * 100)
	}
}