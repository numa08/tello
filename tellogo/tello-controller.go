package tellogo

import (
	"fmt"
	"net"
)

const localControllerAddress = ":9000"
const telloControllerAddress = "192.168.10.1:8889"

type TelloController struct {
	receiveChannel chan struct{}
	commandChannel chan TelloCommand
	conn           *net.UDPConn
	callback       TelloCommandCallbackType
}

func newTelloController() *TelloController {
	return new(TelloController)
}

func (this *TelloController) Start(callback TelloCommandCallbackType) error {
	addr, err := net.ResolveUDPAddr("udp", telloControllerAddress)
	if err != nil {
		return err
	}
	laddr, err := net.ResolveUDPAddr("udp", localControllerAddress)
	if err != nil {
		return err
	}
	this.callback = callback
	conn, err := net.DialUDP("udp", laddr, addr)
	this.conn = conn
	this.receiveChannel = make(chan struct{})
	this.commandChannel = make(chan TelloCommand)
	go this.receive()
	go this.SendCommand(Command)
	return nil
}

func (this *TelloController) End() {
	if this.receiveChannel != nil {
		close(this.receiveChannel)
	}
}

func (this *TelloController) SendCommand(command TelloCommand) {
	if this.conn == nil {
		return
	}
	this.commandChannel <- command
	_, err := this.conn.Write([]byte(command))
	if err != nil {
		fmt.Printf("error %s \n", err.Error())
		go this.End()
		return
	}
}

func (this *TelloController) receive() {
	defer func() {
		if this.conn != nil {
			this.conn.Close()
		}
	}()
	for {
		command := <-this.commandChannel
		if this.conn == nil {
			go this.End()
			return
		}

		buf := make([]byte, 1024)
		n, err := this.conn.Read(buf[0:])
		if err != nil {
			fmt.Printf("error %s \n", err.Error())
			go this.End()
			return
		}
		result := string(buf[0:n])
		if this.callback != nil {
			this.callback.OnCommandExecuted(string(command), string(ConvertTelloCommandResult(result)))
		}
		select {
		case <-this.receiveChannel:
			return
		default:
		}
	}
}
