package main 

import (
 "fmt"
 "net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
  
}

func loginsubmit(w http.ResponseWriter, r *http.Request) {

  
}
func handler(w http.ResponseWriter, r *http.Request) {
  switch r.URL.Path {
  case "/login":
    login(w,r)
   case "/loginsubmit":
    loginsubmit(w,r)
  default:
    fmt.Fprintf(w, "Hello You are Doing Wrong ")
    }

}
func main(){

  

  
}