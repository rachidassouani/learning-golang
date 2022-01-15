package main

import (
	"fmt"
	"net/http"

	"github.com/rachidassouani/go-project/pkg/hanldlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", hanldlers.Home)
	http.HandleFunc("/about", hanldlers.About)

	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
