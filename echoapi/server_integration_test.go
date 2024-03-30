//go:build integration

package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGetAllUser(t *testing.T) {
	var c Users
	body := bytes.NewBufferString(`{"name":"John Doe", "age":25}`)
	err := request(http.MethodPost, uri("users"), body).Decode(&c)
	if err != nil {
		t.Fatal("can't send request", err)
	}

	var us []Users
	res := request(http.MethodGet, uri("api", "users"), nil)
	err = res.Decode(&us)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Greater(t, len(us), 0)
}

func TestCreateUser(t *testing.T) {
	body := bytes.NewBufferString(`{"name":"Jane Doe", "age":26}`)

	var u Users
	res := request(http.MethodPost, uri("api/users"), body)
	err := res.Decode(&u)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NotEqual(t, 0, u.ID)
	assert.Equal(t, "Jane Doe", u.Name)
	assert.Equal(t, 26, u.Age)
}

func TestGetUserByID(t *testing.T) {
	//c := seedUser(t)
	var u Users
	res := request(http.MethodGet, uri("api", "users", "1"), nil)
	err := res.Decode(&u)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, 1, u.ID)
	assert.Equal(t, "John Doe", u.Name)
	assert.Equal(t, 25, u.Age)
}

func TestUpdateUserById(t *testing.T) {
	t.Skipf("skip this test")
}

func TestDeleteUserById(t *testing.T) {
	t.Skipf("skip this test")
}

func seedUser(t *testing.T) Users {
	var c Users
	body := bytes.NewBufferString(`{"name":"John Doe", "age":25}`)
	err := request(http.MethodPost, uri("users"), body).Decode(&c)

	if err != nil {
		t.Fatal("can't send request", err)
	}

	return c

}

func uri(paths ...string) string {
	host := "http://localhost:8080"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func request(method, url string, body io.Reader) *Response {

	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", "Basic YXBpZGVzaWduOjQ1Njc4")
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}
