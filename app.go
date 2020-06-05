package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", meta)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

const metaPage = `<html>
<meta name="go-import" content="go-get-issue-9357.appspot.com git https://github.com/rsc/go-get-issue-9357">
`

func meta(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(metaPage))
}
