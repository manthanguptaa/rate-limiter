package main

import (
	"net/http"
	"time"
)

var ipTokenBucket = make(map[string]*TokenBucket)

func handleTokenBucket(w http.ResponseWriter, r *http.Request) {
	ip_address := r.RemoteAddr
	_, ok := ipTokenBucket[ip_address]
	if !ok {
		ipTokenBucket[ip_address] = InitializeTokenBucket()
	}
	valid_request := ipTokenBucket[ip_address].IsRequestAllowed()
	if !valid_request {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("Too many requests!\n"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("You haven't breached the limit!\n"))
	}
}

func main() {
	http.HandleFunc("/token-bucket", handleTokenBucket)

	go func() {
		for {
			time.Sleep(time.Second)
			for _, tb := range ipTokenBucket {
				tb.Refill()
			}
		}
	}()

	http.ListenAndServe(":3000", nil)
}
