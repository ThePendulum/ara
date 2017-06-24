package main

import (
	"net/http"
	"strconv"
	"net"
	"log/syslog"
	"ara/color"
)

func server(logger *syslog.Writer, hsv *color.Hsv) {
	address := net.JoinHostPort(config.Server.Host, strconv.FormatInt(int64(config.Server.Port), 10))
	http.ListenAndServe(address, http.FileServer(http.Dir("./public")))
}
