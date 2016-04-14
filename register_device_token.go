package iterable

import (
	"encoding/json"
)

const (
	REGISTER_DEVICE_TOKEN_URL = "users/registerDeviceToken"
)

type RegisterDeviceTokenRequest struct {
	Email  string                 `json:"email"`
	Device map[string]interface{} `json:"device"`
}

func (client *Client) RegisterDeviceToken(registerDeviceTokenRequest *RegisterDeviceTokenRequest) error {
	registerDeviceTokenRequestBytes, err := json.Marshal(registerDeviceTokenRequest)
	if err != nil {
		return err
	}

	_, err = client.iterableCall("POST", REGISTER_DEVICE_TOKEN_URL, registerDeviceTokenRequestBytes)
	if err != nil {
		return err
	}

	return nil
}
