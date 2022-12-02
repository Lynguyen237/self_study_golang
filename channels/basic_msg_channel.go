package main

import "fmt"

func main() {
	messages := make(chan string) // channels are typed by the value they convey

	go func() {
		messages <- "first message" // SEND a value into a channel using the channel <- synatax
	}()

	msg := <-messages // <- channel syntax RECEIVES a value from the channel
	fmt.Println(msg)
}

/*
By default sends and receives block until both the sender and receiver are ready.
This property allowed us to wait at the end of our program for the "ping" message
without having to use any other synchronization.
https://gobyexample.com/channels
*/
