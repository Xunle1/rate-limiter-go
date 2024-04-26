package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xunle/rate-limiter-go/token_bucket"
)

var tb *token_bucket.TokenBucket

func main() {
	tb = token_bucket.DefaultTokenBucket()
	http.HandleFunc("/echo", echo)
	log.Println("Start serving...")
	http.ListenAndServe("localhost:8080", nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	if tb.Get() != nil {
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}
	msg := r.URL.Query().Get("msg")
	if msg == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(&struct {
			ResponseMessage string
		}{
			ResponseMessage: "No 'msg' query param.",
		})
		w.Write(resp)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("echo: " + msg))
}
