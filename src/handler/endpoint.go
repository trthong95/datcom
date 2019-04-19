package handler

import (
	"git.d.foundation/datcom/backend/src/service"
)

type CoreHandler struct {
	service *service.Service
}

func NewCoreHandler(service *service.Service) *CoreHandler {
	return &CoreHandler{
		service: service,
	}
}

// func (c *CoreHandler) GetUser(rw, resp) {
// 	req => input
// 	out c.GetUser(input)
// 	out => resp
// }
