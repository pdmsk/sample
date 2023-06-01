package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Simple struct {
	Name        string
	Description string
	Url         string
}

type Factory struct {
	Url string
}

func SimpleFactory(url string) *Factory {
	return &Factory{Url: url}
}

func (f *Factory) GetSimple() *Simple {
	return &Simple{
		Name:        "Artsiom",
		Description: "World",
		Url:         f.Url,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	factory := SimpleFactory(r.Host)
	simple := factory.GetSimple()

	jsonOutput, _ := json.Marshal(simple)

	fmt.Fprintln(w, string(jsonOutput))
}

func main() {
	fmt.Println("Server started on port 4444")
	http.HandleFunc("/api", handler)
	log.Fatal(http.ListenAndServe(":4444", nil))
}
