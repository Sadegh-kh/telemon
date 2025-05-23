package rest

import (
	"net/http"

	"github.com/Sadegh-kh/telemon/internal/customer/application/command"
	"github.com/Sadegh-kh/telemon/internal/customer/domain"
	resquest "github.com/Sadegh-kh/telemon/internal/customer/interfaces/rest/request"
	"github.com/Sadegh-kh/telemon/pkg/httpx"
	"github.com/go-chi/chi"
)

func (h *CustomerCommandHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var req resquest.CreateCustomerRequest

	if err := httpx.BindJSON(r, &req); err != nil {
		httpx.ResposndErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	cmd := command.CreateCustomerCommand{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		NationalID:  domain.NationalID(req.NationalID),
		RegionID:    domain.RegionID(req.RegionID),
	}

	if err := h.commandService.CreateCustomer(r.Context(), cmd); err != nil {
		httpx.ResposndErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	httpx.ResponsdJSON(w, http.StatusCreated, nil)

}

func (h *CustomerCommandHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req resquest.UpdateCustomerRequest

	if err := httpx.BindJSON(r, &req); err != nil {
		httpx.ResposndErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	cmd := command.UpdateCustomerInfoCommand{
		CustomerID:  domain.CustomerID(id),
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		NationalID:  domain.NationalID(req.NationalID),
		RegionID:    domain.RegionID(req.RegionID),
	}

	if err := h.commandService.UpdateCustomer(r.Context(), cmd); err != nil {
		httpx.ResposndErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	httpx.ResponsdJSON(w, http.StatusOK, nil)
}

func (h *CustomerCommandHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.commandService.DeleteCustomer(r.Context(), domain.CustomerID(id)); err != nil {
		httpx.ResposndErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	httpx.ResponsdJSON(w, http.StatusNoContent, nil)
}
