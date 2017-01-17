# JSONConf

[![Godoc Reference][godoc-img]][godoc]
[![Build Status][ci-img]][ci]
[![Go Report Card][goreportcard-img]][goreportcard]

Package for parsing .json configuration files to defined structures.

## Usage
``` $ go get github.com/ewwwwwqm/jsonconf ```

## Example
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
```go
{[1 34 77703]}
```

[godoc]: http://godoc.org/github.com/ewwwwwqm/jsonconf
[godoc-img]: https://godoc.org/github.com/ewwwwwqm/jsonconf?status.svg
[ci-img]: https://travis-ci.org/ewwwwwqm/jsonconf.svg?branch=master
[cov-img]: https://coveralls.io/repos/github/ewwwwwqm/jsonconfbadge.svg?branch=master
[ci]: https://travis-ci.org/ewwwwwqm/jsonconf
[cov]: https://coveralls.io/github/ewwwwwqm/jsonconf?branch=master
[goreportcard-img]: https://goreportcard.com/badge/github.com/ewwwwwqm/jsonconf
[goreportcard]: https://goreportcard.com/report/github.com/ewwwwwqm/jsonconf
