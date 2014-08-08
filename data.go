// Package data is a martini middleware wrapper around
// github.com/albrow/go-data-parser.
package data

import (
	"github.com/albrow/go-data-parser"
	"github.com/go-martini/martini"
	"net/http"
)

type Data map[string]string

// Returns a martini middleware which injects a data object into
// the martini context. The data object is composed of form data
// from the body of the request and url query parameters. See
// http://godoc.org/github.com/albrow/go-data-parser
func Parser() martini.Handler {
	return func(c martini.Context, req *http.Request) {
		if d, err := data.Parse(req); err != nil {
			panic(err)
		} else {
			c.Map(d)
		}
	}
}
