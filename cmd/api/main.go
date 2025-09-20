package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	port := ":8081"
	http.Handle("/metrics", promhttp.Handler())
	fmt.Printf("listening on port : %s", port)
	http.ListenAndServe(port, nil)
}
