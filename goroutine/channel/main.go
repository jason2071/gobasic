package main

import "fmt"

// sum หาผลรวมของทุกค่าใน slice s แล้วส่งผลลัพธ์เข้า channel c
// แทนการ return ค่า เพื่อให้ฝั่งเรียกรับผลลัพธ์จาก goroutine
// ที่ทำงานพร้อมกัน (concurrent) ได้
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // ส่ง sum เข้า channel c (block จนกว่าจะมีคนรับ)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// สร้าง unbuffered channel ชนิด int
	// การ send และ receive จะ block จนกว่าอีกฝั่งจะพร้อม
	// channel จึงทำหน้าที่เป็นจุด synchronize ไปในตัว
	c := make(chan int)

	// แบ่ง slice ครึ่งหนึ่ง แล้วหาผลรวมแต่ละครึ่งใน goroutine แยกกัน
	// ทั้งสอง goroutine ทำงานพร้อมกัน และส่งผลลัพธ์เข้า c
	go sum(s[:len(s)/2], c)  // ครึ่งแรก: [7, 2, 8]
	go sum(s[len(s)/2:], c)  // ครึ่งหลัง: [-9, 4, 0]

	// รับค่า 2 ค่าจาก c — ค่าละหนึ่งจากแต่ละ goroutine
	// แต่ละการ receive จะ block จนกว่า goroutine จะส่งผลลัพธ์มา
	// ลำดับไม่การันตี: x อาจเป็นผลรวมของครึ่งไหนก็ได้ ขึ้นกับว่าใครเสร็จก่อน
	x, y := <-c, <-c // รับค่าจาก c

	fmt.Println(x, y, x+y) // พิมพ์ผลรวมย่อยทั้งสอง และผลรวมทั้งหมด (12)
}
