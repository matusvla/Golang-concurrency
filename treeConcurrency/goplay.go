package main

import (
	"ConcurrencyTreeBrowser/tree"
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
	ch <- t.Value
	wg.Done()
	wg.Wait()
	close(ch)
}

func printTreeNoConcurrency(t *tree.Tree) {
	if t == nil {
		return
	}
	fmt.Println(t.Value)
	printTreeNoConcurrency(t.Left)
	printTreeNoConcurrency(t.Right)
}

func main() {
	ch := make(chan int)
	t := tree.New(2, 10)
	wg.Add(1)
	go WalkInit(t, ch)
	for i := range ch {
		fmt.Println(i)
	}
	wg.Wait()
	// printTreeNoConcurrency(t)
}
