package utils

import (
	"encoding/json"
	"io"
)

func StructToMap(this interface{}) (newMap map[string]interface{}) {
	data, err := json.Marshal(this)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &newMap)
	if err != nil {
		return
	}
	return
}

func MapToStruct(this map[string]interface{}) (Newobj interface{}) {
	data, err := json.Marshal(this)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &Newobj)
	if err != nil {
		return
	}
	return
}

func MapToJson(this map[string]interface{}) (jsonObj []byte) {
	jsonObj, err := json.Marshal(this)
	if err != nil {
		return
	}
	return
}

func JsonBodyToMap(this io.ReadCloser) (Newobj map[string]interface{}) {
	err := json.NewDecoder(this).Decode(&Newobj)
	if err != nil {
		return
	}
	return
}
