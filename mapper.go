package main

import (
	"encoding/json"
	"reflect"
)

type MapItem struct {
	From Item
	To   string
}

type Item struct {
	Field string
	Type  Type
}

type Type string

const (
	String  Type = "string"
	Number  Type = "float64"
	Boolean Type = "bool"
	Object  Type = "struct"
	Array   Type = "slice"
	Unknown Type = ""
)

func Mapping(mapItem []MapItem, source, target []byte) []byte {
	mpSource, _ := jsonToMap(source)
	mpTarget, _ := jsonToMap(target)
	for _, val := range mapItem {
		foundValue := find(mpSource, val.From.Field, val.From.Type)
		mpTarget = changeValue(mpTarget, val.To, foundValue)
	}
	return mapToJSON(mpTarget)
}

func mapToJSON(mapData map[string]interface{}) []byte {
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

func checkType(data interface{}) Type {
	defer recover()
	dataType := reflect.TypeOf(data).Kind().String()
	switch dataType {
	case "string":
		return String
	case "float64":
		return Number
	case "bool":
		return Boolean
	case "struct":
		return Object
	case "slice":
		return Array
	default:
		return Unknown
	}
}

func find(dataMap map[string]interface{}, keyName string, dataType Type) interface{} {
	for key, value := range dataMap {
		if key == keyName && checkType(value) == dataType {
			return value
		}
		switch value.(type) {
		case map[string]interface{}:
			return find(value.(map[string]interface{}), keyName, dataType)
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
