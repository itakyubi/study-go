package main

/*
#cgo LDFLAGS: -L . -lapi -lstdc++
#cgo CFLAGS: -I ./
#include "api.h"
*/

/*
#cgo LDFLAGS: -L . -lvideo_stream_api -lstdc++
#cgo CFLAGS: -I ./lib
#include "video_stream_api.h"
*/
import "C"
import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	C.init()

	for {
		lastHandleTime := time.Now().UnixNano() / 1e6
		imageDir, err := ioutil.ReadDir("/home/work/app/data/images/")
		if err != nil {
			println("read image dir error:" + err.Error())
			return
		}

		for _, imageFile := range imageDir {
			if !strings.HasSuffix(imageFile.Name(), ".jpg") {
				continue
			}

			sleepTime := 100 - (time.Now().UnixNano()/1e6 - lastHandleTime)
			if sleepTime > 0 {
				time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			}

			output := Output{
				StreamKey: "0",
				Image:     "/home/work/app/data/images/" + imageFile.Name(),
			}
			msg, _ := json.Marshal(output)
			C.receive(C.CString(string(msg)))

			lastHandleTime = time.Now().UnixNano() / 1e6
		}
	}
}

type Output struct {
	StreamKey string `json:"streamKey"`
	Image     string `json:"image"`
}
