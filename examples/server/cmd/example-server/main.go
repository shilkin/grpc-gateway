package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/shilkin/grpc-gateway/examples/server"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := server.Run(); err != nil {
		glog.Fatal(err)
	}
}
