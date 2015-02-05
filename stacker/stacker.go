package main

import (
	"fmt"
	"github.com/alexgula/go-playground/stacker/stack"
)

func main() {
	var haystack stack.Stack
	haystack.Push("hay").Push(-15).Push([]string{"pin", "clip", "needle"}).Push(81.5)
	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
