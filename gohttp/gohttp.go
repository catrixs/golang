package gohttp

import(
	"fmt"
	"os"
	"log"
	"time"
	"mime"
	"net/http"
	"path/filepath"
)

type HttpServer struct {
	htdocs  string
	port    int
	timeout int
}

func (s *HttpServer) handleResource(w http.ResponseWriter, r *http.Request) {
	filename := s.htdocs + r.URL.Path
	fmt.Printf("read file: %s\n", filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	contentType := mime.TypeByExtension(filepath.Ext(filename))
	w.Header().Set("Content-Type", contentType)
	http.ServeFile(w, r, filename)
}

func (s *HttpServer) Start() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", s.handleResource)
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", s.port),
		Handler: multiplexer,
		ReadTimeout:    time.Duration(s.timeout) * time.Second,
		WriteTimeout:   time.Duration(s.timeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}

func Server(htdocs string, port int, timeout int) *HttpServer {
	s := &HttpServer {
		htdocs : htdocs,
		port   : port,
		timeout: timeout,
	}
	return s
}
