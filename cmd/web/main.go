package main

import (
	"fmt"
	"net/http"
	"github.com/alpapie/go_web/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/contact",handlers.Contact )

	fmt.Println("le site est ouvert sur: http://127.0.0.1:8080  --avec le port", port)
	http.ListenAndServe(port, nil)

}
