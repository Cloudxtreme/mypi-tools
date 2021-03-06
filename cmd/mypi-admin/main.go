package main

import (
	"os"
	"strings"

	"github.com/dueckminor/mypi-tools/go/ginutil"
	"github.com/dueckminor/mypi-tools/go/pki"
	"github.com/dueckminor/mypi-tools/go/webhandler"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	pki.Setup()

	r := gin.Default()

	wh := &webhandler.WebHandler{}
	err := wh.SetupEndpoints(r)
	if err != nil {
		panic(err)
	}

	mypiAdmin := os.Getenv("MYPI_ADMIN")
	if strings.HasPrefix(mypiAdmin, "http://") || strings.HasPrefix(mypiAdmin, "https://") {
		r.NoRoute(ginutil.SingleHostReverseProxy(mypiAdmin))
	} else {
		if len(mypiAdmin) == 0 {
			mypiAdmin = "dist"
		}
		l := static.LocalFile(mypiAdmin, false)
		r.Use(static.Serve("/", l))
	}

	r.Run()
}
