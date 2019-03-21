package main

import (
	"ConcurencyTreeBrowser/tree"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer wg.Done()
	if t.Left != nil {
		wg.Add(1)
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		wg.Add(1)
		go Walk(t.Right, ch)
	}
	ch <- t.Value
}

func WalkInit(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		wg.Add(1)
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		wg.Add(1)
		go Walk(t.Right, ch)
	}
	wg.Done()
	wg.Wait()
	close(ch)
}

func main() {
	ch := make(chan int)
	t := tree.New(2)
	wg.Add(1)
	go WalkInit(t, ch)
	for i := range ch {
		fmt.Println(i)
	}
	wg.Wait()
}
