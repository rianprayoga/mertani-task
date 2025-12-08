package httpserver

import (
	"device-service/cmd/handler"
	"device-service/internal/repository"
	"fmt"
	"log"
	"net/http"
)

type httpServer struct {
	port        string
	httpHandler handler.HttpHandler
}

func NewHttpServer(port string, db repository.Repo) *httpServer {

	h := handler.HttpHandler{
		Db: db,
	}

	return &httpServer{
		port:        port,
		httpHandler: h,
	}
}

func (s *httpServer) Run() error {

	err := http.ListenAndServe(
		fmt.Sprintf(":%s", s.port),
		s.httpHandler.Routes())

	if err != nil {
		log.Fatal(err)
	}

	return nil

}
