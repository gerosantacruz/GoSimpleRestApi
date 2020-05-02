package main

import "net/http"

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contact", contactPage)

	//start Server
	http.ListenAndServe(":3000", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact page"))
}
