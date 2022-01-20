package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello MDP")

	//http.HandleFunc("/", homePageHandler)

	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

/*func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello MDP")
	checkError(err)
}*/

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
