package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello wold")
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
