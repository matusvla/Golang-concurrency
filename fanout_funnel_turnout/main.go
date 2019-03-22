package main

import (
	"fmt"
	"sync"
)

//Fanout takes one channel and distributes it to two
func Fanout(wg *sync.WaitGroup, In <-chan int, OutA, OutB chan int) {
	defer wg.Done()
	for data := range In {
		select {
		case OutA <- data:
		case OutB <- data:
		}
	}
	close(OutA)
	close(OutB)
}

func observeAndWrite(wg *sync.WaitGroup, ch <-chan int, channelName string) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(channelName, ": ", i)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	In := make(chan int)
	OutA := make(chan int)
	OutB := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			In <- i
		}
		close(In)
	}()
	wg.Add(3)
	go Fanout(wg, In, OutA, OutB)
	go observeAndWrite(wg, OutA, "A")
	go observeAndWrite(wg, OutB, "B")
	wg.Wait()
}
