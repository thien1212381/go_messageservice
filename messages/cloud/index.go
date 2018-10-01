package cloud

type CloudDevice interface {
	GetDeviceToken() string
}

type CloudDevices []CloudDevice

func (c CloudDevices) GetDeviceTokens() []string {
	result := make([]string, len(c))
	for i, d := range c {
		result[i] = d.GetDeviceToken()
	}
	return result
}

type CloudMessage struct {
	Badge   int                    `json:"badge"`
	Content string                 `json:"content"`
	Payload map[string]interface{} `json:"payload"`
}

type PushResult struct {
	CloudDevice
	Message string
}

type ICloudMessage interface {
	PushDevice(device CloudDevice, message CloudMessage) (PushResult, error)
	PushDevices(devices CloudDevices, message CloudMessage) ([]PushResult, error)
}
