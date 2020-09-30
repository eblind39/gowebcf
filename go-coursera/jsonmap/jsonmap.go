// This sample program demonstrates how to encode map data into a JSON string.

package main

import (
	"encoding/json" // Encoding and Decoding Package
	"fmt"
)

func main() {
	// Create a map of key/value pairs and parses the data into JSON
	emp := make(map[string]interface{})
	emp["name"] = "Mark Taylor"
	emp["jobtitle"] = "Software Developer"
	/*emp["phone"] = map[string]interface{}{
		"home":   "123-466-799",
		"office": "564-987-654",
	}*/
	emp["email"] = "markt@gmail.com"

	// Marshal the map into a JSON string.
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonStr := string(empData)
	fmt.Println("The JSON data is:")
	fmt.Println(jsonStr)

}
