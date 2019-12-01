package main

import (
	"fmt"
	"os"
)

func main1() {
	fmt.Println("a")
	var s, sep string
	for i:=1; i<len(os.Args);i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
