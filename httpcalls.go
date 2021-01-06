/*
Copyright © 2020 Marco Ostaska

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

// Package httpcalls is a wrapper to make GET, DELETE and POST requests for apis
package httpcalls

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Insecure variable skips certificate validation when true
var Insecure bool

// ErrInvalidCertificate is returned by http call when try to secure connect in an environment with an invalid certificate
var ErrInvalidCertificate = errors.New("you may try running it using --insecure to skip certificate validation")

// APIData an abstraction to API
type APIData struct {
	API       string    // API section /api/xyz
	ARGS      string    // any extra arguments to be passed to API
	Payload   io.Reader // payload for posting
	Result    []byte    // result body from call
	AuthKey   string    // authorization key (e.g "Authorization")
	AuthValue string    // authorization value (e.g "Basic xyz")
	URL       string    // URL
}

// DeleteRequest make delete call to sl1 API
func (a *APIData) DeleteRequest() error {
	return a.httpcalls("DELETE")

}

// NewPost make new post to sl1 API
func (a *APIData) NewPost() error {
	return a.httpcalls("POST")
}

// NewRequest make new call to sl1 API and unmarshal the call result to given struct pointer
func (a *APIData) NewRequest(v interface{}) error {
	err := a.httpcalls("GET")
	if err != nil {
		return err
	}

	return json.Unmarshal(a.Result, &v)
}

// isReachable checks if url is reachable
// used for internal purposes only
func isReachable(url string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: Insecure},
	}

	timeout := time.Duration(15 * time.Second)
	c := http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	_, err := c.Get(url)
	if err != nil {
		if strings.Contains(err.Error(), "certificate") {
			return fmt.Errorf("%v\n\n%v", err, ErrInvalidCertificate)
		}
	}

	return err
}

// httpcalls make the http calls
// GET, POST and DELETE
func (a *APIData) httpcalls(method string) error {
	if err := isReachable(a.URL); err != nil {
		return err
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: Insecure}
	url := a.URL + a.API + a.ARGS

	client := &http.Client{}
	req, err := http.NewRequest(method, url, a.Payload)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	if a.AuthKey != "" {
		req.Header.Add(a.AuthKey, a.AuthValue)
	}

	res, err := client.Do(req)
	defer func() {
		cerr := res.Body.Close()
		if err == nil {
			err = cerr
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	a.Result = body

	return nil
}

// QueryGraphQL queries graphql
func (a *APIData) QueryGraphQL(query string, v interface{}) error {
	q := map[string]string{"query": query}
	jsonValue, err := json.Marshal(q)
	if err != nil {
		return err
	}
	a.Payload = bytes.NewBuffer(jsonValue)

	if err := a.NewPost(); err != nil {
		return err
	}

	return json.Unmarshal(a.Result, &v)

}
