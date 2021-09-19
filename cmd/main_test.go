package main

import (
	"testing"
)

// checking for a valid return value of getJSON().
func TestGetJSON(t *testing.T) {
	url := "https://api.chucknorris.io/jokes/random"
	msg, err := getJSON(url)

	if len(msg) == 0 || err != nil {
		t.Fatalf(`getJSON() = %q, %v, want len(msg) > 0, nil`, msg, err)
	}
}

// checking for a valid return value of getStringFromJSON().
func TestGetStringFromJSON(t *testing.T) {
	s := `{"categories":[],"created_at":"2020-01-05 13:42:28.664997","icon_url":"https://assets.chucknorris.host/img/avatar/chuck-norris.png","id":"5kWKoqquQBGo6Vw9B3OALA","updated_at":"2020-01-05 13:42:28.664997","url":"https://api.chucknorris.io/jokes/5kWKoqquQBGo6Vw9B3OALA","value":"Chuck Norris' apache name is Chuck Norris."}`
	dataJSON := []byte(s)
	msg, err := getStringFromJSON(dataJSON)

	if msg == "" || err != nil {
		t.Fatalf(`getStringFromJSON() = %q, %v, want msg != "", nil`, msg, err)
	}
}

// checking for a valid return value of getRandomJoke().
func TestGetRandomJoke(t *testing.T) {
	if err := getRandomJoke(); err != nil {
		t.Fatalf(`getRandomJoke() = %v, nil`, err)
	}
}

// checking for a valid return value of getCategories().
func TestGetCategories(t *testing.T) {
	msg, err := getCategories()

	if len(msg) == 0 || err != nil {
		t.Fatalf(`getCategories() = %q, %v, want len(msg) > 0, nil`, msg, err)
	}
}

// checking for a valid return value of getRandomJoke().
func TestSaveDump(t *testing.T) {
	if err := getRandomJoke(); err != nil {
		t.Fatalf(`getRandomJoke() = %v, nil`, err)
	}
}
