package tidal

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	token   string = "MbjR4DLXz1ghC4rV"
	urlBase string = "https://api.tidalhifi.com/v1/"
)

func makeRequest(x string, urlExt string, data url.Values) (*http.Response, error) {
	u, _ := url.Parse(urlBase + urlExt)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("token", token)
	u.RawQuery = q.Encode()
	builtURL := u.String()
	req, _ := http.NewRequest(x, builtURL, strings.NewReader(data.Encode()))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("token", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *Tidal) Login(username string, password string) error {
	data := url.Values{}
	data.Add("username", username)
	data.Add("password", password)
	resp, err := makeRequest("POST", "login/username", data)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Failed to login")
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return err
	}
	log.Println(t.UserID)
	return nil
}
