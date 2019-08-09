package tellogo

// TelloCommand is command
type TelloCommand = string

// TelloCommandResult is command result
type TelloCommandResult = string

// TelloCommandCallbackType observes command result than tello receive
type TelloCommandCallbackType interface {
	OnCommandExecuted(command TelloCommand, result TelloCommandResult)
}

// TelloControllerType is interface of tello's controller
type TelloControllerType interface {
	Start(callback TelloCommandCallbackType) error
	End()
	SendCommand(command TelloCommand)
}

// TelloVideoCallbackType observes video frame what tello send
type TelloVideoCallbackType interface {
	OnUpdateVideoFrame(frame []byte)
}

// TelloVideoControllerType is interface of tello's video stream
type TelloVideoControllerType interface {
	Start(callback TelloVideoCallbackType) error
	End()
}

// Tello is endpoint
type Tello struct {
	Controller      TelloControllerType
	VideoController TelloVideoControllerType
}

// NewTello creat Tello instance
func NewTello() *Tello {
	controller := newTelloController()
	videoController := newTelloVideoController()
	tello := &Tello{Controller: controller, VideoController: videoController}
	// ビデオストリーム配信開始/終了は channel を使って自動的に controller に送る
	// コマンド送信忘れ防止
	go func() {
		for {
			command := <-videoController.requestCommandChannel
			controller.SendCommand(command)
		}
	}()
	return tello
}

// Version returns library version
func Version() string {
	return "0.0.1"
}
