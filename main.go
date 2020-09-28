package main

import (
	"fmt"
)

func main() {
	source := `
	{
		"menu": {  
		"id": "file",  
		"value": false,  
		"popup": {  
			"menuitem": [  
				{"value": "New", "onclick": "CreateDoc()"},  
				{"value": "Open", "onclick": "OpenDoc()"},  
				{"value": "Save", "onclick": "SaveDoc()"}  
			],
			"test": {
				"menuitem": "ini string"
			}
		}  
	}}`

	target := `
	{  
		"employee": {  
			"name": "sonoo",   
			"salary": 56000,   
			"married": true  
		}  
	} 
	`

	newData := Mapping([]MapItem{
		{From: Item{"value", Boolean}, To: "name"},
		{From: Item{"menuitem", String}, To: "salary"},
	}, []byte(source), []byte(target))
	fmt.Println(string(newData))
}
