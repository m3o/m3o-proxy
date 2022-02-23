package proxy

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	API = "api.m3o.com"
)

type Handler struct {
	token string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// set the request url
	r.URL.Scheme = "https"
	r.URL.Host = "api.m3o.com"

	// assuming user passes in /v1/helloworld/Call
	req, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	req.Header.Set("Authorization", "Bearer "+h.token)

	ct := r.Header.Get("Content-Type")

	if len(ct) == 0 {
		ct = "application/json"
	}

	req.Header.Set("Content-Type", ct)

	// set all the headers
	for k, v := range r.Header {
		req.Header.Set(k, strings.Join(v, ","))
	}

	// make request to backend
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rsp.Body.Close()

	// read the data
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// set headers
	for k, v := range rsp.Header {
		w.Header().Set(k, strings.Join(v, ","))
	}
	// set status
	w.WriteHeader(rsp.StatusCode)
	// write the data
	w.Write(b)
}

func New(token string) http.Handler {
	return &Handler{token}
}
