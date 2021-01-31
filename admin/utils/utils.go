package utils

import (
	"gopkg.in/yaml.v2"
)

func LoadComment(comment string) (result map[string]interface{}, err error) {
	result = make(map[string]interface{})
	err = yaml.Unmarshal([]byte(comment), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
