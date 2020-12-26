package main

import (
	"fmt"

	"github.com/pavelanni/gopl.io/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, мир!"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
