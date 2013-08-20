package main

import (
	"flag"

	"github.com/feiyang21687/golang/gohttp"
)

func main() {
	var (
		htdocs  = flag.String("htdocs", "", "the root doc of html file")
		port    = flag.Int("port", 8080, "the http server listen port, default is 8080")
		timeout = flag.Int("timeout", 3, "default read and write timeout seconds, default is 3 seconds")
	)
	flag.Parse()

	gohttp := gohttp.Server(*htdocs, *port, *timeout)
	gohttp.Start()
}
