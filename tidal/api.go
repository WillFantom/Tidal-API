package tidal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/mitchellh/mapstructure"
)

const (
	token   string = "MbjR4DLXz1ghC4rV"
	urlBase string = "https://api.tidalhifi.com/v1/"
)

func makeRequest(x string, urlExt string, data url.Values) (*http.Response, error) {
	req, _ := http.NewRequest(x, urlBase+urlExt, strings.NewReader(data.Encode()))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t Tidal) get(data url.Values, urlExt string) ([]byte, error) {
	URL, _ := url.Parse(urlBase + urlExt)
	data.Add("sessionId", t.SessionID)
	data.Add("countryCode", t.CountryCode)
	data.Add("limit", "2")
	URL.RawQuery = data.Encode()
	strURL := URL.String()
	resp, err := http.Get(strURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Bad Request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (t *Tidal) Login(username string, password string) error {
	data := url.Values{}
	data.Add("token", token)
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
	return nil
}

func (t Tidal) Search(term string, zone string) (interface{}, error) {
	searchURL := fmt.Sprintf("search/%s", zone)
	data := url.Values{}
	data.Add("query", term)
	resp, err := t.get(data, searchURL)
	if err != nil {
		return nil, err
	}
	response := SearchResponse{}
	err = json.Unmarshal(resp, &response)
	if err != nil || resp == nil {
		return nil, err
	}
	switch zone {
	case "artists":
		artists := []Artist{}
		mapstructure.Decode(response.Items, &artists)
		return artists, nil
	case "albums":
		albums := []Album{}
		mapstructure.Decode(response.Items, &albums)
		return albums, nil
	case "tracks":
		tracks := []Track{}
		mapstructure.Decode(response.Items, &tracks)
		return tracks, nil
	default:
		return nil, errors.New("Not a valid search zone")
	}
}
