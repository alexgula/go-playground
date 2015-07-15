package main

import (
	"fmt"
)

func main() {
	var d int
	for {
		n, err := fmt.Scanf("%d\n", &d)
		if n == 0 {
			fmt.Println("Read 0 bytes, strange")
			continue
		}
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if d == 42 {
			return
		}
		fmt.Println(d)
	}
}
