package jsonconf

import (
	"log"
	"io/ioutil"
	"encoding/json"
)

func Parse(filename string, scheme interface{}) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	err = json.Unmarshal([]byte(body), scheme)
	if err != nil {
		log.Fatalf("Failed to decode config: %v", err)
	}
}
