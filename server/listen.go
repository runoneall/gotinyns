package server

import (
	"os"
	"strconv"

	"github.com/miekg/dns"
)

var sem chan struct{}

func Listen() {
	MAX_CONCURRENCY, err := strconv.Atoi(os.Getenv("MAX_CONCURRENCY"))
	if err != nil {
		panic(err)
	}

	sem = make(chan struct{}, MAX_CONCURRENCY)

	dns.HandleFunc(".", handler)
	server := &dns.Server{Addr: ":53", Net: "udp"}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
