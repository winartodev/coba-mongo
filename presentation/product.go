package presentation

import (
	"encoding/json"
	"net/http"
	"winartodev/coba-mongodb/entity"
	"winartodev/coba-mongodb/response"
	"winartodev/coba-mongodb/usecase"

	"github.com/julienschmidt/httprouter"
)

type ProductPresentation interface {
	FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

type productPresentation struct {
	usecase usecase.ProductUsecase
}

func NewProductPresentation(usecase usecase.ProductUsecase) ProductPresentation {
	return &productPresentation{usecase: usecase}
}

func (p *productPresentation) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := p.usecase.FindAll(r.Context())
	if err != nil {
		response.HttpResponseFailed(w, r, http.StatusInternalServerError, err)
		return
	}

	response.HttpResponseSuccess(w, r, res)
	return
}

func (p *productPresentation) Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	product := entity.Product{}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		response.HttpResponseFailed(w, r, http.StatusBadRequest, err)
		return
	}

	err := p.usecase.Insert(r.Context(), product)
	if err != nil {
		response.HttpResponseFailed(w, r, http.StatusInternalServerError, err)
		return
	}

	response.HttpResponseSuccess(w, r, product)
	return
}
