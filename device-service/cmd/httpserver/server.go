package httpserver

import (
	"device-service/cmd/handler"
	"fmt"
	"log"
	"net/http"
)

type httpServer struct {
	port        string
	httpHandler http.Handler
}

func NewHttpServer(port string) *httpServer {

	h := handler.HttpHandler{}

	return &httpServer{
		port:        port,
		httpHandler: h.Routes(),
	}
}

func (s *httpServer) Run() error {

	err := http.ListenAndServe(
		fmt.Sprintf(":%s", s.port),
		s.httpHandler)

	if err != nil {
		log.Fatal(err)
	}

	return nil

}
