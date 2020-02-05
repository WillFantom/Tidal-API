package tidal

import (
	"log"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	var tdl Tidal
	username, exists := os.LookupEnv("TIDAL_USER")
	if !exists {
		log.Fatalln("Could not get Tidal username from ENV")
		t.Fail()
	}
	password, exists := os.LookupEnv("TIDAL_PASS")
	if !exists {
		log.Fatalln("Could not get Tidal password from ENV")
		t.Fail()
	}
	err := tdl.Login(username, password)
	if err != nil {
		t.FailNow()
	}
	log.Print("Logged In")
}

func TestSearch(t *testing.T) {
	username, _ := os.LookupEnv("TIDAL_USER")
	password, _ := os.LookupEnv("TIDAL_PASS")
	var tdl Tidal
	tdl.Login(username, password)
	_, err := tdl.Search("El Camino", "albums")
	if err != nil {
		t.FailNow()
	}
	_, err = tdl.Search("Merry Happy", "tracks")
	if err != nil {
		t.FailNow()
	}
	_, err = tdl.Search("Beach Bunny", "artists")
	if err != nil {
		t.FailNow()
	}
	t.Log("Search Success")
}
