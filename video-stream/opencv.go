package main

import (
	"fmt"
	"github.com/hybridgroup/mjpeg"
	"gocv.io/x/gocv"
	"log"
	"net/http"
)

/*
	使用opencv+http的方式实现视频流，本质上不是真正的视频流，而是一张张图片播放达到近似于视频流的效果，最终播放效果存在卡顿，暂时作为兜底方案。
	原理：
	1、opencv读取视频源
	2、帧图片转换成Mat类型
	3、Mat类型编码成jpg的二进制
	4、持续推送jpg的二进制给http
	ps. 如果是基于上游图片的，则前三步可以合并成jpg文件转jpg二进制数据
*/

var (
	deviceID int
	err      error
	webcam   *gocv.VideoCapture
	stream   *mjpeg.Stream
)

func main() {
	// open webcam
	// webcam, err = gocv.OpenVideoCapture("rtsp://wowzaec2demo.streamlock.net/vod/mp4:BigBuckBunny_115k.mov")
	// webcam, err = gocv.OpenVideoCapture("/Users/wuao/Desktop/dcell/demo.mp4")
	webcam, err = gocv.OpenVideoCapture("rtsp://172.24.217.156:554/live/rtsp")
	if err != nil {
		fmt.Printf("Error opening capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// create the mjpeg stream
	stream = mjpeg.NewStream()

	// start capturing
	go mjpegCapture()

	// start http server
	http.Handle("/", stream)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func mjpegCapture() {
	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		buf, _ := gocv.IMEncode(".jpg", img)
		stream.UpdateJPEG(buf)
	}
}
