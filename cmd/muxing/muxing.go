package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", nameHandler).Methods(http.MethodGet)
	router.HandleFunc("/bad", internalServerErrorHandler).Methods(http.MethodGet)
	router.HandleFunc("/data", dataHandler).Methods(http.MethodPost)
	router.HandleFunc("/headers", headersHandler).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name, present := params["PARAM"]
	w.Header().Set("Content-Type", "text/plain")

	if !present {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`$PARAM not found`))
		return
	}
	// default code 200 OK will be set.
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(nil)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading param", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	w.Write([]byte(fmt.Sprintf("I got message:\n%s", buf)))

}

func headersHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		http.Error(w, "Header 'a' not set.", http.StatusBadRequest)
		return
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		http.Error(w, "Header 'b' not set.", http.StatusBadRequest)
		return
	}
	w.Header().Set("a+b", strconv.Itoa(a+b))
}
