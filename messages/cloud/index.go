package cloud

type CloudDevice interface {
	GetDeviceToken() string
}

type BasicCloudDevice struct {
	DeviceToken string `json:"device_token"`
}

func (b BasicCloudDevice) GetDeviceToken() string {
	return b.DeviceToken
}

type CloudDevices []CloudDevice

func (c CloudDevices) GetDeviceTokens() []string {
	result := make([]string, len(c))
	for i, d := range c {
		result[i] = d.GetDeviceToken()
	}
	return result
}

func NewBasicCloudDevice(dv string) BasicCloudDevice {
	return BasicCloudDevice{dv}
}

func NewBasicCloudDevices(dv []string) []BasicCloudDevice {
	result := make([]BasicCloudDevice, len(dv))
	for _, d := range dv {
		result = append(result, NewBasicCloudDevice(d))
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
	Error   string
}

type ICloudMessage interface {
	PushDevice(device CloudDevice, message CloudMessage) (PushResult, error)
	PushDevices(devices CloudDevices, message CloudMessage) ([]PushResult, error)
}
