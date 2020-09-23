package main

import (
	"encoding/json"
)

type MapItem struct {
	From string
	To   string
}

// type Item struct {
// 	Field      string
// 	Type       Type
// 	Validation string
// }

// type Type int

// const (
// 	String Type = iota
// 	Number
// 	Boolean
// 	Object
// 		Array
// )

func Mapping(mapItem []MapItem, source, target []byte) []byte {
	mpSource, _ := jsonToMap(source)
	mpTarget, _ := jsonToMap(target)
	for _, val := range mapItem {
		foundValue := find(mpSource, val.From)
		mpTarget = changeValue(mpTarget, val.To, foundValue)
	}
	return mapToJson(mpTarget)
}

// func typeValidation(value interface{}, dataType Type) bool {
// 	switch value.(type) {
// 	case string:
// 		if dataType == String {
// 			return true
// 		}
// 	case int32, float32:
// 		if dataType == Number {
// 			return true
// 		}
// 	case bool :
// 		if dataType == Boolean {
// 			return true
// 		}

// 	}
// }

func mapToJson(mapData map[string]interface{}) []byte {
	b, _ := json.MarshalIndent(mapData, "", "  ")
	return b
}

func jsonToMap(jsonData []byte) (map[string]interface{}, error) {
	var mp map[string]interface{}
	err := json.Unmarshal(jsonData, &mp)
	if err != nil {
		return nil, err
	}
	return mp, nil
}

func find(dataMap map[string]interface{}, keyName string) interface{} {
	for key, value := range dataMap {
		if key == keyName {
			return value
		}
		switch value.(type) {
		case map[string]interface{}:
			return find(value.(map[string]interface{}), keyName)
		}
	}
	return nil
}

func changeValue(dataMap map[string]interface{}, sourceKey string, newValue interface{}) map[string]interface{} {
	for key, value := range dataMap {
		if key == sourceKey {
			dataMap[key] = newValue
		}
		switch value.(type) {
		case map[string]interface{}:
			dataMap[key] = changeValue(value.(map[string]interface{}), sourceKey, newValue)
		}
	}
	return dataMap
}
