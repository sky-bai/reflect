package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w,"%s",r.URL)
}
func main() {
	http.HandleFunc("/golang", index)
	http.ListenAndServe("0.0.0.0:8080",nil)
}

