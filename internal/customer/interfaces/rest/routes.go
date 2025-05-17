package rest

import "github.com/go-chi/chi"

func (h *CustomerHandler) RegisterRouters(r chi.Router) {
	r.Get("/{id}", h.GetCustomerByID)
	r.Get("/{national_id}", h.GetCustomerByNationalID)
	r.Post("/", h.CreateCustomer)
	r.Put("/{id}", h.UpdateCustomer)
	r.Delete("/{id}", h.DeleteCustomer)
}
