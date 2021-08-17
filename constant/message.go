package constant

const (
	ServerError = 5000

	MissingInformation = 40001
)

var messages = map[int]string{
	MissingInformation: "Please input enough information.",
	ServerError:        "Server error, please contact administrator",
}

func GetMessageFromCode(code int) (message string) {
	message = messages[code]
	return
}
