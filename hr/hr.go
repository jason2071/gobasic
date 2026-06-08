package hr

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person      Person
	Designation string
}

func Profile() {

	employee := Employee{
		Person: Person{
			Name: "John Doe",
			Age:  30,
		},
		Designation: "Sale",
	}

	fmt.Printf("%+v\n", employee)

}
