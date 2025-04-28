package product

import (
	"net/http"

	"github.com/aarav345/ecom-go/types"
	"github.com/aarav345/ecom-go/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	Store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{Store: store}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodGet)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	ps, err := h.Store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}
