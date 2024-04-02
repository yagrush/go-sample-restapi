package persistence

import (
	"context"

	"github.com/yagrush/go-sample-module-a/pkg/util"
	"github.com/yagrush/go-sample-restapi/app/domain/model"
	"github.com/yagrush/go-sample-restapi/app/domain/repository"
)

type SamplePersistence struct{}

func NewSamplePersistence() repository.SampleRepository {
	return &SamplePersistence{}
}

func (tp SamplePersistence) GetFuga(c context.Context) (*model.Fuga, error) {
	m := model.NewFuga("fuga")
	return m, nil
}

func (tp SamplePersistence) CalcAddInt64(c context.Context, a, b int64) (int64, error) {
	return util.CalcAddInt64(a, b), nil
}
