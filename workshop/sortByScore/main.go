package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Score int
}

func sortScore(students []Student) []Student {
	sort.Slice(students, func(i, j int) bool {
		if students[i].Score > students[j].Score {
			return true
		}
		return false
	})

	// sort.Slice(students, func(i, j int) bool {
	// 	return students[i].Score > students[j].Score
	// })

	return students
}

func main() {
	students := []Student{{"Ann", 80}, {"Bob", 95}, {"Cat", 88}}
	// r := sortScore(students)
	// fmt.Println(r)

	sortScore(students)
	fmt.Println(students)

}
