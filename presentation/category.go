package presentation

import (
	"encoding/json"
	"net/http"
	"winartodev/coba-mongodb/entity"
	"winartodev/coba-mongodb/response"
	"winartodev/coba-mongodb/usecase"

	"github.com/julienschmidt/httprouter"
)

type Category interface {
	FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	FindOne(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type category struct {
	uc usecase.CategoryUsecase
}

func NewCategoryPresentation(categoryUsecase usecase.CategoryUsecase) category {
	return category{uc: categoryUsecase}
}

func (c *category) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := c.uc.FindAll(r.Context())
	if err != nil {
		response.HttpResponseFailed(w, r, http.StatusInternalServerError, err)
		return
	}

	response.HttpResponseSuccess(w, r, res)
	return
}

func (c *category) FindOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	slug := p.ByName("slug")
	res, err := c.uc.FindOne(r.Context(), slug)
	if err != nil {
		response.HttpResponseFailed(w, r, http.StatusInternalServerError, err)
		return
	}

	response.HttpResponseSuccess(w, r, res)
	return
}

func (c *category) Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	category := entity.Category{}
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		response.HttpResponseFailed(w, r, http.StatusBadRequest, err)
		return
	}

	err := c.uc.Insert(r.Context(), category)
	if err != nil {
		response.HttpResponseFailed(w, r, http.StatusInternalServerError, err)
		return
	}

	response.HttpResponseSuccess(w, r, "data success created")
}

func (c *category) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	slug := p.ByName("slug")

	category := entity.Category{}
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		response.HttpResponseFailed(w, r, http.StatusBadRequest, err)
		return
	}

	err := c.uc.Update(r.Context(), slug, category)
	if err != nil {
		response.HttpResponseFailed(w, r, http.StatusInternalServerError, err)
		return
	}

	response.HttpResponseSuccess(w, r, "data success updated")
}

func (c *category) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	slug := p.ByName("slug")

	err := c.uc.Delete(r.Context(), slug)
	if err != nil {
		response.HttpResponseFailed(w, r, http.StatusInternalServerError, err)
		return
	}

	response.HttpResponseSuccess(w, r, "data success deleted")
	return
}
