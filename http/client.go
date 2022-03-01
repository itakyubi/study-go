package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		body := Request{
			Timestamp: time.Now().UnixNano() / 1e6,
			Image:     "/Users/wuao/Desktop",
		}
		bodyByte, _ := json.Marshal(body)
		request, _ := http.NewRequest("POST", "http://127.0.0.1:8000/msg", bytes.NewBuffer(bodyByte))

		client := http.Client{}
		resp, _ := client.Do(request)

		if resp != nil {
			defer resp.Body.Close()
		}

		result, _ := ioutil.ReadAll(resp.Body)
		println(string(result))
	}

}

type Request struct {
	Timestamp int64  `json:"Timestamp,omitempty"`
	Image     string `json:"Image,omitempty"`
}
