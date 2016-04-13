package iterable

import (
	"encoding/json"
)

type registerDeviceTokenRequest struct {
	Email  string
	Device map[string]interface{}
}

func (client *Client) RegisterDeviceToken(registerDeviceTokenRequest *registerDeviceTokenRequest) error {

	registerDeviceTokenRequestBytes, err := json.Marshal(registerDeviceTokenRequest)
	if err != nil {
		return err
	}

	_, err = client.iterableCall(REGISTER_DEVICE_TOKEN_URL, "POST", registerDeviceTokenRequestBytes)
	if err != nil {
		return err
	}

	return nil
}
