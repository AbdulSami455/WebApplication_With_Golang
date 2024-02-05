package main

/*
import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {

	//making of http client
	client := &http.Client{}

	//fmt.Print("hello")
	baseURL := "https://httpbin.org"
	// basic get Request to handle Request

	resp, err := client.Get(baseURL + "/get")
	if err != nil {
		fmt.Print("Error Happens")
	}
	defer resp.Body.Close()
	getBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(getBody))

	//Post request
	postBody := []byte(`{"key": "value"}`)
	postResp, err := client.Post(baseURL+"/post", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("Error on POST request:", err)
		return
	}
	defer postResp.Body.Close()
	postBodyResp, _ := io.ReadAll(postResp.Body)
	fmt.Println("POST Response:", string(postBodyResp))

	//There are other http Requests like POST,PUT and PATCH

}
*/
