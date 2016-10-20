package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jraedisch/go_sse_example/events"
)

func main() {
	http.Handle("/", Router())
	log.Fatal(http.ListenAndServeTLS(":8081", "selfsigned.crt", "selfsigned.key", nil))
}

// Router returns a router with registered handlers.
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/events", EventsHandler).Methods("GET")
	return r
}

// EventsHandler sends a simple message every 5 seconds.
func EventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	notify := w.(http.CloseNotifier).CloseNotify()
	conn := true
	go func() {
		<-notify
		conn = false
	}()

	for conn {
		msg := events.Message{Text: time.Now().String()}
		jsn, _ := json.Marshal(msg)
		fmt.Fprintf(w, "data: %s\n\n", jsn)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		} else {
			log.Println("no flush")
		}
		time.Sleep(5 * time.Second)
	}
	log.Println("conn closed")
}
