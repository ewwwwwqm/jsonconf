// Package for parsing JSON files.
package jsonconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Parse JSON using filename and scheme interface.
func Parse(filename string, scheme interface{}) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Failed to read config:", err)
	}
	err = json.Unmarshal([]byte(body), scheme)
	if err != nil {
		fmt.Println("Failed to decode config:", err)
	}
}
