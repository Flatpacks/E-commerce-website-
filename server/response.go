// wrap json reponse
package server

const (
	CodeInvalidDataError  = 40002
	CodeParseRequestError = 40201
	CodeDataNotFoundError = 40401
	CodeInternalErr       = 50001
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Code:    0,
		Data:    data,
		Message: "success",
	}
}

func NewE