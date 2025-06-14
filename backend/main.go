package main

import (
	"fmt"
	"net/http"
	"backend/router"
)

func main() {
	r := router.SetupRouter()
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", r)
}