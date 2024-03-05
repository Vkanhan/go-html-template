package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var temp *template.Template

func main() {

	var err error

	temp, err = template.ParseGlob("*.html")
	if err != nil {
		log.Fatal("Error parsing template", err)
		os.Exit(1)
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/cart", cartHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/pay", payHandler)

	fmt.Println("Listening to port 8080..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := temp.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template", err)
	}
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	if err := temp.ExecuteTemplate(w, "cart.html", nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template", err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if err := temp.ExecuteTemplate(w, "login.html", nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template", err)
	}
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	if err := temp.ExecuteTemplate(w, "pay.html", nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template", err)
	}
}
