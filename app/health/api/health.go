package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/yoas0bi/micro-toolkit/app/health"
	"github.com/yoas0bi/micro-toolkit/http/restful/response"
)

func (h *handler) Check(r *restful.Request, w *restful.Response) {
	req := health.NewHealthCheckRequest()
	resp, err := h.service.Check(
		r.Request.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, NewHealth(resp))
}
