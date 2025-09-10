package main

import (
	"fmt"
	"net/http"
)

const PORT = 3000

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Handling incoming users")
	})

	fmt.Println("Server is running on port:", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
