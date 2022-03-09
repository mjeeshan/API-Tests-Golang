package main

import (
	"fmt"
	"net/http"

	"github.com/jeeshan12/apimongo/router"
)

func main() {

	fmt.Println("Server is starting ...")
	r := router.Router()
	http.ListenAndServe(":4000", r)
	fmt.Println("Server is running ...")
}
