package repository

import (
	"context"

	"github.com/yagrush/go-sample-restapi/app/domain/model"
)

type SampleRepository interface {
	GetFuga(context.Context) (*model.Fuga, error)
	CalcAddInt64(context.Context, int64, int64) (int64, error)
}
