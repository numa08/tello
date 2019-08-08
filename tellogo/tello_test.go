package tellogo

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var callbackChannel chan string

type controllerCallback struct{}

func (this *controllerCallback) OnCommandExecuted(command string, result string) {
	fmt.Printf("on command executed:%s result %s\n", command, result)
	callbackChannel <- fmt.Sprintf("%s%s", command, result)
}

var videoBuffer []byte

type videoControllerCallback struct{}

func (this *videoControllerCallback) OnUpdateVideoFrame(frame []byte) {
	fmt.Printf("on video frame updated\n")
	videoBuffer = append(videoBuffer, frame...)
}

func TestVideo(t *testing.T) {
	videoBuffer = []byte{}
	tello := NewTello()
	controller := tello.Controller
	video := tello.VideoController
	videoCallback := new(videoControllerCallback)
	err := controller.Start(nil)
	if err != nil {
		t.Fatalf("tello start failed %s\n", err.Error())
		return
	}
	video.Start(videoCallback)
	select {
	case <-time.After(15 * time.Second):
	}
	if len(videoBuffer) <= 0 {
		t.Fatal("video frame is empty")
	} else {
		videoFile, _ := os.Create("video.raw")
		videoFile.Write(videoBuffer)
	}
	video.End()
	controller.End()
}

func TestRunTello(t *testing.T) {
	callbackChannel = make(chan string)
	tello := NewTello()
	controller := tello.Controller
	callback := new(controllerCallback)
	err := controller.Start(callback)
	if err != nil {
		t.Fatalf("tello start failed %s", err.Error())
	}
	select {
	case res := <-callbackChannel:
		fmt.Printf("has response %s\n", res)
		if res != "commandok" {
			t.Fatalf("tello response is invalided %s", res)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("response timeout")
	}
	controller.End()
}
