package products

import (
	"lectronic/src/databases/orm/models"
	"lectronic/src/interfaces"
	"lectronic/src/libs"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type ProductController struct {
	srvc interfaces.ProductSrvc
}

func NewController(srvc interfaces.ProductSrvc) *ProductController {
	return &ProductController{srvc}
}

func (c *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	c.srvc.GetAll().Send(w)
}

func (c *ProductController) GetByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.GetByID(id).Send(w)
}

func (c *ProductController) Popular(w http.ResponseWriter, r *http.Request) {

	c.srvc.Popular().Send(w)
}

func (c *ProductController) Search(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	query, ok := vars["s"]
	if !ok {
		libs.GetResponse("Missing query parameter", 400, true).Send(w)
		return
	}

	c.srvc.Search(query[0]).Send(w)
}

func (c *ProductController) Add(w http.ResponseWriter, r *http.Request) {

	var data models.Product

	imageName := r.Context().Value("imageName").(string)
	data.Image = imageName

	err := schema.NewDecoder().Decode(&data, r.MultipartForm.Value)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.Add(&data).Send(w)
}

func (c *ProductController) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	var data models.Product

	imageName := r.Context().Value("imageName").(string)
	data.Image = imageName

	err := schema.NewDecoder().Decode(&data, r.MultipartForm.Value)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.Update(id, &data).Send(w)
}

func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.Delete(id).Send(w)
}
