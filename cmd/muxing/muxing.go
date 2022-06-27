package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
	_ "github.com/stretchr/testify"
)

const getNameParameterKey = "param"

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := httprouter.New()

	router.GET("/name/:"+getNameParameterKey, nameHandler)
	router.GET("/bad", badHandler)
	router.POST("/data", dataHandler)
	router.POST("/headers", headersHandler)

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

func nameHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", ps.ByName(getNameParameterKey))
}

func badHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusInternalServerError)
}

func dataHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I got message:\n"+string(data))
}

func headersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("a+b", strconv.Itoa(a+b))
	w.WriteHeader(http.StatusOK)
}
