package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot)

	fmt.Println("listening to port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error while listening to port 8080")
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving / endpoint")
	io.WriteString(w, "hello users !!!")
}
