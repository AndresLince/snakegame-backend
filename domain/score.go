package domain

import (
	"context"

	"../entity"
)

type ScoreInterface interface {
	GetAll(ctx context.Context, limit int64) ([]*entity.Score, error)
	Create(ctx context.Context, p *entity.Score) (int, error)
}
