package main

import (
	"authoritymanage/initialize"
	"authoritymanage/router"
)

func init() {
	initialize.Init()
}

func main() {
	engine := router.GetEngine()
	if err := engine.Run(":8010"); err != nil {
		panic(err)
	}
}
