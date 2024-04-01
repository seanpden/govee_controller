package clihandler

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	apiwrapper "github.com/seanpden/govee_controller/pkg/api_wrapper"
)

type deviceSliceFlag []string

func (d *deviceSliceFlag) String() string {
	return fmt.Sprintf("%v", *d)
}

func (d *deviceSliceFlag) Set(value string) error {
	inputs := strings.Split(value, ",")
	*d = append(*d, inputs...)
	return nil
}

func handleTurnDeviceOnOff(device deviceSliceFlag, value string, APIKEY string) {
	if value == "on" {
		fmt.Println("Turning device on")
		data, err := apiwrapper.TurnDeviceOn(device, APIKEY)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data)
		return
	}
	if value == "off" {
		fmt.Println("Turning device off")
		data, err := apiwrapper.TurnDeviceOff(device, APIKEY)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data)
		return
	}
	fmt.Println("Invalid value")
	return
}

func handleListDevices(APIKEY string) {
	fmt.Println("Getting list of devices")
	data, err := apiwrapper.ListDevices(APIKEY)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	return
}

func handleGetDeviceState(device deviceSliceFlag, APIKEY string) {
	fmt.Println("Getting device state")
	data, err := apiwrapper.GetManyDeviceStates(device, APIKEY)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	return
}

func handleSetBrightness(device deviceSliceFlag, value string, APIKEY string) {
	fmt.Println("Setting device brightness")
	brightnessLevel, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(err)
	}
	data, err := apiwrapper.SetDeviceBrightness(device, brightnessLevel, APIKEY)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	return
}

func handleSetColor(device deviceSliceFlag, value string, APIKEY string) {
	fmt.Println("Setting device color")
	color := strings.Split(value, ",")
	r, err := strconv.Atoi(color[0])
	if err != nil {
		fmt.Println(err)
	}
	g, err := strconv.Atoi(color[1])
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.Atoi(color[2])
	if err != nil {
		fmt.Println(err)
	}
	data, err := apiwrapper.SetDeviceRGB(device, r, g, b, APIKEY)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	return
}

func handleColorTemp(device deviceSliceFlag, value string, APIKEY string) {
	fmt.Println("Setting device color temp")
	colorTemp, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(err)
	}
	data, err := apiwrapper.SetDeviceColorTemp(device, colorTemp, APIKEY)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	return
}

func HandleCLI(APIKEY string) {
	var Device deviceSliceFlag
	Cmd := flag.String("CMD", "", "what command to execute")
	Value := flag.String("VALUE", "", "what the command value is. e.g. 'on', 'off', '255,255,255', etc.")
	flag.Var(&Device, "DEVICE", "what device(s) to execute command on")
	flag.Parse()

	// if "all" is in the device slice, get all device names and set it to the device slice
	if len(Device) == 1 && Device[0] == "all" {
		data, err := apiwrapper.ListDevices(APIKEY)
		if err != nil {
			fmt.Println(err)
		}
		for _, device := range data.Data.Devices {
			Device = append(Device, device.DeviceName)
		}
	}

	if *Cmd == "turn" {
		handleTurnDeviceOnOff(Device, *Value, APIKEY)
		return
	}

	if *Cmd == "list" {
		handleListDevices(APIKEY)
		return
	}

	if *Cmd == "get" {
		// TODO: Fix response
		handleGetDeviceState(Device, APIKEY)
		return
	}

	if *Cmd == "brightness" {
		handleSetBrightness(Device, *Value, APIKEY)
		return
	}

	if *Cmd == "color" {
		handleSetColor(Device, *Value, APIKEY)
		return
	}

	if *Cmd == "color_temp" {
		handleColorTemp(Device, *Value, APIKEY)
		return
	}

	fmt.Println("Invalid inputs. Please try again or use the -h flag for help")
}
