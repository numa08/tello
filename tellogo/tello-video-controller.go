package tellogo

import (
	"fmt"
	"net"
)

const localVideoAddress = ":11111"

type TelloVideoController struct {
	conn                  *net.UDPConn
	callback              TelloVideoCallbackType
	receiveChannel        chan struct{}
	requestCommandChannel chan TelloCommand
}

func newTelloVideoController() *TelloVideoController {
	return &TelloVideoController{requestCommandChannel: make(chan TelloCommand)}
}

func (this *TelloVideoController) Start(callback TelloVideoCallbackType) error {
	laddr, err := net.ResolveUDPAddr("udp", localVideoAddress)
	if err != nil {
		return err
	}
	this.callback = callback
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return err
	}
	this.conn = conn
	this.receiveChannel = make(chan struct{})
	go this.receive()
	this.requestCommandChannel <- StreamOn
	return nil
}

func (this *TelloVideoController) End() {
	this.requestCommandChannel <- StreamOff
	if this.receiveChannel != nil {
		close(this.receiveChannel)
	}
}

func (this *TelloVideoController) send(frame []byte) {
	if this.callback == nil {
		return
	}
	this.callback.OnUpdateVideoFrame(frame)
}

func (this *TelloVideoController) receive() {
	defer func() {
		if this.conn != nil {
			this.conn.Close()
		}
	}()
	for {
		buf := make([]byte, 2048)
		n, _, err := this.conn.ReadFrom(buf)
		if err != nil {
			fmt.Printf("error %s \n", err.Error())
		}
		go this.send(buf[0:n])
		select {
		case <-this.receiveChannel:
			return
		default:
		}
	}
}
