package main

import (
	"cybercops.in/routers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// ------------------------------------------------------------------------
// Logger is a middleware handler that does request logging
// ------------------------------------------------------------------------
type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length, Authorization, Accept,X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	start := time.Now()
	fmt.Printf("%s %s %v\n", r.Method, r.URL.Path, time.Since(start))
	l.handler.ServeHTTP(w, r)
}

// ------------------------------------------------------------------------
// NewLogger constructs a new Logger middleware handler
// ------------------------------------------------------------------------
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

func main() {
	addr := ":3333"
	//init new router
	mux := mux.NewRouter()
	// request middleware for the request processing
	wapmux := NewLogger(mux)
	// add mux router the router path
	routers.Router(mux)

	fmt.Println("server is listening on port", addr)
	//servser listening on specified port as we configured
	log.Fatal(http.ListenAndServe(addr, wapmux))
}
