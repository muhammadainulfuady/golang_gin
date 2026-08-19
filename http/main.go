package main

import (
	"fmt"
	"net/http"
)

func AuthorizationHandler(next http.Handler) {
	fmt.Println("Monggo lewat cong")
	next.ServeHTTP(nil, nil)
}

func masuk(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Halo cong sudah lewat")
}

func main() {
	AuthorizationHandler(http.HandlerFunc(masuk))
}
