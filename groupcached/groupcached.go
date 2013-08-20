package main

import (
		"fmt"
		"os"
		"github.com/golang/groupcache"
)

func main() {
	me := "http://10.229.13.84:9527"
	peers := groupcache.NewHTTPPool(me)
	fmt.Fprintf(os.Stdout, "peers:%s", peers)
	
	var echo = groupcache.NewGroup("echo_server", 64 << 20, groupcache.GetterFunc(
					func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
						fmt.Fprintf(os.Stdout, "groupcache getter %s\n", key)
						dest.SetString(key)
						return nil
					}))

	var buffer string
	err := echo.Get(nil, "test", groupcache.StringSink(&buffer))
	fmt.Fprintf(os.Stdout, "groupcache get %s %s=%s\n", "test", err, buffer)
}
