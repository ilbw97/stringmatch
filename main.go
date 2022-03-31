package main

import (
	"fmt"
	"net/http"
	"stringmatch/ahocorasick"
	"stringmatch/kmp"
)

func main() {
	// pprof
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6666", nil))
	}()
	kmp.Dorun()
	ahocorasick.Dotest()

}
