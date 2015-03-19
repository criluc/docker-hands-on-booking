package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/booking/{user}", Booking)

    log.Fatal(http.ListenAndServe(":18273", router))
}

func Booking(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    user := vars["user"]
    fmt.Fprintf(w, "\nCiao %s, thanks for your registration!\nSee you at the hands-on.\n\n", user)
    log.Printf("IP: %s, User: %s", r.RemoteAddr, user);
}

