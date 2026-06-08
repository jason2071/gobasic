package jsonDemo

import (
	"encoding/json"
	"fmt"
	"gobasic/models"
	"io"
	"net/http"
)

type TodoJsonPlaceholder struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func TodoListJsonPlaceholder() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	dataStruct := []models.Todo{}
	json.Unmarshal(body, &dataStruct)
	fmt.Printf("len %d\n", len(dataStruct))

	// dataStruct -> json
	json, err := json.MarshalIndent(dataStruct, "", "	")
	if err != nil {
		return
	}

	fmt.Println(string(json))
}
