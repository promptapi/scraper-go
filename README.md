![Version](https://img.shields.io/badge/version-0.0.0-orange.svg)
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

## Example Usage

```go
// main.go

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
