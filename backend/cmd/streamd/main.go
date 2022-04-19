package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	byteChunks := []byte("AAGGJJMMZZ$$")
	http.HandleFunc("/audio", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}
		w.Header().Set("Transfer-encoding", "chunked")
		w.Header().Set("Connection", "Keep-Alive")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		for i := 0; i <= 11; i += 2 {
			w.Write([]byte(fmt.Sprintln(byteChunks[i], byteChunks[i+1])))
			flusher.Flush() // Trigger "chunked" encoding and send a chunk...
			time.Sleep(1000 * time.Millisecond)
		}

	})

	fmt.Println("Server started at localhost 9000")
	http.ListenAndServe(":9000", nil)

}
