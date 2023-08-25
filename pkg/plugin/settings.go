package plugin

import "encoding/json"

type jsonSettings struct {
	Host string `json:"host"`
}

func unmarshalSettings(m json.RawMessage) (*jsonSettings, error) {
	var settings jsonSettings
	err := json.Unmarshal(m, &settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}
