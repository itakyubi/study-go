package main

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	// ffmpeg -re -i demo.flv -c copy -f flv rtmp://localhost:1935/test/rfBd56ti2SMtYvSgD5xAV0YU99zampta7Z7S575KLkIZ9PYk

	_ = ffmpeg.Input("/Users/wuao/Desktop/demo.mp4").
		HFlip().DrawBox(100, 50, 200, 200, "red", 5).
		Output("/Users/wuao/Desktop/demo.avi").
		OverWriteOutput().ErrorToStdOut().Run()

}
