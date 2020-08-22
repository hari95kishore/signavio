package main

import (
    "net/http"
    "io/ioutil"
    "log"
    "encoding/json"
    "strings"
)


func main() {
    response, err := http.Get("http://signavio-server.com")
    if err != nil {
        log.Fatal(err)
    }

    respbody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    message := make(map[string]string)
    err = json.Unmarshal(respbody, &message)
    if err != nil {
        log.Fatal(err)
    }

    reverse_msg := strings.Fields(message["message"])
    for i, j := 0, len(reverse_msg)-1; i < j; i, j = i+1, j-1 {
        reverse_msg[i], reverse_msg[j] = reverse_msg[j], reverse_msg[i]
    }
    log.Print(strings.Join(reverse_msg, " "))
}
