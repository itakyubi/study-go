package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

var channel chan string

func main() {
	channel = make(chan string)
	go func() {
		for {
			s := <-channel
			println("get from chan:" + s)
		}

	}()

	http.HandleFunc("/msg", Handler)
	http.ListenAndServe(":8000", nil)

}

func Handler(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	println("body:" + string(body))

	channel <- string(body)

	response := Response{Code: 200}
	responseByte, _ := json.Marshal(response)
	io.WriteString(w, string(responseByte))
}

type Response struct {
	Code int `json:"code,omitempty"`
}
