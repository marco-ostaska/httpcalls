# httpcalls

[![GoDoc](https://godoc.org/github.com/marco-ostaska/httpcalls?status.svg)](https://godoc.org/github.com/marco-ostaska/httpcalls)
[![Go Report Card](https://goreportcard.com/badge/github.com/marco-ostaska/httpcalls)](https://goreportcard.com/report/github.com/marco-ostaska/httpcalls)
[![Build Status](https://travis-ci.com/marco-ostaska/httpcalls.svg?branch=main)](https://travis-ci.com/marco-ostaska/httpcalls)

    import "github.com/marco-ostaska/httpcalls"

Package httpcalls is a wrapper to make GET, DELETE and POST requests for apis

## Usage

```go
var ErrInvalidCertificate = errors.New("you may try running it using --insecure to skip certificate validation")
```
ErrInvalidCertificate is returned by http call when try to secure connect in an
environment with an invalid certificate

```go
var Insecure bool
```
Insecure variable skips certificate validation when true

#### type APIData

```go
type APIData struct {
	API       string    // API section /api/xyz
	ARGS      string    // any extra arguments to be passed to API
	Payload   io.Reader // payload for posting
	Result    []byte    // result body from call
	AuthKey   string    // authorization key (e.g "Authorization")
	AuthValue string    // authorization value (e.g "Basic xyz")
	URL       string    // URL
}
```

APIData an abstraction to API

#### func (*APIData) DeleteRequest

```go
func (a *APIData) DeleteRequest() error
```
DeleteRequest make delete call to sl1 API

#### func (*APIData) NewPost

```go
func (a *APIData) NewPost() error
```
NewPost make new post to sl1 API

#### func (*APIData) NewRequest

```go
func (a *APIData) NewRequest(v interface{}) error
```
NewRequest make new call to sl1 API and unmarshal the call result to given
struct pointer

#### func (*APIData) QueryGraphQL

```go
func (a *APIData) QueryGraphQL(query string, v interface{}) error
```
QueryGraphQL queries graphql
