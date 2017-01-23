package main

import (
	"fmt"

	"github.com/proullon/roman"
)

func main() {
	for i := 1; i < 200; i++ {
		fmt.Printf("Chapter %3d: %s\n", i, roman.Value(i))
	}
}
