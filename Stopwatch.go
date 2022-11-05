package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	time   time.Time
	splits []time.Time
}

func (sw *Stopwatch) Start() {
	sw.time = time.Now()
	sw.splits = make([]time.Time, 0)
}

func (sw *Stopwatch) SaveSplit() {
	sw.splits = append(sw.splits, sw.time)
}

func (sw Stopwatch) GetResults() []time.Duration {
	var retResults []time.Duration
	for _, splitTime := range sw.splits {
		retResults = append(retResults, splitTime.Sub(sw.time))
	}

	return retResults
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
