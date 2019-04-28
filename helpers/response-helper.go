package helpers

import (
	"github.com/go-chi/render"
	m "dragon-fruit-expert-system/models/response"
	"net/http"
)

func ResWithoutMeta(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	res := m.ResWithoutMeta{Code: code, Status: http.StatusText(code), Data:data}

	render.Status(r, code)
	render.JSON(w, r, res)
}

func ResWithMeta(w http.ResponseWriter, r *http.Request, code int, limit uint, page uint, count uint, data interface{}) {
	countPage, remainder := count/limit, count%limit
	if remainder > 0 {
		countPage = countPage + 1
	}

	meta := m.Meta{
		Limit: limit,
		Page: page,
		CountData: count,
		CountPage: countPage,
	}

	res := m.ResWithMeta{Code: code, Status: http.StatusText(code), Meta: meta, Data: data}

	render.Status(r, code)
	render.JSON(w, r, res)
}

func ResError(w http.ResponseWriter, r *http.Request, code int, msg interface{}) {
	res := m.ResError{Code: code, Message: msg, Status: http.StatusText(code)}

	render.Status(r, code)
	render.JSON(w, r, res)
}

func ResCustom(w http.ResponseWriter, r *http.Request, code int, res interface{}) {
	render.Status(r, code)
	render.JSON(w, r, res)
}