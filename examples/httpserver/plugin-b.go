package main

import (
	"fmt"
	observer "github.com/oguz-yilmaz/go-observer"
)

type PluginB struct{ prod *observer.EventProducer }

func NewPluginB(prod *observer.EventProducer) *PluginB {
	p := &PluginB{prod}

	observer.RegisterEvent(prod, EventOnServerStarting, p.ServerStarting)

	return p
}

func (PluginB) ServerStarting(d string) {
	fmt.Println("PLUGIN B handles ServerStarting with data: ", d)
}
