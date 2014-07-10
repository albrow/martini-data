package data

import (
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
	"strings"
)

type Data map[string]string

func Parser() func(c martini.Context, req *http.Request) {
	return func(c martini.Context, req *http.Request) {
		vals := Data{}
		for key, val := range req.URL.Query() {
			vals[key] = val[0]
		}
		contentType := req.Header.Get("Content-Type")
		if strings.Contains(contentType, "multipart/form-data") {
			if err := req.ParseMultipartForm(2048); err != nil {
				panic(err)
			}
			// convert from map[string][]string to map[string]string
			// the first value for each key is considered the right value
			for key, val := range req.MultipartForm.Value {
				vals[key] = val[0]
			}
		} else if strings.Contains(contentType, "form-urlencoded") {
			if err := req.ParseForm(); err != nil {
				panic(err)
			}
			// convert from map[string][]string to map[string]string
			// the first value for each key is considered the right value
			for key, val := range req.PostForm {
				vals[key] = val[0]
			}
		}
		c.Map(vals)
	}
}

func (d Data) KeyExists(key string) bool {
	_, found := d[key]
	return found
}

func (d Data) Get(key string) string {
	return d[key]
}

func (d Data) GetInt(key string) int {
	str, found := d[key]
	if !found {
		return 0
	} else {
		if result, err := strconv.Atoi(str); err != nil {
			panic(err)
		} else {
			return result
		}
	}
}

func (d Data) GetStrings(key string) []string {
	return strings.Split(d[key], ",")
}
