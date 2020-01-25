package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	morestrings "github.com/user/hello/more-strings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
