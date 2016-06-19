package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"Name":"Danqi", "Age":20, "Friends":["Tobin", "Chris"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Printf("Error parsing json file : %v", err)
		return
	}
	m := f.(map[string]interface{})
	
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an arrray:")
			for i, u := range vv {
				fmt.Println(i,u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}	
	}
}
