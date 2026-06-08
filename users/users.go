package users

import (
	"encoding/json"
	"fmt"
	"gobasic/models"
	"io"
	"net/http"
)

func User() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	dataStruct := []models.User{}
	json.Unmarshal(body, &dataStruct)
	fmt.Printf("len %d\n", len(dataStruct))

	// dataStruct -> json
	json, err := json.MarshalIndent(dataStruct, "", "	")
	if err != nil {
		return
	}

	fmt.Println(string(json))
}
