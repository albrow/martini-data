package data

import (
	"github.com/albrow/go-data-parser"
	"github.com/go-martini/martini"
	"net/http"
)

type Data map[string]string

func Parser() func(c martini.Context, req *http.Request) {
	return func(c martini.Context, req *http.Request) {
		if d, err := data.Parse(req); err != nil {
			panic(err)
		} else {
			c.Map(d)
		}
	}
}
