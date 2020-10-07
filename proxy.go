package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

//CODE is runnable code
var (
	CODE   string
	URL    string
	LISTEN string
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func logRequest(req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	log.Println("request body:\n", string(body))
	log.Println("request headers:\n", req.Header)
}

func logResponse(res *http.Response) {
	body, _ := ioutil.ReadAll(res.Body)
	res.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	log.Println("response body:\n", string(body))
	log.Println("response headers:\n", res.Header)
}

func redirectRequest(res http.ResponseWriter, req *http.Request) {
	urlReverse, _ := url.Parse(URL)

	proxy := httputil.NewSingleHostReverseProxy(urlReverse)

	director := proxy.Director
	proxy.Director = func(req *http.Request) {

		req.URL.Host = urlReverse.Host
		req.URL.Scheme = urlReverse.Scheme
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))

		log.Printf("request %s%s redirect to %s%s", req.Host, req.URL.Path, urlReverse.Host, req.URL.Path)

		logRequest(req)

		director(req)
	}

	proxy.ModifyResponse = func(res *http.Response) error {
		i := interp.New(interp.Options{})

		i.Use(stdlib.Symbols)

		_, err := i.Eval(CODE)
		if err != nil {
			panic(err)
		}

		v, err := i.Eval("temp.Run")
		if err != nil {
			panic(err)
		}

		run := v.Interface().(func(*http.Response))

		run(res)

		logResponse(res)

		return nil
	}

	proxy.ServeHTTP(res, req)

}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	redirectRequest(res, req)
}

func server() {
	http.HandleFunc("/", handleRequestAndRedirect)

	if err := http.ListenAndServe(LISTEN, nil); err != nil {
		panic(err)
	}
}

func setUp() {
	log.Println("starting reverse proxy")
	CODE = strings.ReplaceAll(getEnv("CODE", ""), "?", "\n")
	log.Println("go runtime code", CODE)
	URL = getEnv("REVERSE_PROXY_SERVER_REDIRECT_URL", "")
	log.Println("redirect url", URL)
	LISTEN = getEnv("REVERSE_PROXY_SERVER", "")
	log.Println("listening", LISTEN)
}

func main() {
	setUp()
	server()
}
