package main

import (
	"fmt"
)

// Assignment from go course
// https://coinbase.udemy.com/course/go-the-complete-developers-guide/learn/practice/9178/instructor-solution#overview
func main() {

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}

}
