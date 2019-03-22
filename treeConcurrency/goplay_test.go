package main

import (
	"ConcurrencyTreeBrowser/tree"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		tr   tree.Tree
		want [10]int
	}{
		{tree.New(2, 10), []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}},
	}
	for _, c := range tests {
		got := displayTreeNoConcurrency(c.tr)
		if got != want {
			t.Errorf("displayTreeNoConcurrency")
		}
	}
}
