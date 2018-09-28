package main

import (
        "net/http"
        "os"
)

func main() {
        http.Handlefunc("/", return ContainerHandler)
        if err := http.ListenAndServe(":80", nil); err != nil {
                panic(err)
        }
}

func returnContainerHandler(w http.ResponseWriter, r *http.Request){
        message, _ := os.Hostname()
        w.Write([]byte(message))
}
