package Networking

import "net/http"

// HTTPServer creates and handles listener, later forwards data
type HTTPServer struct {
}

func (server *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
