package main

import (
	"log"
	"net/http"
	"time"

	"github.com/codescalersinternships/dateTime-Server-FarahTharwat/pkg"
)

func main() {
	dateTimeFunc := func() string {
		return time.Now().Format(time.RFC1123)
	}
	http.Handle("/dateTime", pkg.DateTimeEndpoint(dateTimeFunc))
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
