package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT string = "3000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	fmt.Println("Server listening on port", PORT)

	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
