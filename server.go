package main

// Import Packages and Libraries 
import (
	"fmt"
	"net/http"
	"html/template"
)

//Func for HTML template Rendering
func login(w http.ResponseWriter, r *http.Request) {
  var filename ="login1.html"
  t, err :=template.ParseFiles(filename)
  if err != nil {
    fmt.Println("Error in parsing the template")
    }
  t.Execute(w,filename)
  
}

//func loginsubmit(w http.ResponseWriter, r *http.Request) {

  
//}
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

  http.HandleFunc("/", handler)
  http.ListenAndServe("",nil)

  
}
