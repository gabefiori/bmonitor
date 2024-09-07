package main

import (
	"github.com/bmonitor/server"
	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func init() {
	spinhttp.Handle(server.Handle)
}

func main() {}
