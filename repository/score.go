package repository

import (
	"context"
	"errors"
	"fmt"

	"../domain"
	"../entity"
	"github.com/jackc/pgx/pgxpool"
)

type ScoreRepository struct {
	Conn *pgxpool.Pool
}

func NewSQLScoreRepository(pool *pgxpool.Pool) domain.ScoreInterface {
	return &ScoreRepository{
		Conn: pool,
	}
}

func (scoreRepository *ScoreRepository) Create(ctx context.Context, p *entity.Score) (int, error) {
	query := "INSERT INTO score (username, score, created_at) VALUES ( $1, $2, current_timestamp) RETURNING Id"
	lastInsertId := 0

	if p.Score == 0 {
		return 0, errors.New("missing or invalid score")
	}
	if p.UserName == "" {
		return 0, errors.New("missing or invalid user name")
	}

	err := scoreRepository.Conn.QueryRow(ctx, query, p.UserName, p.Score).Scan(&lastInsertId)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return lastInsertId, nil
}

func (sr *ScoreRepository) GetAll(ctx context.Context, limit int64) ([]*entity.Score, error) {

	query := "SELECT id, username, score, created_at FROM score ORDER BY score DESC, created_at ASC LIMIT $1"
	return sr.fetch(ctx, query, limit)
}

func (scoreRepository *ScoreRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*entity.Score, error) {
	rows, err := scoreRepository.Conn.Query(ctx, query, args[0])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	scores := make([]*entity.Score, 0)
	for rows.Next() {
		data := new(entity.Score)

		err := rows.Scan(&data.Id, &data.UserName, &data.Score, &data.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		scores = append(scores, data)
	}

	return scores, nil
}
