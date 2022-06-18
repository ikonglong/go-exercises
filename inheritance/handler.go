package inheritance

import "fmt"

type handler interface {
	translate(httpReq string) string
	validate(req string) error
	handle(req string) error
}

type handlerWrapper struct {
	handler
}

func (h *handlerWrapper) handle(req string) error {
	h.handler.translate(req)
	h.handler.validate(req)
	h.handler.handle(req)
	return nil
}

type handlerA struct {
}

func (h handlerA) translate(req string) string {
	fmt.Println("translate")
	return req
}

func (h handlerA) validate(req string) error {
	fmt.Println("validate")
	return nil
}

func (h handlerA) handle(req string) error {
	fmt.Println("handle")
	return nil
}
