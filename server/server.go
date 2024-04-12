package server

import (
	"Netology/props"
	"Netology/pythagoreancalculation/handler"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

const (
	endpoint = "/pythogorean-calculation"
)

type Server struct {
	host        string
	port        int
	contextRoot string
	doOnce      sync.Once
}

func NewServer(properties *props.Properties) *Server {
	server := new(Server)
	server.host = properties.Server.Host
	server.port = properties.Server.Port
	server.contextRoot = properties.Server.ContextRoot
	return server
}

// ConfigureAPI configures the API with all the endpoints with respective handlers
func (s *Server) ConfigureAPI(calculate handler.Calculate) {
	s.doOnce.Do(func() {
		ConfigureApi(s.contextRoot, calculate, s.port)
	})
}

func ConfigureApi(contextRoot string, calculate handler.Calculate, port int) {
	var router = mux.NewRouter()
	router.HandleFunc(contextRoot+endpoint, func(rw http.ResponseWriter, r *http.Request) {
		HandlePythogoreanCalculation(rw, r, calculate)
	})

	log.Printf("\nApplication is running in : %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	if err != nil {
		log.Fatalf("Failure to start Go http server: %v", err)
		return
	}
}
