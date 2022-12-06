package main

import "fmt"

func main() {
	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
		fmt.Println("inside for loop - i is: ", i)
		fmt.Println("inside for loop - &i is: ", &i)
		fmt.Println("inside for loop - out array is: ", out)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
	/*
		Values: 3 3 3
		Addresses: 0x40e020 0x40e020 0x40e020
		https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
	*/
}
