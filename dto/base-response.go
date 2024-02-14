package dto

type BaseResponse struct {
	Data         interface{} `json:"data"`
	Success      bool        `json:"success"`
	MessageTitle string      `json:"messageTitle"`
	Message      string      `json:"message"`
}

func FailedResponse(msg string, responseCode int) (BaseResponse, int) {
	return BaseResponse{
		Data:         nil,
		Success:      false,
		MessageTitle: "error",
		Message:      msg,
	}, responseCode
}

func SuccessResponse(data any, msg string, responseCode int) (BaseResponse, int) {
	return BaseResponse{
		Data:         data,
		Success:      true,
		MessageTitle: "success",
		Message:      msg,
	}, responseCode
}
