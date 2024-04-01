package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/seanpden/govee_controller/pkg/structs"
)

// PrettyPrintJSON formats JSON data and prints it to the console.
//
// The function takes a parameter `body` of type `any`, which represents the JSON data to be formatted.
// It returns an error if there is an issue with the JSON formatting process.
func PrettyPrintJSON(body any) error {
	formatted_data, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(formatted_data))
	return nil
}

// SaveToJSON saves data to a JSON file.
//
// The function takes in a parameter 'data' of type 'any', which represents the data to be saved.
// It does not return any value.
func SaveToJSON(data any) {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("devices.json", file, 0666)
}

func LoadFromJSON(filepath string) (structs.ListDevicesResponse, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return structs.ListDevicesResponse{}, err
	}

	// var data map[string]interface{}
	var data structs.ListDevicesResponse
	err = json.Unmarshal(file, &data)
	if err != nil {
		return structs.ListDevicesResponse{}, err
	}

	return data, nil
}
