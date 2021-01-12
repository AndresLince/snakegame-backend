package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../domain"
	"../entity"
	"../repository"
	"../utilities"
	"github.com/jackc/pgx/pgxpool"
)

type ScoreHandler struct {
	repository domain.ScoreInterface
}

func NewScoresHandler(pool *pgxpool.Pool) *ScoreHandler {
	return &ScoreHandler{
		repository: repository.NewSQLScoreRepository(pool),
	}
}

func (scoreHandler *ScoreHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	limit := utilities.ParseInt64(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}

	payload, err := scoreHandler.repository.GetAll(r.Context(), limit)
	if err != nil {
		utilities.ErrorResponse(w, http.StatusInternalServerError, "Server Error")
		return
	}

	utilities.Response(w, http.StatusOK, payload)
	return
}

func (scoreHandler *ScoreHandler) Create(w http.ResponseWriter, r *http.Request) {

	score := entity.Score{}
	json.NewDecoder(r.Body).Decode(&score)

	id, err := scoreHandler.repository.Create(r.Context(), &score)
	if err != nil {
		utilities.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	lastId := strconv.Itoa(id)

	utilities.Response(w, http.StatusCreated, map[string]string{"message": "Score creado correctamente.", "id": lastId})
}
