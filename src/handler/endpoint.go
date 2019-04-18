package handler

type CoreHandler {
	service *service.Service
}

func NewCoreHandler( service *service.Service) {
	return CoreHandler{

	}
}

func (c *CoreHandler) GetUser(rw, resp) {
	req => input
	out c.GetUser(input)
	out => resp
}