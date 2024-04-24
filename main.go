package main

import (
    "fmt"
    "log"
    "net/http"
)


func main() {

    var tls_config = TSLConfig {
        cert: "",
        key:  "",
        port: ":443",
    }

    fmt.Println("Listening on port " + tls_config.port + "...");
    err := http.ListenAndServeTLS(":443", tls_config.cert, tls_config.key, nil); 
    if err != nil {
        log.Println(err.Error())
    }
    
}
