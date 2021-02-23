package main

import (
	"github.com/golang/glog"

	"go-intro/rpc/mailbox/mailbox"

	"flag"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	flag.Parse()

	mbox := mailbox.NewMailboxService()

	rpc.Register(mbox)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":8000")
	if e != nil {
		glog.Fatalf("Error listening on :8000: %v", e)
	}

	glog.Infoln("Listening on port 8000...")

	http.Serve(l, nil)
}
