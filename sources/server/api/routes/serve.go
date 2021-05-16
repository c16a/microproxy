package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func HandleRoute(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse("http://www.google.com")

	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}
