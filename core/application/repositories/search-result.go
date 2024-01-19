package repositories

import (
	"health/core/application/entities"
	"math/big"
)

type SearchResult[T entities.Entity] struct {
	Items       T       `json:"items"`
	Total       big.Int `json:"total"`
	CurrentPage int     `json:"current_page"`
	PerPage     int     `json:"per_page"`
	LastPage    int     `json:"last_page"`
}
