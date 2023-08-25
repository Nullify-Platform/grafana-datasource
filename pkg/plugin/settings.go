package plugin

import (
	"encoding/json"
	"errors"
)

type jsonSettings struct {
	Host string `json:"host"`
}

func unmarshalSettings(m json.RawMessage) (*jsonSettings, error) {
	if m == nil {
		return nil, errors.New("settings are empty")
	}
	var settings jsonSettings
	err := json.Unmarshal(m, &settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}
