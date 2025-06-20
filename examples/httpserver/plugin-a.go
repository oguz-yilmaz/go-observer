package main

import (
	"fmt"
	observer "github.com/oguz-yilmaz/go-observer"
)

type PluginA struct{ prod *observer.EventProducer }

func NewPluginA(prod *observer.EventProducer) *PluginA {
	p := &PluginA{prod}

	observer.RegisterEvent(prod, EventOnServerStarting, p.ServerStarting)
	observer.RegisterEvent(prod, EventOnServerStarted, p.ServerStarted)
	observer.RegisterEvent(prod, EventOnServerStartError, p.ServerError)

	return p
}

func (PluginA) ServerStarting(d string) {
	fmt.Println("PLUGIN A handles ServerStarting with data: ", d)
}

func (PluginA) ServerStarted(d ServerStartedData) {
	fmt.Printf("PLUGIN A handles ServerStarted with message: %s on port %d\n", d.Message, d.Port)
	fmt.Printf("Only Plugin A handles ServerStarted event\n")
}

func (PluginA) ServerError(d ServerStartedErrorData) {
	fmt.Printf("PLUGIN A handles ServerError with message: %s, error: %v on port %d\n", d.Message, d.err, d.Port)
}
