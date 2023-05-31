# README

This server has been generated by [Autostrada](https://autostrada.dev/) as boilerplate for Rundoo server take-home assignment.


## Getting started

Make sure that you're in the root of the project directory, fetch the dependencies with `go mod tidy`, then run the application using `go run ./cmd/web`:

```
$ go mod tidy
$ go run ./cmd/web
```

## Project Limitations:
- SKU are 6 characters (can be characters, numbers etc.)
- Category is defined as HARDWARE, PAINT, WOOD for focused definition
- Name is freeform as string and can have spaces
- Cannot have duplicate SKUs in DB

Then visit [http://localhost:4444](http://localhost:4444) in your browser to ensure that your service is running.

## Notes during Development process
- mongo driver for primitive ID 
- set up Tests (signatures) for repository, usecase implementation per requirements (TDD)
- implement API, but with minimal specifications for user ease and scope
- enforced SKU to 6 character
- made categories defaulted to three types
- name can include spaces with key mapping
- full-text string match means exact match on key
- dedupes against multiple databases

Retrospective Overview:
- had a bunch of small issues resolving (old) laptop with GO configuration issues + new project
- decided to try to implement TDD for best practices and wrote (some) unit tests first before any actual development

