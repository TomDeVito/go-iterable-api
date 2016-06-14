package iterable

import (
	"encoding/json"
)

const (
	DELETE_USER_URL = "users/delete"
)

type DeleteUserRequest struct {
	Email string `json:"email"`
}

func (client *Client) DeleteUser(deleteUserRequest *DeleteUserRequest) error {
	deleteUserRequestBytes, err := json.Marshal(deleteUserRequest)
	if err != nil {
		return err
	}

	_, err = client.iterableCall("POST", DELETE_USER_URL, deleteUserRequestBytes)
	if err != nil {
		return err
	}

	return nil
}
