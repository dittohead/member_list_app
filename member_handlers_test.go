package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetMembersHandler(t *testing.T) {
	members = []Member{
		{1, "Richard Hammond", "r.hammond@gmail.com", "01.01.1982"},
	}
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(getMemberHandler)
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Invalid status code: got %v, should be %v",
			status, http.StatusOK)
	}

	expected := Member{1, "Richard Hammond", "r.hammond@gmail.com", "01.01.1982"}

	bm := []Member{}
	err = json.NewDecoder(recorder.Body).Decode(&bm)
	if err != nil {
		t.Fatal(err)
	}

	actual := bm[0]
	if actual != expected {
		t.Errorf("Corrupted body: got %v want %v", actual, expected)
	}
}

func TestCreateMembersHandler(t *testing.T) {

	members = []Member{
		{1, "J.Clarkson", "hate.ferrari@amazon.com", "02.12.1994"},
	}

	form := newCreateMembersForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(createMemberHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("Wrong status code: got %v, expect %v",
			status, http.StatusOK)
	}

	expected := Member{1, "James May", "James@May.dan", getRegDate()}

	if err != nil {
		t.Fatal(err)
	}

	actual := members[1]

	if actual != expected {
		t.Errorf("Wrong body reiceved: got %v, expect %v", actual, expected)
	}
}

func newCreateMembersForm() *url.Values {
	form := url.Values{}
	form.Set("name", "James May")
	form.Set("email", "James@May.dan")
	return &form
}
