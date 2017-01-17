JSONConf
========

Package for parsing .json configuration files to defined structures.

Usage
-----
``` $ go get github.com/ewwwwwqm/jsonconf ```

Example
-------
``` config.json ```
```json
{
	"service_ids": [1, 34, 77703]
}
```

``` main.go ```
```go
package main

import (
	"fmt"
	"github.com/ewwwwwqm/jsonconf"
)

type Config struct {
	ServiceIds []int `json:"service_ids"`
}

var cfg Config

func main() {
	jsonconf.Parse("config.json", &cfg)
	fmt.Println(cfg)
}
```

Output:
  {[1 34 77703]}
