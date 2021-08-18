package constant

//Error code
const (
	ServerError        = 5000
	MissingInformation = 40001
	InvalidInformation = 40002
	UserNotFound       = 41000
)

// Error message
const (
	MsgUserNotFound    = "user not found"
	MsgEmailDuplicated = "This email has been used"
	MsgInvalidPassword = "Invalid password"
	MsgServerError     = "Server error, please contact administrator"
)

var messages = map[int]string{
	MissingInformation: "Please input enough information.",
	ServerError:        MsgServerError,
	UserNotFound:       MsgUserNotFound,
}

func GetMessageFromCode(code int) (message string) {
	message = messages[code]
	return
}
