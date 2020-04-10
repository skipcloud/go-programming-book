package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/skipcloud/go-programming-book/ch7/7.16/calculator"
)

/*
	Write a web-based calculator program
*/

const serverPort = "8000"
const host = "localhost"

type Server struct {
	port       string
	host       string
	calculator *calculator.Calculator
	output     *template.Template
}

func main() {
	c := calculator.New()
	server := Server{
		host:       host,
		port:       serverPort,
		calculator: c,
		output:     c.NewTemplate(),
	}
	http.HandleFunc("/", server.Output)
	http.HandleFunc("/input", server.Input)
	http.HandleFunc("/reset", server.Reset)
	server.Start()
}

func (s *Server) Input(w http.ResponseWriter, req *http.Request) {
	data := req.URL.Query().Get("data")
	if data != "" {
		s.calculator.Input(data)
	}
	s.output.Execute(w, nil)
}

func (s *Server) Output(w http.ResponseWriter, req *http.Request) {
	err := s.calculator.Calculate()
	if err != nil {
		w.WriteHeader(http.StatusTeapot)
		s.calculator.Reset()
	}
	s.output.Execute(w, nil)
}

func (s *Server) Reset(w http.ResponseWriter, req *http.Request) {
	s.calculator.Reset()
	s.output.Execute(w, nil)
}

func (s *Server) URL() string {
	return fmt.Sprintf("%s:%s", s.host, s.port)
}

func (s *Server) Start() {
	fmt.Printf("server starting on %s\n", s.URL())
	log.Fatal(http.ListenAndServe(":"+s.port, nil))
}
