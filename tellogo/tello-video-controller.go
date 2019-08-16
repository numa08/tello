package tellogo

import (
	"fmt"
	"net"
)

const localVideoAddress = ":11111"

type telloVideoController struct {
	conn                  *net.UDPConn
	callback              TelloVideoCallbackType
	receiveChannel        chan struct{}
	requestCommandChannel chan TelloCommand
}

func newTelloVideoController() *telloVideoController {
	return &telloVideoController{requestCommandChannel: make(chan TelloCommand)}
}

func (tv *telloVideoController) Start(callback TelloVideoCallbackType) error {
	laddr, err := net.ResolveUDPAddr("udp", localVideoAddress)
	if err != nil {
		return err
	}
	tv.callback = callback
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return err
	}
	tv.conn = conn
	tv.receiveChannel = make(chan struct{})
	go tv.receive()
	tv.requestCommandChannel <- StreamOn
	return nil
}

func (tv *telloVideoController) End() {
	tv.requestCommandChannel <- StreamOff
	if tv.receiveChannel != nil {
		close(tv.receiveChannel)
	}
}

func (tv *telloVideoController) send(frame []byte) {
	if tv.callback == nil {
		return
	}
	tv.callback.OnUpdateVideoFrame(frame)
}

func (tv *telloVideoController) receive() {
	defer func() {
		if tv.conn != nil {
			tv.conn.Close()
		}
	}()
	var frame = []byte{}
	for {
		buf := make([]byte, 2048)
		n, _, err := tv.conn.ReadFrom(buf)
		frame = append(frame, buf[0:n]...)
		if err != nil {
			fmt.Printf("error %s \n", err.Error())
		}
		if n != 1460 {
			// did receive 1 frame
			go tv.send(frame)
			frame = []byte{}
		}

		select {
		case <-tv.receiveChannel:
			return
		default:
		}
	}
}
