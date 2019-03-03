package main

import (
	"fmt"
	"io"
	"net/http"
)

func Print1to20() int {
	res := 0
	for i := 1; i <= 20; i++ {
		res += i
	}
	return res
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1> helo,this is my first page! </h1>")
	io.WriteString(w, "<h1> hahahahaha <h1>")
}

func main() {
	fmt.Println("first page")
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil)
}
