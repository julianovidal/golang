package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func buildResponse(a string) string {
    code := "<html><head></head><body><a href=\"#\"><img src=\"http://www.iab.net/media/image/300x250.gif\" /></a></body></html>"
    return code
}

func main() {
	
	myServer := &http.Server {
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/display", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		ad := q.Get("ad")
		fmt.Println(ad)
        htmlCodeResponse := buildResponse(ad)
        fmt.Fprintf(w, "%s", htmlCodeResponse)
	})

	log.Fatal(myServer.ListenAndServe())
}