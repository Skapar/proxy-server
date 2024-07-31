package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Skapar/proxy-server/internal/proxy"

	_ "github.com/Skapar/proxy-server/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Proxy Server API
// @version 1.0
// @description This is a proxy server API.
// @host localhost:8080
// @BasePath /
func main() {
	http.HandleFunc("/proxy", proxy.ProxyHandler)
	http.HandleFunc("/health", proxy.HealthCheckHandler)

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Proxy server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}