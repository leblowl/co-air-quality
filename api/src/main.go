package main

import (
	"co-air-quality.api/src/app"
)

func main() {
	var opts = map[string]string{"http-address": ":3030"}
	var a = app.App{}

	a.Init(opts)
	a.Start()
}
