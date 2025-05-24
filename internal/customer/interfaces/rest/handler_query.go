package rest

import (
	"net/http"

	"github.com/Sadegh-kh/telemon/internal/customer/domain"
	"github.com/Sadegh-kh/telemon/pkg/httpx"
	"github.com/go-chi/chi"
)

func (h *CustomerQueryHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	customer, err := h.queryService.GetCustomerByID(r.Context(), domain.CustomerID(id))
	if err != nil {
		httpx.ResposndErrorJSON(w, http.StatusNotFound, err)
		return
	}

	httpx.ResponsdJSON(w, http.StatusOK, customer)
}

func (h *CustomerQueryHandler) GetCustomerByNationalID(w http.ResponseWriter, r *http.Request) {
	nationalID := chi.URLParam(r, "national_id")
	customer, err := h.queryService.GetCustomerByNatinalID(r.Context(), domain.NationalID(nationalID))
	if err != nil {
		httpx.ResposndErrorJSON(w, http.StatusNotFound, err)
		return
	}

	httpx.ResponsdJSON(w, http.StatusOK, customer)
}
