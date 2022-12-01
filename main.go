package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	// Serve Home
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileserver := template.Must(template.ParseFiles(os.Getenv("KO_DATA_PATH") + "/index.html"))
		fileserver.Execute(w, true)
	})

	// Serve about
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fileserver := template.Must(template.ParseFiles(os.Getenv("KO_DATA_PATH") + "/about.html"))
		fileserver.Execute(w, true)
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	log.Fatal(http.ListenAndServe(port, nil))

}
