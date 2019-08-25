package main

import (
	"net/http"

	"github.com/rcliao/go-starter-template/web"
)

func main() {
	http.HandleFunc("/hello", web.HelloHandler())
	http.ListenAndServe(":9000", nil)
}
