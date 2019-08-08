package tellogo

import (
	"fmt"
	"testing"
	"time"
)

var callbackChannel chan string

type callback struct{}

func (this *callback) OnCommandExecuted(command string, result string) {
	fmt.Printf("on command executed:%s result %s\n", command, result)
	callbackChannel <- fmt.Sprintf("%s%s", command, result)
}

func TestRunTello(t *testing.T) {
	callbackChannel = make(chan string)
	tello := NewTelloControllerType()
	callback := new(callback)
	err := tello.Start(callback)
	if err != nil {
		t.Fatalf("tello start failed %s", err.Error())
	}
	select {
	case res := <- callbackChannel:
		fmt.Printf("has response %s\n", res)
		if res != "commandok" {
			t.Fatalf("tello response is invalided %s", res)
		}
	case <- time.After(5 * time.Second):
		t.Fatal("response timeout")
	}
}
