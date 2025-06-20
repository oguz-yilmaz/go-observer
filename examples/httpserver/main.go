package main

import (
	"fmt"
	observer "github.com/oguz-yilmaz/go-observer"
	"net/http"
)

type App struct {
	*observer.EventProducer
}

func NewApp() *App {
	return &App{
		EventProducer: observer.NewEventProducer(),
	}
}

func (a *App) helloHandler(w http.ResponseWriter, r *http.Request) {
	observer.FireEvent(a.EventProducer, EventOnHelloHandler, "Hello handler called")

	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	app := NewApp()

	NewPluginA(app.EventProducer)
	NewPluginB(app.EventProducer)

	http.HandleFunc("/", app.helloHandler)

	observer.FireEvent(app.EventProducer, EventOnServerStarting, "[ServerInfo:] Starting server at http://localhost:8080")

	observer.FireEvent(app.EventProducer, EventOnServerStarted, ServerStartedData{
		Message: "[ServerMessage:] Server Started successfully",
		Port:    8080,
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		observer.FireEvent(app.EventProducer, EventOnServerStartError, ServerStartedErrorData{
			Message: "[ServerError:] Server failed to start",
			err:     err,
			Port:    8080,
		})

		fmt.Println("Server failed:", err)
	}
}
