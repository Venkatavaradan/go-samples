package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	//"time"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	//fmt.Println(t.Value)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	//ch1val := 0
	//ch2val := 0
	for i := 0; i < 10; i++ {
		ch1val := <-ch1
		ch2val := <-ch2
		fmt.Println("ch1:", ch1val, "ch2:", ch2val)
		if ch1val != ch2val {
			return false
		}

	}
	return true
}

func main() {
	//ch := make (chan int );
	fmt.Println("Here..")
	//go Walk(tree.New(1), ch)
	//time.Sleep(2100* time.Millisecond)
	fmt.Println("Done..")
	flag := Same(tree.New(1), tree.New(2))
	fmt.Println(flag)

}
