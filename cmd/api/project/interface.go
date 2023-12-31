package project

import (
	"github.com/tekpriest/poprev/internal/model"
	"github.com/tekpriest/poprev/internal/query"
)

type CreateProjectData struct {
	Amount     float64 `json:"amount,omitempty"`
	ArtistName string  `json:"artist_name,omitempty"`
}

type QueryProjectData struct {
	Limit  int64  `query:"limit"`
	Page   int64  `query:"page"`
	Status string `query:"status"`
}

type FetchProjectsData struct {
	Projects []model.Project  `json:"projects"`
	Meta     query.Pagination `json:"meta,omitempty"`
}

type FetchAllProjectTransactionsData struct {
	Transactions []model.Transaction `json:"transactions"`
	Meta         query.Pagination    `json:"meta,omitempty"`
}

type QueryTransactionsData struct {
	Limit  int64  `query:"limit"`
	Page   int64  `query:"page"`
	Status string `query:"status"`
}
