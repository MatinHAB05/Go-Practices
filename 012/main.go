package main

import (
	"fmt"
	"net/http"
)

func IdentifyHandler(w http.ResponseWriter, r *http.Request) {
	a := r.Form
	b ,c , := r.FormFile("mamad")
	w.Write([]byte(fmt.Sprintf("%v\n", a)))
	w.
}

func main() {
	my_mux := http.NewServeMux()

	my_mux.HandleFunc("/", IdentifyHandler)
	my_mux.HandleFunc("/mamad/", IdentifyHandler)
	my_mux.HandleFunc("/home/", IdentifyHandler)
	my_mux.HandleFunc("/test/", IdentifyHandler)

	server := http.Server{
		Addr:    ":8097",
		Handler: my_mux,
	}

	fmt.Println("START!")
	fmt.Println(server.ListenAndServe())
	fmt.Println("END!")

}
