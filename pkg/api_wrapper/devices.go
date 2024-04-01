package apiwrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seanpden/govee_controller/pkg/structs"
	"github.com/seanpden/govee_controller/pkg/utils"
	_ "github.com/seanpden/govee_controller/pkg/utils"
)

// createHeader generates the headers for an HTTP request.
//
// It takes in a pointer to an http.Request object and an API key string.
// The function modifies the request object by adding three header fields:
// "accept" with the value "application/json", "content-type" with the value
// "application/json", and "Govee-API-Key" with the value of the API key.
//
// Parameters:
// - req: A pointer to an http.Request object.
// - APIKEY: The API key to include in the request header.
func createHeader(req *http.Request, APIKEY string) {
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Govee-API-Key", APIKEY)
}

func constructPayload(device string, model string, cmd structs.Command) (io.Reader, error) {
	// TODO
	payload := structs.Payload{
		Device: device,
		Model:  model,
		Cmd:    cmd,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

// makeRequest sends an HTTP request to the specified URL using the given method and API key.
//
// Parameters:
// - method: The HTTP method to use for the request.
// - url: The URL to send the request to.
// - APIKEY: The API key to include in the request header.
//
// Returns:
// - []byte: The response body as a byte array.
// - error: An error if any occurred during the request or response handling.
func makeRequest(method string, url string, payload io.Reader, APIKEY string) ([]byte, error) {
	// create the request, err handling
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	createHeader(req, APIKEY)

	// make the request, err handling
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// check if response is a 200
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// read body and return response as []byte
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// ListDevices retrieves a list of devices using the provided API key.
//
// Parameters:
//
// - APIKEY: The API key used to authenticate the request.
//
// Returns:
//
// - structs.ListDevicesResponse: The response containing the list of devices.
// - error: An error if the API request fails.
func ListDevices(APIKEY string) (structs.ListDevicesResponse, error) {
	// instantiate vars needed for api request
	url := "https://developer-api.govee.com/v1/devices"
	method := "GET"

	// make request, error handling
	body, err := makeRequest(method, url, nil, APIKEY)
	if err != nil {
		return structs.ListDevicesResponse{}, err
	}

	// convert response body to go struct
	var structuredBody structs.ListDevicesResponse
	err = json.Unmarshal(body, &structuredBody)
	if err != nil {
		return structs.ListDevicesResponse{}, err
	}

	return structuredBody, nil
}

// GetDeviceState retrieves the state of a device from the Govee API.
//
// Parameters:
//
//	device (string): The ID of the device.
//	model (string): The model of the device.
//	APIKEY (string): The API key for authentication.
//
// Returns:
//
// - DeviceStateResponse: The response containing the device state.
// - error: An error if the API request fails.
func GetDeviceState(device string, model string, APIKEY string) (structs.DeviceStateResponse, error) {
	// instantiate vars needed for api request
	url := fmt.Sprintf("https://developer-api.govee.com/v1/devices/state?device=%s&model=%s", device, model)
	method := "GET"

	// make request, error handling
	body, err := makeRequest(method, url, nil, APIKEY)
	if err != nil {
		return structs.DeviceStateResponse{}, err
	}

	// convert response body to go struct
	var structuredBody structs.DeviceStateResponse
	err = json.Unmarshal(body, &structuredBody)
	if err != nil {
		return structs.DeviceStateResponse{}, err
	}

	return structuredBody, nil
}

// GetManyDeviceStates loads a list of devices from a json file and retrieves the state of each supplied device.
//
// Parameters:
//
// - devices: A slice of strings representing the devices for which to retrieve the state.
// - APIKEY: A string representing the API key to be used for authentication.
//
// Returns:
//
// - []structs.DeviceStateResponse: A slice of structs representing the device states.
// - error: An error object if there was a problem loading the devices or retrieving their states.
func GetManyDeviceStates(devices []string, APIKEY string) ([]structs.DeviceStateResponse, error) {
	// load a list of devices from a json file
	devicesJSON, err := utils.LoadFromJSON("devices.json")
	if err != nil {
		return []structs.DeviceStateResponse{}, err
	}

	var deviceStates []structs.DeviceStateResponse

	// iterate over list of supplied devices, if the device is in the list
	// get its state and append it to devicesStates
	for _, device := range devices {
		for _, deviceJSON := range devicesJSON.Data.Devices {
			if device == deviceJSON.DeviceName {
				deviceState, err := GetDeviceState(deviceJSON.Device, deviceJSON.Model, APIKEY)

				if err != nil {
					return []structs.DeviceStateResponse{}, err
				}
				deviceStates = append(deviceStates, deviceState)
			}
		}
	}

	return deviceStates, nil

}

// TurnDeviceOn turns on a list of devices using the provided API key.
//
// The function takes in two parameters:
// - devices: a slice of strings representing the devices to be turned on.
// - apiKey: a string representing the API key required for authentication.
//
// The function returns a structs.ControlDeviceResponse and an error.
func TurnDeviceOn(devices []string, apiKey string) (structs.ControlDeviceResponse, error) {
	devicesJSON, err := utils.LoadFromJSON("devices.json")
	if err != nil {
		return structs.ControlDeviceResponse{}, err
	}

	url := "https://developer-api.govee.com/v1/devices/control"
	cmd := structs.Command{
		Name:  "turn",
		Value: "on",
	}

	var response structs.ControlDeviceResponse

	for _, device := range devices {
		for _, deviceJSON := range devicesJSON.Data.Devices {
			if device == deviceJSON.DeviceName {
				payload, err := constructPayload(deviceJSON.Device, deviceJSON.Model, cmd)
				if err != nil {
					return response, err
				}
				body, err := makeRequest("PUT", url, payload, apiKey)
				if err != nil {
					return response, err
				}

				err = json.Unmarshal(body, &response)
				if err != nil {
					return response, err
				}
			}
		}
	}

	return response, nil
}

func TurnDeviceOff(devices []string, apiKey string) (structs.ControlDeviceResponse, error) {
	devicesJSON, err := utils.LoadFromJSON("devices.json")
	if err != nil {
		return structs.ControlDeviceResponse{}, err
	}

	url := "https://developer-api.govee.com/v1/devices/control"
	cmd := structs.Command{
		Name:  "turn",
		Value: "off",
	}

	var response structs.ControlDeviceResponse

	for _, device := range devices {
		for _, deviceJSON := range devicesJSON.Data.Devices {
			if device == deviceJSON.DeviceName {
				payload, err := constructPayload(deviceJSON.Device, deviceJSON.Model, cmd)
				if err != nil {
					return response, err
				}
				body, err := makeRequest("PUT", url, payload, apiKey)
				if err != nil {
					return response, err
				}

				err = json.Unmarshal(body, &response)
				if err != nil {
					return response, err
				}
			}
		}
	}

	return response, nil

}

func SetDeviceBrightness(devices []string, brightness int, apiKey string) (structs.ControlDeviceResponse, error) {
	devicesJSON, err := utils.LoadFromJSON("devices.json")
	if err != nil {
		return structs.ControlDeviceResponse{}, err
	}

	// check if brightness is between 0-100
	if brightness < 0 || brightness > 100 {
		return structs.ControlDeviceResponse{}, fmt.Errorf("brightness must be between 0-100")
	}

	url := "https://developer-api.govee.com/v1/devices/control"
	cmd := structs.Command{
		Name:  "brightness",
		Value: brightness,
	}

	var response structs.ControlDeviceResponse

	for _, device := range devices {
		for _, deviceJSON := range devicesJSON.Data.Devices {
			if device == deviceJSON.DeviceName {
				payload, err := constructPayload(deviceJSON.Device, deviceJSON.Model, cmd)
				if err != nil {
					return response, err
				}
				body, err := makeRequest("PUT", url, payload, apiKey)
				if err != nil {
					return response, err
				}

				err = json.Unmarshal(body, &response)
				if err != nil {
					return response, err
				}
			}
		}
	}

	return response, nil

}

func SetDeviceRGB(devices []string, r int, g int, b int, apiKey string) (structs.ControlDeviceResponse, error) {
	devicesJSON, err := utils.LoadFromJSON("devices.json")
	if err != nil {
		return structs.ControlDeviceResponse{}, err
	}

	// check if values are between 0 and 255
	if r > 255 || r < 0 || g > 255 || g < 0 || b > 255 || b < 0 {
		return structs.ControlDeviceResponse{}, fmt.Errorf("r, g, and b must be between 0 and 255")
	}

	url := "https://developer-api.govee.com/v1/devices/control"
	cmd := structs.Command{
		Name: "color",
		Value: struct {
			Name string `json:"name"`
			R    int    `json:"r"`
			G    int    `json:"g"`
			B    int    `json:"b"`
		}{
			Name: "Color",
			R:    r,
			G:    g,
			B:    b,
		},
	}

	var response structs.ControlDeviceResponse

	for _, device := range devices {
		for _, deviceJSON := range devicesJSON.Data.Devices {
			if device == deviceJSON.DeviceName {
				payload, err := constructPayload(deviceJSON.Device, deviceJSON.Model, cmd)
				if err != nil {
					return response, err
				}
				body, err := makeRequest("PUT", url, payload, apiKey)
				if err != nil {
					return response, err
				}

				err = json.Unmarshal(body, &response)
				if err != nil {
					return response, err
				}
			}
		}
	}

	return response, nil

}

func SetDeviceColorTemp(devices []string, colorTemp int, apiKey string) (structs.ControlDeviceResponse, error) {
	devicesJSON, err := utils.LoadFromJSON("devices.json")
	if err != nil {
		return structs.ControlDeviceResponse{}, err
	}

	// check if colorTemp is between 2000-9000
	if colorTemp < 2000 || colorTemp > 9000 {
		return structs.ControlDeviceResponse{}, fmt.Errorf("colorTemp must be between 2000-9000")
	}

	url := "https://developer-api.govee.com/v1/devices/control"
	cmd := structs.Command{
		Name:  "colorTem",
		Value: colorTemp,
	}

	var response structs.ControlDeviceResponse

	for _, device := range devices {
		for _, deviceJSON := range devicesJSON.Data.Devices {
			if device == deviceJSON.DeviceName {
				payload, err := constructPayload(deviceJSON.Device, deviceJSON.Model, cmd)
				if err != nil {
					return response, err
				}
				body, err := makeRequest("PUT", url, payload, apiKey)
				if err != nil {
					return response, err
				}

				err = json.Unmarshal(body, &response)
				if err != nil {
					return response, err
				}
			}
		}
	}

	return response, nil

}
