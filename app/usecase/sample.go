package usecase

import (
	"fmt"
	"net/http"
	"strconv"

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

type SampleCalcAddInt64Usecase struct {
	R repository.SampleRepository
}

func (u *SampleCalcAddInt64Usecase) Serve(w http.ResponseWriter, r *http.Request) error {
	aStr := r.FormValue("a")
	a, err := strconv.ParseInt(aStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid request parameter a=%s", aStr)
	}

	bStr := r.FormValue("b")
	b, err := strconv.ParseInt(bStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid request parameter b=%s", bStr)
	}

	ret, err := u.R.CalcAddInt64(r.Context(), a, b)
	if err != nil {
		return fmt.Errorf("error: CalcAddInt64")
	}

	fmt.Fprint(w, ret)
	return nil
}
