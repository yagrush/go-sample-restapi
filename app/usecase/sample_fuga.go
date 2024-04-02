package usecase

import (
	"fmt"
	"net/http"

	"github.com/yagrush/go-sample-restapi/app/domain/repository"
)

type SampleFugaUsecase struct {
	R repository.SampleRepository
}

func (u *SampleFugaUsecase) Serve(w http.ResponseWriter, r *http.Request) error {
	m, err := u.R.GetFuga(r.Context())
	if err != nil {
		return err
	}

	fmt.Fprint(w, m.GetMessage())
	return nil
}
