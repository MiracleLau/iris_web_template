package model

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) Response {
	return Response{
		Success: true,
		Message: "操作成功",
		Data:    data,
	}
}

func Error(msg string) Response {
	return Response{
		Success: false,
		Message: msg,
	}
}