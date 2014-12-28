package main

import (
    "net/http"
    "log"
    "fmt"
    )

var _host string
var _port = "12000"

func main() {
    http.HandleFunc("/", homePage)
    listenOn := fmt.Sprintf("%s:%s", _host, _port)
    fmt.Printf("Statistics http server is listening on:\n\t%s\n", listenOn);
    if err := http.ListenAndServe(listenOn, nil); err != nil {
        log.Fatal("failed to start server", err)
    }
}
