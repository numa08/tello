package tellogo

const (
	// Ok is tello command is Ok
	Ok TelloCommandResult = "ok"
	// Error is tello command is error
	Error TelloCommandResult = "error"
)

// ConvertTelloCommandResult returns tello command result as ok/erro.
// str is althought ok or error, nether that when return error
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
