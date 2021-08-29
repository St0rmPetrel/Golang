package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/St0rmPetrel/Golang/jun_test_1/db"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkResponseHeader(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected response code %v. Got %v\n", expected, actual)
	}
}

func TestGet(t *testing.T) {
	rdb, err := db.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	app := setupApp(rdb)

	req := httptest.NewRequest("GET", "/json/hackers", nil)

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	checkResponseCode(t, http.StatusOK, res.StatusCode)
	checkResponseHeader(t, "application/json", res.Header.Get("Content-Type"))
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
}
