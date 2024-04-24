package main

import (
    "fmt"
    "log"
    "net/http"
)


func main() {

    var tls_config = TSLConfig {
        Cert: "",
        Key:  "",
        Port: ":443",
    }

    fmt.Println("Listening on port " + tls_config.Port + "...");
    err := http.ListenAndServeTLS(":443", tls_config.Cert, tls_config.Key, nil); 
    if err != nil {
        log.Println(err.Error())
    }
    
}
