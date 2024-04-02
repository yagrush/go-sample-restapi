package infrastructure

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/yagrush/go-sample-restapi/app/adapter/api"
	"github.com/yagrush/go-sample-restapi/app/handler"
	"github.com/yagrush/go-sample-restapi/app/infrastructure/persistence"
	"github.com/yagrush/go-sample-restapi/app/usecase"
	"github.com/yagrush/go-sample-restapi/app/util"
)

func NewMux() (*http.ServeMux, error) {
	mux := http.NewServeMux()
	vali := validator.New()

	mux.HandleFunc("/sample/fuga", api.Middleware(util.New(handler.SampleFugaHandler{
		U: usecase.SampleFugaUsecase{
			R: persistence.NewSamplePersistence(),
		},
		V: vali,
	}).Serve, http.MethodGet))

	mux.HandleFunc("/sample/calcAddInt64", api.Middleware(util.New(handler.SampleCalcAddInt64Handler{
		U: usecase.SampleCalcAddInt64Usecase{
			R: persistence.NewSamplePersistence(),
		},
		V: vali,
	}).Serve, http.MethodGet))

	return mux, nil
}
