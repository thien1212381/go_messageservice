package cloud

import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
)

type FirebasePusherConfig struct {
	KeyPush            string `json:"key_push"`
	TitlePush          string `json:"title_push"`
	IconPush           string `json:"icon_push"`
	SoundPush          string `json:"sound_push"`
	IsContentAvailable bool   `json:"is_content_available"`
	CollapseKey        string `json:"collapse_key"`
}

type FirebasePusher struct {
	fcmClient *fcm.FcmClient
	config    *FirebasePusherConfig
}

func (f *FirebasePusher) PushDevice(device CloudDevice, message CloudMessage) (PushResult, error) {
	result, err := f.PushDevices(CloudDevices{device}, message)
	return result[0], err
}

func (f *FirebasePusher) PushDevices(devices CloudDevices, message CloudMessage) ([]PushResult, error) {
	pushResult := make([]PushResult, len(devices))
	for i, device := range devices {
		pushResult[i].CloudDevice = device
	}

	notif := fcm.NotificationPayload{
		Title: f.config.TitlePush,
		Icon:  f.config.IconPush,
		Sound: f.config.SoundPush,
		Body:  message.Content,
		Badge: fmt.Sprintf("%d", message.Badge),
	}

	f.fcmClient.NewFcmRegIdsMsg(devices.GetDeviceTokens(), message.Payload)
	f.fcmClient.SetNotificationPayload(&notif)

	if f.config.CollapseKey != "" {
		f.fcmClient.SetCollapseKey(f.config.CollapseKey)
	}

	f.fcmClient.SetContentAvailable(f.config.IsContentAvailable)

	status, err := f.fcmClient.Send()
	if err != nil {
		return pushResult, err
	}

	for i, ps := range status.Results {
		for _, v := range ps {
			pushResult[i].Message = v
		}
	}

	return pushResult, nil
}

func NewFirebasePusher(config *FirebasePusherConfig) *FirebasePusher {
	fcmClient := fcm.NewFcmClient(config.KeyPush)

	if config.IconPush == "" {
		config.IconPush = "ic_launcher"
	}

	if config.SoundPush == "" {
		config.SoundPush = "default"
	}

	return &FirebasePusher{
		fcmClient: fcmClient,
		config:    config,
	}
}
