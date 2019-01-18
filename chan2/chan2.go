package main

import (
	"fmt"
	"time"
)

func main() {
	range_chan()
	//select_chan()
}

/*
  channel with range
*/
func range_chan() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i ++ {
			ch <- i
		}
	}()

	go func() {
		for data := range ch {
			fmt.Println(data)
		}
	}()

	time.Sleep(2 * time.Second)
}


/*
  select with chan
*/
func select_chan() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i ++ {
			ch <- i
		}
	}()

	go func() {
		for {
			select {
			case data, ok := <-ch:
				if ok {
					fmt.Println("data ", data)
				} else {
					break
				}
			}
		}
	}()

	time.Sleep(5 * time.Second)
}