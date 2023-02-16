package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello I am the main function bitches!!")
	greater()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	//running a server in golang
	// in case of servers we usually logs crashes
	// we need to build as well using go build .
	// and run using go run main.go
	// tidy tool is uses for fixing mod file

	// go mod verify is used for verifying the hashes
	// go list is used for listing all the dependedncies
	// go list all
	// go list -m all
	// go list -m -versions github.com/gorilla/mux this will give us all the availabel versions
	// go mod why github.com/gorilla/mux will tell why we are dependent on this dependency
	// go mod graph will tell us the entire grph of dependency
	// go mod edit is useful for editing the go mod file
	// -go for changing go version
	// -module for module editing
	// production command go mod vendor will give us something
	// similar to node modules
	// go run -mod=vendor main.go to refer from vendor folder
	// instead of cache
	log.Fatal(http.ListenAndServe(":4000", r))
}

// in mod init we add the name of destination wher we are going to host our code
// we are now moving from workspace to modules in golang after 2019
// got get is used for installing the modules

func greater() {
	fmt.Println("Dhatt teri mayi ka chodo!!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// response writer is used for sendin back a response
	w.Write([]byte("<h1>Teri maa ka bhosda on desi dehati sex vidos</h1>"))
}
