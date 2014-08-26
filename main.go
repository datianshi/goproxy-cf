package main

import (
	"github.com/datianshi/goproxy-cf/reverseproxy"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	remote, err := url.Parse("http://liberty-music.cfapps.io")
	if err != nil {
		panic(err)
	}

	proxy := reverseproxy.NewSingleHostReverseProxy(remote)
	log.Println(remote.Host)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func handler(p *reverseproxy.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		p.ServeHTTP(w, r)
	}
}
