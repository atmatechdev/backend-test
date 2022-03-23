package api

type Response struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Error   error       `json:"error,omitempty"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(data interface{}, message string) Response {
	var status uint = 200
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return res
}

func ErrorResponse(status uint, message string) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
	return res
}

func TokenInvalidResponse() Response {
	res := Response{
		Status:  401,
		Message: "Please provide valid authorization token",
		Data:    nil,
	}
	return res
}
