package pkg

import (
	"fmt"
	"net/http"
)

func DateTimeEndpoint(dateTime func() string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%v", dateTime())
	}
}
