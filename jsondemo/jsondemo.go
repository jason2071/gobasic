package jsonDemo

import (
	"encoding/json"
	"fmt"
	"gobasic/models"
	"log"
)

var data = `[
		{
			"userId": 1,
			"id": 1,
			"title": "delectus aut autem",
			"completed": false
		},
		{
			"userId": 1,
			"id": 2,
			"title": "quis ut nam facilis et officia qui",
			"completed": false
		},
		{
			"userId": 1,
			"id": 3
		}
	]`

func TodoList() {
	dataStruct := []models.Todo{}
	json.Unmarshal([]byte(data), &dataStruct)
	fmt.Printf("%+v\n", dataStruct)

	// dataStruct -> json
	json, err := json.MarshalIndent(dataStruct, "", "	")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(json))
}
