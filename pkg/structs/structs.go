package structs

type ListDevicesResponse struct {
	Data struct {
		Devices []struct {
			Device       string   `json:"device"`
			Model        string   `json:"model"`
			DeviceName   string   `json:"deviceName"`
			Controllable bool     `json:"controllable"`
			Retrievable  bool     `json:"retrievable"`
			SupportCmds  []string `json:"supportCmds"`
			Properties   struct {
				ColorTem struct {
					Range struct {
						Min int `json:"min"`
						Max int `json:"max"`
					} `json:"range"`
				} `json:"colorTem"`
			} `json:"properties"`
		} `json:"devices"`
	} `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Payload struct {
	Device string  `json:"device"`
	Model  string  `json:"model"`
	Cmd    Command `json:"cmd"`
}

type Command struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type _Command struct {
	Value string
}

type DeviceStateResponse struct {
	Data    Data   `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}
type Color struct {
	R int `json:"r,omitempty"`
	G int `json:"g,omitempty"`
	B int `json:"b,omitempty"`
}
type Properties struct {
	Online           bool   `json:"online,omitempty"`
	PowerState       string `json:"powerState,omitempty"`
	Brightness       int    `json:"brightness,omitempty"`
	ColorTemInKelvin *int   `json:"colorTemInKelvin,omitempty"`
	ColorTem         *int   `json:"colorTem,omitempty"`
	Color            *Color `json:"color,omitempty"`
}
type Data struct {
	Device     string       `json:"device,omitempty"`
	Model      string       `json:"model,omitempty"`
	Properties []Properties `json:"properties,omitempty"`
}

type ControlDeviceResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}
