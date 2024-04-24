package main 

import (
    "crypto"
   	"database/sql"
    "encoding/json"
    "net/http"
    "log"
)

func authenticate_request(r *http.Request) (User, error) {
    var user User;


    user_cookie, err := r.Cookie("ibcm-auth");
    if err != nil {
        return user, err;
    }

    err = json.Unmarshal([]byte(user_cookie.Value), user);
    if err != nil {
        return user, err;
    }

}


