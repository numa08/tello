package tellogo

type TelloCommand string

type TelloCommandResult string

type TelloCommandCallbackType interface {
	OnCommandExecuted(command string, result string)
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

func NewTelloControllerType() TelloControllerType {
	return new(TelloController)
}

func NewTelloVideoControllerType() TelloVideoControllerType {
	return new(TelloVideoController)
}
