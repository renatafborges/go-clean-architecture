package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]HandlerInput
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]HandlerInput),
		WebServerPort: serverPort,
	}
}

type HandlerInput struct {
	method  string
	handler http.HandlerFunc
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Handlers[path] = HandlerInput{
		method:  method,
		handler: handler,
	}
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {

	s.Router.Use(middleware.Logger)
	for path, handlerInput := range s.Handlers {
		s.Router.Method(handlerInput.method, path, handlerInput.handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
