package jsonDemo

import (
	"encoding/json"
	"gobasic/models"
	"net/http"
	"os"
)

func JsonEncoderDecoder() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return
	}

	jsonDecoder := json.NewDecoder(resp.Body)

	dataStruct := []models.User{}
	dataStruct2 := []models.User{}

	jsonDecoder.Decode(&dataStruct)
	jsonDecoder.Decode(&dataStruct2)
	resp.Body.Close()

	jsonEncoder := json.NewEncoder(os.Stdout)
	jsonEncoder.Encode(dataStruct)

	// fmt.Println(jsonEncoder)
}
