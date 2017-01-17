package jsonconf_test

import (
	"fmt"
	"io/ioutil"
	"os"

    "github.com/ewwwwwqm/jsonconf"
)

func ExampleParse() {
	// temporary filename
	fname := "./config.json.test"

	// set JSON inside config
	json := []byte(`{"name": "Max"}`)
	err := ioutil.WriteFile(fname, json, 0644)
    if err != nil {
        panic(err)
    }

    // define Config structure
	type Config struct {
		Name string `json:"name"`
	}
	
	// assign local variable
	var cfg Config
	
	// parse JSON file
	jsonconf.Parse(fname, &cfg)
	fmt.Println(cfg.Name)
	// Output:
	// Max

	err = os.Remove(fname)
	if err != nil {
		panic(err)
	}
}
