package main

import (
  "fmt"
  "net/http"
)

func helloworldapp(w http.ResponseWriter, r *http.Request) {
  switch r.URL.Path{
    case "/":
      fmt.Fprint(w, "Hello, world")
    case "/hello":
    fmt.Fprint(w, "Hello sami")
  }
  fmt.Printf("Handling Request with %s Request: ", r.Method)
}

func main() {
  http.HandleFunc("/", helloworldapp)

  server := http.Server{
    Addr:"",
    Handler: nil,
    ReadTimeout: 1000,
    WriteTimeout: 1000,  
  }
  server.ListenAndServe()
}
