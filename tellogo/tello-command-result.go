package tellogo

const (
	Ok    TelloCommandResult = "ok"
	Error TelloCommandResult = "error"
)

func ConvertTelloCommandResult(str string) TelloCommandResult {
	switch str {
	case "ok":
		return Ok
	case "error":
		return Error
	default:
		return Error
	}
}
