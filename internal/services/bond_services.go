package services

import (
	"math"

	"github.com/cryskram/hercules/internal/dto"
	"github.com/cryskram/hercules/internal/models"
	repository "github.com/cryskram/hercules/internal/repositories"
)

type BondService interface {
	GetAll(filter dto.BondFilter) (*dto.PaginatedResponse[models.Bond], error)
	GetByISIN(isin string) (*models.Bond, error)
}

type bondService struct {
	repo repository.BondRepository
}

func NewBondService(repo repository.BondRepository) BondService {
	return &bondService{
		repo: repo,
	}
}

func (s *bondService) GetByISIN(isin string) (*models.Bond, error) {
	return s.repo.GetByISIN(isin)
}

func (s *bondService) GetAll(filter dto.BondFilter) (*dto.PaginatedResponse[models.Bond], error) {

	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.Limit <= 0 {
		filter.Limit = 20
	}

	if filter.Limit > 100 {
		filter.Limit = 100
	}

	if filter.Sort == "" {
		filter.Sort = "yield"
	}

	if filter.Order == "" {
		filter.Order = "desc"
	}

	bonds, total, err := s.repo.GetAll(filter)

	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(filter.Limit)))

	return &dto.PaginatedResponse[models.Bond]{
		Data: bonds,

		Meta: dto.PaginationMeta{
			Page:  filter.Page,
			Limit: filter.Limit,

			TotalItems: total,
			TotalPages: totalPages,

			HasNext:     filter.Page < totalPages,
			HasPrevious: filter.Page > 1,
		},
	}, nil
}
