package test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	apiwrapper "github.com/seanpden/govee_controller/pkg/api_wrapper"
	"github.com/seanpden/govee_controller/pkg/utils"
)

// var APIKEY = handleEnvVar()

func handleEnvVar() string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	APIKEY := os.Getenv("APIKEY")
	return APIKEY
}

func TestListDevices(t *testing.T) {
	fmt.Println("TestListDevices")
	APIKEY := handleEnvVar()
	data, err := apiwrapper.ListDevices(APIKEY)
	if err != nil {
		log.Fatal(err)
	}
	// pretty print the list of devices
	utils.PrettyPrintJSON(data)

	fmt.Println()
}

func TestListDevicesWrongAPIKEY(t *testing.T) {
	fmt.Println("TestListDevicesWrongAPIKEY")
	APIKEY := "invalid"
	_, err := apiwrapper.ListDevices(APIKEY)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}

func TestGetDeviceState(t *testing.T) {
	fmt.Println("TestGetDeviceState")
	device := "F7:31:D1:39:38:38:5B:62"
	model := "H6072"
	APIKEY := handleEnvVar()
	data, err := apiwrapper.GetDeviceState(device, model, APIKEY)
	if err != nil {
		log.Fatal(err)
	}

	// pretty print the device stateA
	utils.PrettyPrintJSON(data)

	fmt.Println()
}

func TestSaveToJSON(t *testing.T) {
	fmt.Println("TestSaveToJSON")
	APIKEY := handleEnvVar()
	data, err := apiwrapper.ListDevices(APIKEY)
	if err != nil {
		log.Fatal(err)
	}
	utils.SaveToJSON(data)

	// pretty print the list of devices
	utils.PrettyPrintJSON(data)
	fmt.Println()
}

func TestLoadFromJSON(t *testing.T) {
	fmt.Println("TestLoadFromJSON")
	data, err := utils.LoadFromJSON("devices.json")
	if err != nil {
		log.Fatal(err)
	}

	// pretty print the list of devices
	utils.PrettyPrintJSON(data)
	fmt.Println()
}

func TestGetDeviceManyStates(t *testing.T) {
	fmt.Println("TestGetDeviceManyStates")
	devices := []string{"Lyra (Office: Right)"}
	APIKEY := handleEnvVar()
	data, err := apiwrapper.GetManyDeviceStates(devices, APIKEY)
	if err != nil {
		log.Fatal(err)
	}

	// pretty print the list of devices
	utils.PrettyPrintJSON(data)
	fmt.Println()
}

func TestTurnDeviceOn(t *testing.T) {
	fmt.Println("TestTurnDeviceOn")
	devices := []string{"F7:31:D1:39:38:38:5B:62"}
	APIKEY := handleEnvVar()
	data, err := apiwrapper.TurnDeviceOn(devices, APIKEY)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrintJSON(data)
	fmt.Println()
}

func TestTurnDeviceOff(t *testing.T) {
	fmt.Println("TestTurnDeviceOff")
	devices := []string{"F7:31:D1:39:38:38:5B:62"}
	APIKEY := handleEnvVar()
	data, err := apiwrapper.TurnDeviceOff(devices, APIKEY)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrintJSON(data)
	fmt.Println()
}

func TestSetDeviceBrightness(t *testing.T) {
	fmt.Println("TestSetDeviceBrightness")
	devices := []string{"F7:31:D1:39:38:38:5B:62"}
	APIKEY := handleEnvVar()
	data, err := apiwrapper.SetDeviceBrightness(devices, 10, APIKEY)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrintJSON(data)
	fmt.Println()

}

func TestSetDeviceRGB(t *testing.T) {
	fmt.Println("TestSetDeviceRGB")
	devices := []string{"Lyra (Office: Left)"}
	APIKEY := handleEnvVar()
	data, err := apiwrapper.SetDeviceRGB(devices, 255, 0, 0, APIKEY)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrintJSON(data)
	fmt.Println()

}

func TestSetDeviceColorTemp(t *testing.T) {
	fmt.Println("TestSetDeviceColorTemp")
	devices := []string{"F7:31:D1:39:38:38:5B:62"}
	APIKEY := handleEnvVar()
	data, err := apiwrapper.SetDeviceColorTemp(devices, 6500, APIKEY)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrintJSON(data)
	fmt.Println()

}
