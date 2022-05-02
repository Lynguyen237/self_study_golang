package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}
