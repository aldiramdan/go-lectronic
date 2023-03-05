package reviews

import (
	"encoding/json"
	"lectronic/src/databases/orm/models"
	"lectronic/src/interfaces"
	"lectronic/src/libs"
	"net/http"

	"github.com/gorilla/mux"
)

type ReviewController struct {
	srvc interfaces.ReviewService
}

func NewController(srvc interfaces.ReviewService) *ReviewController {
	return &ReviewController{srvc}
}

func (c *ReviewController) GetByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.GetByID(id).Send(w)
}

func (c *ReviewController) GetByProductID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.GetByProductID(id).Send(w)
}

func (c *ReviewController) Add(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	var data *models.Review

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	data.UserID = user_id.(string)

	c.srvc.Add(data).Send(w)
}
