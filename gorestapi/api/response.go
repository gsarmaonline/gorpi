package api

type (
	Response struct {
		Request *Request
	}
)

func NewResponse(request *Request) (response *Response) {
	response = &Response{
		Request: request,
	}
	return
}
