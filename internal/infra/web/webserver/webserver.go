package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	pathHandler, ok := s.Handlers[path]
	if !ok {
		s.Handlers[path] = make(map[string]http.HandlerFunc)
		pathHandler = s.Handlers[path]
	}
	pathHandler[method] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, pathHandler := range s.Handlers {
		for method, handler := range pathHandler {
			if method == "POST" {
				s.Router.Post(path, handler)
			} else {
				s.Router.Get(path, handler)
			}
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
