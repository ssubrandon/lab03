package main

import (
	"net/http"
	"os"
)

func main() {
	http:HandleFunc("/", returnContainerHandler)
	if err := http:ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

func returnContainerHandler(w http.ResonseWriter, r *http.Request){
	message, _ := os.Hostname()
	w.Write([]byte(message))
}
