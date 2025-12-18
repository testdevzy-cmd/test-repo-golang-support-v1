package main

func SetupRoutes() {
	handler := &UserHandler{Prefix: "api"}
	req := &Request{Method: "GET", Path: "/users"}
	resp, err := handler.Handle(req)
	if err != nil {
		return
	}
	_ = resp
}
