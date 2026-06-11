package main

import "fmt"

// greet ส่งข้อความทักทายเข้า channel c
func greet(name string, c chan string) {
	c <- "hello " + name
}

func main() {
	// สร้าง channel ชนิด string
	c := make(chan string)

	// รัน greet เป็น goroutine ส่งข้อความเข้า channel
	go greet("john", c)

	// main รอรับค่าจาก channel (block จนกว่าจะมีค่าส่งมา)
	msg := <-c

	fmt.Println(msg) // hello
}
