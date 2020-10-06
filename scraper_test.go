package scraper

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestScrapeWithoutPromptAPIToken(t *testing.T) {
	if os.Getenv("PROMPTAPI_TOKEN") != "" {
		t.Skip("skipping test...")
	}

	s := new(PromptAPI)
	params := new(Params)
	result := new(Result)
	extraHeaders := []*ExtraHeader{}

	if err := s.Scrape(params, extraHeaders, result); err.Error() != "You need to set PROMPTAPI_TOKEN environment variable" {
		t.Errorf("got: %v", err)
	}
}

func TestScrapeBasicRequest(t *testing.T) {
	if os.Getenv("PROMPTAPI_TOKEN") == "" {
		t.Skip("skipping test...")
	}

	s := new(PromptAPI)
	params := &Params{
		URL: "https://pypi.org/classifiers/",
	}

	result := new(Result)
	extraHeaders := []*ExtraHeader{}

	if err := s.Scrape(params, extraHeaders, result); err != nil {
		t.Errorf("got: %v", err)
	}
	if result.URL != params.URL {
		t.Errorf("want: %s, got: %s", params.URL, result.URL)
	}
	if len(result.Data) < 300000 {
		t.Errorf("fetced data length should be greater than 300000 bytes, got: %d", len(result.Data))
	}
}

func TestScrapeBasicRequestWithSave(t *testing.T) {
	if os.Getenv("PROMPTAPI_TOKEN") == "" {
		t.Skip("skipping test...")
	}

	s := new(PromptAPI)
	params := &Params{
		URL: "https://pypi.org/classifiers/",
	}

	result := new(Result)
	extraHeaders := []*ExtraHeader{}

	if err := s.Scrape(params, extraHeaders, result); err != nil {
		t.Errorf("got: %v", err)
	}
	if result.URL != params.URL {
		t.Errorf("want: %s, got: %s", params.URL, result.URL)
	}
	if len(result.Data) < 300000 {
		t.Errorf("fetced data length should be greater than 300000 bytes, got: %d", len(result.Data))
	}

	fileSize, err := s.Save("/tmp/test.html", result)
	if err != nil {
		t.Errorf("got: %v", err)
	}

	if fileSize <= 0 {
		t.Errorf("saved file size should be greater than 0 bytes, got: %d", fileSize)
	}
}

func TestScrapeComplexRequest(t *testing.T) {
	if os.Getenv("PROMPTAPI_TOKEN") == "" {
		t.Skip("skipping test...")
	}

	s := new(PromptAPI)
	params := &Params{
		URL:      "https://pypi.org/classifiers/",
		Country:  "EE",
		Selector: "ul li button[data-clipboard-text]",
	}
	result := new(Result)
	extraHeaders := []*ExtraHeader{}

	if err := s.Scrape(params, extraHeaders, result); err != nil {
		t.Errorf("got: %v", err)
	}
	if result.URL != params.URL {
		t.Errorf("want: %s, got: %s", params.URL, result.URL)
	}
	if len(result.DataSelector) < 700 {
		t.Errorf("fetced array length should be greater than 700 items, got: %d", len(result.DataSelector))
	}
}

func ExamplePromptAPI_Scrape() {
	s := new(PromptAPI)
	params := &Params{
		URL: "https://pypi.org/classifiers/",
	}

	result := new(Result)
	extraHeaders := []*ExtraHeader{}

	err := s.Scrape(params, extraHeaders, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Length of incoming data: %d\n", len(result.Data))
	fmt.Printf("Response headers: %v\n", result.Headers)
	fmt.Printf("Content-Length: %v\n", result.Headers["Content-Length"])
}

func ExamplePromptAPI_Scrape_with_extra_headers() {
	s := new(PromptAPI)
	params := &Params{
		URL: "https://pypi.org/classifiers/",
	}

	result := new(Result)
	extraHeaders := []*ExtraHeader{
		&ExtraHeader{
			Name:  "X-Referer",
			Value: "https://www.google.com",
		},
	}

	err := s.Scrape(params, extraHeaders, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Length of incoming data: %d\n", len(result.Data))
	fmt.Printf("Response headers: %v\n", result.Headers)
	fmt.Printf("Content-Length: %v\n", result.Headers["Content-Length"])
}

func ExamplePromptAPI_Save() {
	s := new(PromptAPI)
	params := &Params{
		URL: "https://pypi.org/classifiers/",
	}

	result := new(Result)
	extraHeaders := []*ExtraHeader{}

	err := s.Scrape(params, extraHeaders, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Length of incoming data: %d\n", len(result.Data))
	fmt.Printf("Response headers: %v\n", result.Headers)
	fmt.Printf("Content-Length: %v\n", result.Headers["Content-Length"])

	fileSize, err := s.Save("/tmp/test.html", result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Size of /tmp/test.html -> %d bytes\n", fileSize)
}
