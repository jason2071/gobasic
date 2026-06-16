package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func topStudent(students []Student) Student {

	if len(students) == 0 {
		return Student{}
	}

	top := students[0]
	for _, s := range students {
		if s.Score > top.Score {
			top = s
		}
	}
	return top
}

func main() {
	students := []Student{
		{"Ann", 80},
		{"Bob", 95},
		{"Cat", 88},
	}
	r := topStudent(students) // ได้ {Bob 95}
	fmt.Println(r)
}
