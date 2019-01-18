package main

import (
	"fmt"
	"time"
)

type MoneyChange struct {
	Uid         uint32
	BeforeMoney int64
	AfterMoney  int64
	Changed     int64
}

//var moneyEvents chan MoneyChange = make(chan MoneyChange, 1000)
var moneyEvents []chan MoneyChange

func init() {
	for i := 0; i < 200; i++ { // 200 是基于money_test.go测试出来的一个基准值
		event := make(chan MoneyChange, 1000)
		moneyEvents = append(moneyEvents, event)
		go eatEvents(event, i)
	}
}

func main() {
	me := MoneyChange{
		Uid:         12554,
		//BeforeMoney: 1000,
	}
	SendEvent(me)

	time.Sleep(5 * time.Second)
}

func SendEvent(e MoneyChange) {
	i := int(e.Uid % 200)
	if i >= 200 {
		fmt.Errorf("eatEvent err find null i:%d  events:%+v", i, e)
		return
	}

	event := moneyEvents[i]
	select {
	case event <- e:
		fmt.Printf("i %v", i)
	default:
		fmt.Printf("SendEvent chan full event %+v", e)
	}
}

func eatEvents(event chan MoneyChange, index int) {
	for e := range event {

		fmt.Printf("eatEvents change money for uid %v before %v after %v changed %v index %v", e.Uid, e.BeforeMoney, e.AfterMoney, e.AfterMoney-e.BeforeMoney, index)
	}
}