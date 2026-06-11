package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}

func main() {
	go say("world")
	say("hello")
	say("wow")
}
