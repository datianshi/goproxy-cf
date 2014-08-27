package main

import (
	"github.com/datianshi/goproxy-cf/reverseproxy"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	cfg := config()

	for key, value := range *cfg {
		remote, err := url.Parse(value)
		if err != nil {
			panic(err)
		}
		proxy := reverseproxy.NewReverseProxy(remote)

		log.Printf("Map %[1]s to %[2]s", key, value)
		http.HandleFunc(key, handler(proxy))
	}

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}
