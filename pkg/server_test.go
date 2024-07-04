package pkg

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDateTimeEndpoint(t *testing.T) {
	/* Creating a request to pass to our handler. We don't have any query parameters for now, so we'll
	   pass 'nil' as the third parameter. */
	req, err := http.NewRequest("GET", "/dateTime", nil)
	if err != nil {
		t.Fatal(err)
	}

	/* Creating a ResponseRecorder (which satisfies http.ResponseWriter) to record the response. */
	rr := httptest.NewRecorder()
	handler := DateTimeEndpoint(dateTimeNow)

	/* Calling their ServeHTTP method directly and pass in our Request and ResponseRecorder. */
	handler.ServeHTTP(rr, req)

	/* Checking the status code is what we expect. */
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	/* Checking the response body is what we expect. */
	expected := "tuesday 15:16"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// server -> on access -> prints date time
	// client -> on access -> receives date time
	//     -> should i test if that's now? (acceptable delta)
	//     -> should i test if the return is a valid datetime ?

}
