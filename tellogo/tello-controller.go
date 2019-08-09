package tellogo

import (
	"fmt"
	"net"
)

const localControllerAddress = ":9000"
const telloControllerAddress = "192.168.10.1:8889"

type telloController struct {
	receiveChannel chan struct{}
	commandChannel chan TelloCommand
	conn           *net.UDPConn
	callback       TelloCommandCallbackType
}

func newTelloController() *telloController {
	return new(telloController)
}

func (tc *telloController) Start(callback TelloCommandCallbackType) error {
	addr, err := net.ResolveUDPAddr("udp", telloControllerAddress)
	if err != nil {
		return err
	}
	laddr, err := net.ResolveUDPAddr("udp", localControllerAddress)
	if err != nil {
		return err
	}
	tc.callback = callback
	conn, err := net.DialUDP("udp", laddr, addr)
	tc.conn = conn
	tc.receiveChannel = make(chan struct{})
	tc.commandChannel = make(chan TelloCommand)
	go tc.receive()
	go tc.SendCommand(Command)
	return nil
}

func (tc *telloController) End() {
	if tc.receiveChannel != nil {
		close(tc.receiveChannel)
	}
}

func (tc *telloController) SendCommand(command TelloCommand) {
	if tc.conn == nil {
		return
	}
	tc.commandChannel <- command
	_, err := tc.conn.Write([]byte(command))
	if err != nil {
		fmt.Printf("error %s \n", err.Error())
		go tc.End()
		return
	}
}

func (tc *telloController) receive() {
	defer func() {
		if tc.conn != nil {
			tc.conn.Close()
		}
	}()
	for {
		command := <-tc.commandChannel
		if tc.conn == nil {
			go tc.End()
			return
		}

		buf := make([]byte, 1024)
		n, err := tc.conn.Read(buf[0:])
		if err != nil {
			fmt.Printf("error %s \n", err.Error())
			go tc.End()
			return
		}
		result := string(buf[0:n])
		if tc.callback != nil {
			tc.callback.OnCommandExecuted(command, ConvertTelloCommandResult(result))
		}
		select {
		case <-tc.receiveChannel:
			return
		default:
		}
	}
}
