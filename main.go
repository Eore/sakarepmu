package main

import (
	"fmt"
)

func main() {
	source := `
	{
		"menu": {  
		"id": "file",  
		"value": "File",  
		"popup": {  
			"menuitem": [  
				{"value": "New", "onclick": "CreateDoc()"},  
				{"value": "Open", "onclick": "OpenDoc()"},  
				{"value": "Save", "onclick": "SaveDoc()"}  
			]  
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
		{From: "value", To: "name"},
		{From: "menuitem", To: "salary"},
	}, []byte(source), []byte(target))
	fmt.Println(string(newData))
}
