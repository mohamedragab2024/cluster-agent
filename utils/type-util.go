package utils

import (
	"encoding/json"
)

func StructToMap(obj interface{}) (newMap map[string]interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &newMap)
	if err != nil {
		return
	}
	return
}
