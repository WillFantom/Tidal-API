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
