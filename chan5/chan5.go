package main

import "fmt"

var (
	Ch = make(chan int, 1)
)

func init() {
	go func() {
		Ch <- 1
	}()
}

type Worker struct {
	Wch  chan chan int
}

func NewWorker() *Worker{
	return &Worker{Wch: make(chan chan int)}
}

func (w *Worker) Start() {
	go func() {
		w.Wch <- Ch
	}()
}

func main() {
	w := NewWorker()
	w.Start()
	select {
	case data, _ := <-w.Wch:
		fmt.Println("data ", data)
		fmt.Println(<-data)
	}
}

