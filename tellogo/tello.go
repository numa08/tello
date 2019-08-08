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

func NewTelloControllerType() TelloControllerType {
	return new(TelloController)
}
