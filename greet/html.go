package greet

import (
	"encoding/json"
	"gobasic/models"
	"html/template"
	"net/http"
	"os"
)

func HtmlTemplate() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}

	todoDecoder := json.NewDecoder(resp.Body)
	todos := []models.Todo{}
	todoDecoder.Decode((&todos))
	resp.Body.Close()

	// todoEncoder := json.NewEncoder(os.Stdout)
	// todoEncoder.Encode(todos)

	indexTemplateFile, err := os.ReadFile("htmlTemplate/index.html")
	if err != nil {
		return
	}

	indexTemplate := template.Must(template.New("index").Parse(string(indexTemplateFile)))

	indexTemplate.Execute(os.Stdout, todos)
}
