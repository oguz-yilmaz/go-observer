package main

import (
	observer "github.com/oguz-yilmaz/go-observer"
)

type ServerStartedData struct {
	Message string
	Port    int
}

type ServerStartedErrorData struct {
	Message string
	err     error
	Port    int
}

var (
	EventOnServerStarting   = observer.Event[string]{Name: "on_server_starting"}
	EventOnServerStarted    = observer.Event[ServerStartedData]{Name: "on_server_started"}
	EventOnServerStartError = observer.Event[ServerStartedErrorData]{Name: "on_server_start_error"}
	EventOnHelloHandler     = observer.Event[string]{Name: "on_hello_handler"}
)
