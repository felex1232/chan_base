package main

import "fmt"

func main() {
	//chan_buff()
	chan_nobuff()
}

/*
  channel with buffer
*/
func chan_buff() {
	var ch = make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<- ch)
}

/*
  channel with no buffer
*/
func chan_nobuff() {
	var ch = make(chan int, 1)
	ch <- 1
	go func() {
		ch <- 2
	}()
	fmt.Println(<- ch)
	fmt.Println(<- ch)
}
