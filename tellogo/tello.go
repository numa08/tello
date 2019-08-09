package tellogo

type TelloCommand = string

type TelloCommandResult = string

type TelloCommandCallbackType interface {
	OnCommandExecuted(command TelloCommand, result TelloCommandResult)
}

type TelloControllerType interface {
	Start(callback TelloCommandCallbackType) error
	End()
	SendCommand(command TelloCommand)
}

type TelloVideoCallbackType interface {
	OnUpdateVideoFrame(frame []byte)
}

type TelloVideoControllerType interface {
	Start(callback TelloVideoCallbackType) error
	End()
}

type Tello struct {
	Controller      TelloControllerType
	VideoController TelloVideoControllerType
}

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
