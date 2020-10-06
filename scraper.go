package scraper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

// Result represents incoming JSON data from Prompt API - Scraper
type Result struct {
	Data         string                 `json:"data,omitempty"`
	DataSelector []string               `json:"data-selector,omitempty"`
	URL          string                 `json:"url"`
	Headers      map[string]interface{} `json:"headers"`
}

// ErrorResponse represents response errors from Prompt API - Scraper
type ErrorResponse struct {
	Message string `json:"message"`
}

// Params represents arguments for Scrape method
type Params struct {
	URL          string
	AuthPassword string
	AuthUsername string
	Cookie       string
	Country      string
	Referer      string
	Selector     string
}

type extraHeader struct {
	name  string
	value string
}

var promptAPIEndpoint = "https://api.promptapi.com/scraper"

// PromptAPI represents type
type PromptAPI struct{}

// Scrape makes API call to Prompt API - Scraper and returns result
func (pa PromptAPI) Scrape(params *Params, headers []*extraHeader, result *Result) error {
	apiKey, ok := os.LookupEnv("PROMPTAPI_TOKEN")
	if !ok {
		return errors.New("You need to set PROMPTAPI_TOKEN environment variable")
	}

	_, err := url.ParseRequestURI(params.URL)
	if err != nil {
		return err
	}

	v := url.Values{}
	v.Set("url", params.URL)

	if params.AuthPassword != "" {
		v.Set("auth_password", params.AuthPassword)
	}
	if params.AuthUsername != "" {
		v.Set("auth_username", params.AuthUsername)
	}
	if params.Cookie != "" {
		v.Set("cookie", params.Cookie)
	}
	if params.Country != "" {
		v.Set("country", params.Country)
	}
	if params.Referer != "" {
		v.Set("referer", params.Referer)
	}
	if params.Selector != "" {
		v.Set("selector", params.Selector)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	requiredURL := fmt.Sprintf("%s?%s", promptAPIEndpoint, v.Encode())

	req, err := http.NewRequest("GET", requiredURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("apikey", apiKey)
	if len(headers) > 0 {
		for _, h := range headers {
			req.Header.Set(h.name, h.value)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		msg := new(ErrorResponse)
		err := json.Unmarshal(body, msg)
		if err != nil {
			return err
		}
		return errors.New(msg.Message)
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	return nil
}

// Save saves fetched data to given file
func (pa PromptAPI) Save(filename string, result *Result) (n int, err error) {
	targetExtension := ".html"
	saveData := result.Data

	if len(result.DataSelector) > 0 {
		targetExtension = ".json"
		buffer := new(bytes.Buffer)
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)

		err = encoder.Encode(result.DataSelector)
		if err != nil {
			return
		}
		saveData = buffer.String()
	}

	fileExtension := filepath.Ext(filename)
	fileBasename := filename[0 : len(filename)-len(fileExtension)]
	fileTargetname := fileBasename + targetExtension

	f, err := os.Create(fileTargetname)
	if err != nil {
		return
	}
	defer f.Close()

	n, errWrite := f.WriteString(saveData)
	if errWrite != nil {
		return
	}
	return
}
