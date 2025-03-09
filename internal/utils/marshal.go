package utils

import "encoding/json"

// MarshalJson marshals a UserProfile into a JSON byte array.
func MarshalJson(input any) ([]byte, error) {
	inputBytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	return inputBytes, nil
}

// UnmarshalJson unmarshals a JSON byte array into a UserProfile.
func UnmarshalJson(data []byte, output any) error {
	err := json.Unmarshal(data, output)
	if err != nil {
		return err
	}

	return nil
}
