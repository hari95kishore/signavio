package main

import (
    "net/http"
    "log"
    "encoding/json"
)


type Greeting struct {
    Id  string `json:"id"`
    Message string `json:"message"`
}

func getGreeting(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        log.Print("Method not allowed")
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    log.Print("received GET request from ", r.Host)
    greet := Greeting{
        Id: "1",
        Message: "Hello world",
    }
    responsejson, err := json.Marshal(greet)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
        return
    }

    w.Header().Add("content-type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responsejson)
}

func main() {
    http.HandleFunc("/", getGreeting)
    http.ListenAndServe(":8080", nil)
}
