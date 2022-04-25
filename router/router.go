package router

import "errors"

const GIN string = "gin"
const MUX string = "mux"

type (
	Router interface {
		Get(path string, handleFunc interface{})
		Post(path string, handleFunc interface{})
		Put(path string, handleFunc interface{})
		Patch(path string, handleFunc interface{})
		Delete(path string, handleFunc interface{})
		Options(path string, handleFunc interface{})
		Run() error
		UseMiddleware(middleware interface{})
	}
)

func New(routerType string) (Router, error) {
	switch routerType {
	case GIN:
		return &RouterGin{}, nil
	case MUX:
		return &RouterMux{}, nil
	}

	return nil, errors.New("invalid router type")
}
