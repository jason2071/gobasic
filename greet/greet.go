package greet

import (
	"fmt"
	"gobasic/hr"
	"html/template"
	"os"
	"strings"
)

func upperCase(s string) string {
	return strings.ToUpper(s)
}

func HiPersonTextTemplate() {
	john := hr.Person{Name: "John", Age: 30}
	elena := hr.Person{Name: "Elena", Age: 25}

	const greetPerson = `Hi I amd {{.Name}}. {{.Age}} years old{{"\n"}}`

	// fmt.Printf("Hi I amd %s. %d years old\n", john.Name, john.Age)
	// fmt.Printf("Hi I amd %s. %d years old\n", elena.Name, elena.Age)

	// greetTemplate, err := template.New("greet").Parse(greetPerson)
	// if err != nil {
	// 	return
	// }

	greetTemplate := template.Must(template.New("greet").Parse(greetPerson))

	fmt.Println(greetTemplate.Name())
	greetTemplate.Execute(os.Stdout, john)
	greetTemplate.Execute(os.Stdout, elena)
	fmt.Println("==========================")

	people := []hr.Person{john, elena}
	const greetPeople = `{{range .}}Hi I am {{.Name | upperCase }}. {{.Age}} years old {{"\n"}}{{end}}`

	// maps := template.FuncMap{"upperCase": upperCase}
	maps := make(template.FuncMap)
	maps["upperCase"] = upperCase

	greetPeopleTemplate := template.Must(
		template.New("greetPeople").
			Funcs(maps).
			Parse(greetPeople),
	)

	fmt.Println(greetPeopleTemplate.Name())
	greetPeopleTemplate.Execute(os.Stdout, people)
}
