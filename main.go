package main

import (
	"fmt"

	"github.com/audaykumar/go-set-generics/set"
)

func main() {
	setInt := make(set.Set[int], 0)
	setInt.Add(1)

	setString := make(set.Set[string], 0)
	setString.Add("test1")
	setString.Add("test2")

	for k := range setString {
		fmt.Println(k)
	}
}
