package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "strings"
    "io/ioutil"
)

func TestForm(t *testing.T) {
    req := httptest.NewRequest("GET", "/form", nil)
    w := httptest.NewRecorder()
    form(w, req)
    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Got status code %d, %d expected.", resp.StatusCode, http.StatusOK)
    }
}

func TestResultSuccess(t *testing.T) {
    req := httptest.NewRequest("POST", "/result", strings.NewReader("numbers=9+45+1+9+7&goroutines=4"))
    req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
    w := httptest.NewRecorder()
    result(w, req)
    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    expected := "Result: [1 7 9 9 45]"
    if ok := strings.HasPrefix(string(body), expected); !ok {
        t.Errorf("Returned %s, expected: %s.", string(body), expected)
    }
}

func TestResultFailure(t *testing.T) {
    req := httptest.NewRequest("POST", "/result", strings.NewReader("numbers=9+45blah+1+9+7&goroutines=4"))
    req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
    w := httptest.NewRecorder()
    result(w, req)
    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)
    expected := "45blah is not a number."
    if ok := strings.HasPrefix(string(body), expected); !ok || resp.StatusCode != http.StatusBadRequest {
        t.Errorf("Returned %s, expected: %s.", string(body), expected)
    }
}