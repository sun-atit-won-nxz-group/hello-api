package handler

import (
	"hello-api/service"
	"net/http"
	"strconv"
)

type Calculator interface {
	Sum(w http.ResponseWriter, r *http.Request)
}

type CalculatorHandler struct {
	service service.CalculatorService
}

func NewCalculatorHandler(service service.CalculatorService) *CalculatorHandler {
	return &CalculatorHandler{
		service: service,
	}
}

func (h *CalculatorHandler) Sum(w http.ResponseWriter, r *http.Request) {
	resp := h.service.Sum(1, 2)

	w.Write([]byte(strconv.Itoa(resp)))
}
