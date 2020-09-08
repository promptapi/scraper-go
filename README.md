![Version](https://img.shields.io/badge/version-0.1.1-orange.svg)
![Go](https://img.shields.io/badge/go-1.15.1-black.svg)
[![Documentation](https://godoc.org/github.com/promptapi/scraper-go?status.svg)](https://pkg.go.dev/github.com/promptapi/scraper-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/promptapi/scraper-go)](https://goreportcard.com/report/github.com/promptapi/scraper-go)
[![Build Status](https://travis-ci.org/promptapi/scraper-go.svg?branch=main)](https://travis-ci.org/promptapi/scraper-go)

# Prompt API - Scraper - Golang Package

`PromptAPI` struct is a simple golang wrapper for [scraper api][scraper-api]
with few more extra cream and sugar.

## Requirements

1. You need to signup for [Prompt API][promptapi-signup]
1. You need to subscribe [scraper api][scraper-api], test drive is **free!!!**
1. You need to set `PROMPTAPI_TOKEN` environment variable after subscription.

then;

```bash
$ go get -u github.com/promptapi/scraper-go
```

## Example Basic Usage

```go
// main.go
package main

import (
	"fmt"
	"log"

	scraper "github.com/promptapi/scraper-go"
)

func main() {
	s := new(scraper.PromptAPI)

	params := new(scraper.Params)
	params.URL = "https://pypi.org/classifiers/"
	params.Country = "EE"

	result := new(scraper.Result)

	err := s.Scrape(params, result)
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
```

Run:

```bash
$ go run main.go 
Length of incoming data: 321322
Response headers: map[Accept-Ranges:bytes Content-Length:321322 Content-Security-Policy:base-uri 'self'; block-all-mixed-content; connect-src 'self' https://api.github.com/repos/ *.fastly-insights.com sentry.io https://api.pwnedpasswords.com https://2p66nmmycsj3.statuspage.io; default-src 'none'; font-src 'self' fonts.gstatic.com; form-action 'self'; frame-ancestors 'none'; frame-src 'none'; img-src 'self' https://warehouse-camo.ingress.cmh1.psfhosted.org/ www.google-analytics.com *.fastly-insights.com; script-src 'self' www.googletagmanager.com www.google-analytics.com *.fastly-insights.com https://cdn.ravenjs.com; style-src 'self' fonts.googleapis.com; worker-src *.fastly-insights.com Content-Type:text/html; charset=UTF-8 Date:Tue, 08 Sep 2020 19:10:24 GMT ETag:"1ea9p+Hscl37dEKelacPWw" Referrer-Policy:origin-when-cross-origin Strict-Transport-Security:max-age=31536000; includeSubDomains; preload Vary:Accept-Encoding, Cookie, Accept-Encoding X-Cache:MISS, HIT X-Cache-Hits:0, 1 X-Content-Type-Options:nosniff X-Frame-Options:deny X-Permitted-Cross-Domain-Policies:none X-Served-By:cache-bwi5127-BWI, cache-hhn4035-HHN X-Timer:S1599592224.395422,VS0,VE247 X-XSS-Protection:1; mode=block]
Content-Length: 321322
Size of /tmp/test.html -> 321322 bytes
```

You can add url parameters for extra operations. Valid parameters are:

- `AuthPassword`: for HTTP Realm auth password
- `AuthUsername`: for HTTP Realm auth username
- `Cookie`: URL Encoded cookie header.
- `Country`: 2 character country code. If you wish to scrape from an IP address of a specific country.
- `Referer`: HTTP referer header
- `Selector`: CSS style selector path such as `a.btn div li`. If `Selector` is
  enabled, returning result will be collection of data and saved file will be
  in `.json` format.

Example with `Selector`:

```go
// main.go
package main

import (
	"fmt"
	"log"

	scraper "github.com/promptapi/scraper-go"
)

func main() {
	s := new(scraper.PromptAPI)

	params := new(scraper.Params)
	params.URL = "https://pypi.org/classifiers/"
	params.Country = "EE"
	params.Selector = "ul li button[data-clipboard-text]"

	result := new(scraper.Result)

	err := s.Scrape(params, result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Length of incoming data: %d\n", len(result.Data))
	fmt.Printf("Length of extracted data: %d\n", len(result.DataSelector))
	fmt.Printf("Response headers: %v\n", result.Headers)
	fmt.Printf("Content-Length: %v\n", result.Headers["Content-Length"])

	fileSize, err := s.Save("/tmp/test.json", result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Size of /tmp/test.json -> %d bytes\n", fileSize)

}
```

Run:

```bash
$ go run main.go 
Length of incoming data: 0
Length of extracted data: 734
Response headers: map[Accept-Ranges:bytes Content-Length:321322 Content-Security-Policy:base-uri 'self'; block-all-mixed-content; connect-src 'self' https://api.github.com/repos/ *.fastly-insights.com sentry.io https://api.pwnedpasswords.com https://2p66nmmycsj3.statuspage.io; default-src 'none'; font-src 'self' fonts.gstatic.com; form-action 'self'; frame-ancestors 'none'; frame-src 'none'; img-src 'self' https://warehouse-camo.ingress.cmh1.psfhosted.org/ www.google-analytics.com *.fastly-insights.com; script-src 'self' www.googletagmanager.com www.google-analytics.com *.fastly-insights.com https://cdn.ravenjs.com; style-src 'self' fonts.googleapis.com; worker-src *.fastly-insights.com Content-Type:text/html; charset=UTF-8 Date:Tue, 08 Sep 2020 19:17:22 GMT ETag:"1ea9p+Hscl37dEKelacPWw" Referrer-Policy:origin-when-cross-origin Strict-Transport-Security:max-age=31536000; includeSubDomains; preload Vary:Accept-Encoding, Cookie, Accept-Encoding X-Cache:HIT, HIT X-Cache-Hits:1, 1 X-Content-Type-Options:nosniff X-Frame-Options:deny X-Permitted-Cross-Domain-Policies:none X-Served-By:cache-bwi5137-BWI, cache-bma1621-BMA X-Timer:S1599592641.178639,VS0,VE1512 X-XSS-Protection:1; mode=block]
Content-Length: 321322
Size of /tmp/test.json -> 173717 bytes
```

Let’s see `/tmp/test.json` file:

```json
[
  "<button class=\"button button--small margin-top margin-bottom copy-tooltip copy-tooltip-w\" data-clipboard-text=\"Development Status :: 1 - Planning\" data-tooltip-label=\"Copy to clipboard\" type=\"button\">\n Copy\n</button>\n",
  "<button class=\"button button--small margin-top margin-bottom copy-tooltip copy-tooltip-w\" data-clipboard-text=\"Development Status :: 2 - Pre-Alpha\" data-tooltip-label=\"Copy to clipboard\" type=\"button\">\n Copy\n</button>\n",
  "<button class=\"button button--small margin-top margin-bottom copy-tooltip copy-tooltip-w\" data-clipboard-text=\"Development Status :: 3 - Alpha\" data-tooltip-label=\"Copy to clipboard\" type=\"button\">\n Copy\n</button>\n",
  "<button class=\"button button--small margin-top margin-bottom copy-tooltip copy-tooltip-w\" data-clipboard-text=\"Development Status :: 4 - Beta\" data-tooltip-label=\"Copy to clipboard\" type=\"button\">\n Copy\n</button>\n",
  "<button class=\"button button--small margin-top margin-bottom copy-tooltip copy-tooltip-w\" data-clipboard-text=\"Development Status :: 5 - Production/Stable\" data-tooltip-label=\"Copy to clipboard\" type=\"button\">\n Copy\n</button>\n",
  "<button class=\"button button--small margin-top margin-bottom copy-tooltip copy-tooltip-w\" data-clipboard-text=\"Development Status :: 6 - Mature\" data-tooltip-label=\"Copy to clipboard\" type=\"button\">\n Copy\n</button>\n",
  "<button class=\"button button--small margin-top margin-bottom copy-tooltip copy-tooltip-w\" data-clipboard-text=\"Development Status :: 7 - Inactive\" data-tooltip-label=\"Copy to clipboard\" type=\"button\">\n Copy\n</button>\n",
  ,
  ,
  ,
  ,
  ,
]
```

## Development

Available rake tasks:

```bash
$ rake -T
rake default                    # Default task, show avaliable tasks
rake release:check              # Do release check
rake release:publish[revision]  # Publish project with revision: major,minor,patch, default: patch
rake serve_doc[port]            # Run doc server
rake test[verbose]              # Run tests
```

- Run tests: `rake test` or `rake test[-v]`
- Run doc server: `rake serve_doc` or `rake serve_doc[9000]`

Release package (*if you have write access*):

1. Commit your changes
1. Run `rake release:check`
1. If all goes ok, run `rake release:publish`

---

---

## License

This project is licensed under MIT

---

## Contributer(s)

* [Prompt API](https://github.com/promptapi) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/promptapi/scraper-go/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'Add awesome features...'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

This project is intended to be a safe,
welcoming space for collaboration, and contributors are expected to adhere to
the [code of conduct][coc].

---

[scraper-api]:      https://promptapi.com/marketplace/description/scraper-api
[promptapi-signup]: https://promptapi.com/#signup-form
[coc]:              https://github.com/promptapi/scraper-go/blob/main/CODE_OF_CONDUCT.md
