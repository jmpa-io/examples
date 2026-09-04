package main

import "net/http"

type Server struct {

}

func NewServer() *Server {
  return &Server{

  }
}

// app is the global server instance used as the HTTP handler.
var app = NewServer()

func (s *Server) handleRootEndpoint(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Welcome to the root endpoint!"))
}

func (s *Server) handleHealthEndpoint(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("OK"))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  switch r.URL.Path {
  case "/":
    s.handleRootEndpoint(w, r)
    return

  case "/health":
    s.handleHealthEndpoint(w, r)
    return
  }
}
