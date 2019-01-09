package main

import (
	"net/http"
	_ "net/http/pprof"
)

func cpuhogger() {
	var acc uint64
	for {
		acc += 1
		if acc&1 == 0 {
			acc <<= 1
		}
	}
}

func main() {
	go http.ListenAndServe("0.0.0.0:8080", nil)
	cpuhogger()
}
