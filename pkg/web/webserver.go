package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

// inspired by: https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831

type WebServer struct {
	router *http.ServeMux
	log    *logrus.Logger
}

func NewWebServer(log *logrus.Logger) *WebServer {
	s := &WebServer{}
	s.log = log
	s.router = http.NewServeMux()
	s.routes()
	return s
}

func (s *WebServer) Start() {
	go func() {
		err := http.ListenAndServe(":8080", s.router)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error starting web server: %s", err.Error()))
		}
	}()
}

func (s *WebServer) routes() {
	prefix := "/aresworld/"
	router := http.NewServeMux()
	router.HandleFunc("/ping", s.handlePing())
	router.HandleFunc("/whoami/", s.handleWhoAmI())
	router.HandleFunc("/gopykube", s.handleGoPyKube())
	s.AddHandle(prefix, router)
}

func (s *WebServer) handlePing() http.HandlerFunc {
	// private scope here because of the closure, which is nice :)
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong!"))
	}
}

func (s *WebServer) handleGoPyKube() http.HandlerFunc {
	// private scope here because of the closure, which is nice :)
	return func(w http.ResponseWriter, r *http.Request) {
		var result interface{}
		getURL := "http://192.168.64.2:30559/"

		req, err := http.NewRequest("GET", getURL, nil)
		if err != nil {
			fmt.Errorf("ERROR")
		}
		req.Header.Add("Accept", "application/json")
		resp, err := httpclient.client.Do(req)
		if err != nil {
			fmt.Errorf("ERROR")

		}

		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			fmt.Errorf("ERROR")

		}
		result, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}
}

func (s *WebServer) handleWhoAmI() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("I am Are"))
		return
	}
}

func (s *WebServer) AddHandle(prefix string, router *http.ServeMux) {
	s.router.Handle(prefix, http.StripPrefix(strings.TrimSuffix(prefix, "/"), router))
}
