package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/yagrush/go-sample-restapi/app/usecase"
)

type SampleFugaHandler struct {
	U usecase.SampleFugaUsecase
	V *validator.Validate
}

func (h *SampleFugaHandler) Serve(w http.ResponseWriter, req *http.Request) error {
	if err := h.V.Struct(req); err != nil {
		return err
	}
	return h.U.Serve(w, req)
}

type SampleCalcAddInt64Handler struct {
	U usecase.SampleCalcAddInt64Usecase
	V *validator.Validate
}

func (h *SampleCalcAddInt64Handler) Serve(w http.ResponseWriter, req *http.Request) error {
	if err := h.V.Struct(req); err != nil {
		return err
	}
	return h.U.Serve(w, req)
}
