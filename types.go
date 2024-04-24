package main


type User struct {
    Username string    `json:"username"`
    Password string    `json:"password"`
}

type TSLConfig struct {
    Cert string 
    Key  string
    Port string
}
