package api

type BaseRequest struct {
	Data    string
	KeyType KeyType
	Action  Action // get, put, delete, list ...
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	CodeSuccess = 200
	CodeFail    = 500 + iota
	CodeActionNotAllowed
)

const (
	MsgSuccess = "success"
)

func SuccessResp(data interface{}) BaseResponse {
	return BaseResponse{
		Code:    CodeSuccess,
		Message: MsgSuccess,
		Data:    data,
	}
}

func ErrResp(code int, msg string) BaseResponse {
	return BaseResponse{
		Code:    code,
		Message: msg,
	}
}
