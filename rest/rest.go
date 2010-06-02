// Wraps the http package with a HTTP method and header aware muxer.
// Code derived from the http package implementation of DefaultServeMux.

package rest

import (
	"http"
	"log"
	"container/vector"
	"os"
	"path"
	"regexp"
)

// Does path match pattern?
func pathMatch(pattern, path string) bool {
	if len(pattern) == 0 {
		// should not happen
		return false
	}
	n := len(pattern)
	if pattern[n-1] != '/' {
		return pattern == path
	}
	return len(path) >= n && path[0:n] == pattern
}

// Return the canonical path for p, eliminating . and .. elements.
func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash except for root;
	// put the trailing slash back if necessary.
	if p[len(p)-1] == '/' && np != "/" {
		np += "/"
	}
	return np
}

// Is the request method in the accepted methods list? A nil accepted methods list accepts all methods
func methodMatch(methods []string, requested string) bool {
	if methods == nil {
		return true
	}
	for _, method := range methods {
		if method == requested {
			return true
		}
	}
	return false
}

// Do the headers match every required header?
func requiredHeadersMatch(requiredHeaders map[string]string, req *http.Request) bool {
	if requiredHeaders == nil {
		return true
	}

	log.Stdoutf("Checking %#v against %#v", requiredHeaders, req)

	for k, v := range requiredHeaders {
		switch k {
		default:
			if match, _ := regexp.MatchString(v, req.Header[k]); !match {
				return false
			}
		case "User-Agent":
			if match, _ := regexp.MatchString(v, req.UserAgent); !match {
				return false
			}
		case "Host":
			if match, _ := regexp.MatchString(v, req.Host); !match {
				return false
			}
		}
	}
	return true
}

type RestRoute struct {
	Pattern         string
	Methods         []string
	RequiredHeaders map[string]string
	Handler         http.Handler
}

type RestMux struct {
	v *vector.Vector
}

func NewRestMux() *RestMux { return &RestMux{new(vector.Vector)} }

var DefaultRestMux = NewRestMux()

func (mux *RestMux) ServeHTTP(c *http.Conn, req *http.Request) {
	log.Stdoutf("%s on %s", req.Method, req.URL)
	// Clean path to canonical form and redirect.
	if p := cleanPath(req.URL.Path); p != req.URL.Path {
		c.SetHeader("Location", p)
		c.WriteHeader(http.StatusMovedPermanently)
		return
	}


	// Most-specific (longest) pattern wins.
	var h http.Handler
	var n = 0
	for i := 0; i < mux.v.Len(); i++ {
		var r = mux.v.At(i).(RestRoute)
		// match path
		if !pathMatch(r.Pattern, req.URL.Path) {
			continue
		} else if !methodMatch(r.Methods, req.Method) {
			continue
		} else if !requiredHeadersMatch(r.RequiredHeaders, req) {
			continue
		}

		// longest (most specific) pattern wins
		if h == nil || len(r.Pattern) > n {
			n = len(r.Pattern)
			h = r.Handler
		}
	}
	if h == nil {
		h = http.NotFoundHandler()
	}
	h.ServeHTTP(c, req)
}

func Handle(pattern string, handler http.Handler, methods []string, requiredHeaders map[string]string) {
	var route RestRoute

	route.Pattern = pattern
	route.Handler = handler
	route.Methods = methods
	route.RequiredHeaders = requiredHeaders

	DefaultRestMux.v.Push(route)
}

// Wrapper for http.ListenAndServe that replaces DefaultServeMux with
// DefaultRestMux
func ListenAndServe(addr string, handler http.Handler) os.Error {
	if handler == nil {
		handler = DefaultRestMux
	}
	log.Stdoutf("Starting Server on %s", addr)
	return http.ListenAndServe(addr, handler)
}
