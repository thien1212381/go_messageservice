package cloud

import "testing"

func TestFirebasePusher_PushDevice(t *testing.T) {
	pusher := NewFirebasePusher(&FirebasePusherConfig{
		KeyPush:   "",
		TitlePush: "Test",
	})

	tokenTest := ""
	pushResult, err := pusher.PushDevice(NewBasicCloudDevice(tokenTest), CloudMessage{
		Badge:   1,
		Content: "Test FCM",
		Payload: map[string]interface{}{},
	})

	if err != nil {
		t.Error(err)
	} else {
		if pushResult.CloudDevice.GetDeviceToken() == tokenTest {
			if pushResult.Error == "" && pushResult.Message == "" {
				t.Error("No Response From FCM.")
			} else {
				t.Logf("%#v", pushResult)
			}
		} else {
			t.Errorf("NotExpected: Expected: %s - Result: %s", pushResult.CloudDevice.GetDeviceToken(), tokenTest)
		}
	}
}
