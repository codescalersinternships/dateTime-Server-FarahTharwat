package pkg

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestDateTimeEndpointGin(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dateTime", nil)
	app := NewApp()
	app.GetRouter().ServeHTTP(recorder, req)
	t.Run("Returns 200 status code", func(t *testing.T) {
		assertEqual(t, 200, recorder.Code)
	})
	t.Run("Returns tuesday 15:16 as response", func(t *testing.T) {
		assertEqual(t, "tuesday 15:16", recorder.Body.String())
	})

}

func assertEqual(t *testing.T, want, got any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v \n want %v", got, want)
	}
}
