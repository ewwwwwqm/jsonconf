// Package for parsing JSON files.
package jsonconf

import (
	"log"
	"io/ioutil"
	"encoding/json"
)

// Parse JSON using filename and scheme interface.
func Parse(filename string, scheme interface{}) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Failed to read config: %v", err)
	}
	err = json.Unmarshal([]byte(body), scheme)
	if err != nil {
		log.Println("Failed to decode config: %v", err)
	}
}
