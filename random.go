package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number from [0:5)
	fmt.Println(rand.Intn(5))

	//Generate a random number x where x is in range 5<=x<=20
	rangeLower := 5
	rangeUpper := 20
	randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
	fmt.Println(randomNum)

}
