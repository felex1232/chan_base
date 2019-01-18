package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rq = make([]Request, 0, 100)

const wc = 2

func main() {
	quit := make(chan bool)
	request := make(chan *Request, 1000)

	for i := 0; i < 3; i++ {
		rq = append(rq, Request{int(time.Now().UTC().UnixNano()), "R", func() []int {
			t := []int{}
			for i, c := 0, randInt(0, 5); i < c; i++ {
				t = append(t, int(time.Now().UTC().UnixNano()))
			}
			return t
		}()})
	}

	go Serve(request, quit)
	for {
		fmt.Println("lenrq", len(rq))
		for i := 0; len(rq) > 0 && i <= wc; i++ {
			r := rq[0]
			rq = rq[1:]
			request <- &r
		}
		time.Sleep(time.Microsecond)
	}

}

type Request struct {
	id   int
	msg  string
	jobs []int
}

func (r *Request) String() string {
	return fmt.Sprintf("request id: %d", r.id)
}

func Serve(request chan *Request, quit chan bool) {
	for i := 0; i < wc; i++ {
		id := i
		go handle(request, id)
	}
	<-quit
}

func handle(queue chan *Request, id int) {
	fmt.Println("handle start", id)
	for r := range queue {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("done", id, r)
		for _, t := range r.jobs {
			fmt.Println("job", t)
			rq = append(rq, Request{t, "Job", func() []int {
				x := []int{}
				for i, c := 0, randInt(1, 3); i < c; i++ {
					x = append(x, i)
				}
				return x
			}()})
		}
	}
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}