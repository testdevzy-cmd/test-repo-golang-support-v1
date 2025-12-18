package main

type Request struct {
	Method string
	Path   string
	Body   interface{}
}

type Response struct {
	Status int
	Data   interface{}
}

type UserHandler struct {
	Prefix string
}

func (h *UserHandler) Handle(req *Request) Response {
	return Response{Status: 200, Data: "ok"}
}
