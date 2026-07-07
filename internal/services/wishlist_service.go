package services

import (
	"github.com/cryskram/hercules/internal/dto"
	"github.com/cryskram/hercules/internal/models"
	repository "github.com/cryskram/hercules/internal/repositories"
)

type WishlistService interface {
	Create(req dto.CreateWishlistRequest) error
	GetAll() ([]dto.WishlistResponse, error)
	GetByID(id string) (*dto.WishlistResponse, error)
	Update(id string, req dto.UpdateWishlistRequest) error
	Delete(id string) error
	AddBond(wishlistID string, bondISIN string) error
	RemoveBond(wishlistID string, bondISIN string) error
	GetWishlistBonds(id string) ([]dto.BondResponse, error)
}

type wishlistService struct {
	wishlistRepo repository.WishlistRepository
	bondRepo     repository.BondRepository
}

type WishlistWithCount struct {
	ID          string
	Name        string
	Description *string
	Color       *string
	BondCount   int64
}

func NewWishlistService(
	wishlistRepo repository.WishlistRepository,
	bondRepo repository.BondRepository,
) WishlistService {

	return &wishlistService{
		wishlistRepo: wishlistRepo,
		bondRepo:     bondRepo,
	}
}

func (s *wishlistService) Create(
	req dto.CreateWishlistRequest,
) error {

	wishlist := models.Wishlist{
		Name:        req.Name,
		Description: req.Description,
		Color:       req.Color,
	}

	return s.wishlistRepo.Create(&wishlist)
}

func (s *wishlistService) GetAll() ([]dto.WishlistResponse, error) {
	wishlists, err := s.wishlistRepo.GetAll()

	if err != nil {
		return nil, err
	}

	response := make([]dto.WishlistResponse, 0)

	for _, wishlist := range wishlists {
		count, _ := s.wishlistRepo.GetBondCount(wishlist.ID)
		response = append(response, dto.WishlistResponse{
			ID:          wishlist.ID,
			Name:        wishlist.Name,
			Description: wishlist.Description,
			Color:       wishlist.Color,
			BondCount:   int(count),
		})
	}

	return response, nil
}

func (s *wishlistService) AddBond(
	wishlistID string,
	bondISIN string,
) error {
	_, err := s.wishlistRepo.GetByID(wishlistID)

	if err != nil {
		return err
	}

	_, err = s.bondRepo.GetByISIN(bondISIN)

	if err != nil {
		return err
	}

	return s.wishlistRepo.AddBond(
		wishlistID,
		bondISIN,
	)
}

func (s *wishlistService) RemoveBond(
	wishlistID string,
	bondISIN string,
) error {

	return s.wishlistRepo.RemoveBond(
		wishlistID,
		bondISIN,
	)
}

func (s *wishlistService) Delete(
	id string,
) error {
	_, err := s.wishlistRepo.GetByID(id)

	if err != nil {
		return err
	}

	return s.wishlistRepo.Delete(id)
}

func (s *wishlistService) GetWishlistBonds(
	id string,
) ([]dto.BondResponse, error) {
	_, err := s.wishlistRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	bonds, err := s.wishlistRepo.GetWishlistBonds(id)

	if err != nil {
		return nil, err
	}

	response := make([]dto.BondResponse, 0, len(bonds))

	for _, bond := range bonds {
		response = append(response, dto.ToBondResponse(bond))
	}

	return response, nil
}

func (s *wishlistService) GetByID(
	id string,
) (*dto.WishlistResponse, error) {
	wishlist, err := s.wishlistRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	count, _ := s.wishlistRepo.GetBondCount(id)

	return &dto.WishlistResponse{
		ID:          wishlist.ID,
		Name:        wishlist.Name,
		Description: wishlist.Description,
		Color:       wishlist.Color,
		BondCount:   int(count),
	}, nil
}

func (s *wishlistService) Update(
	id string,
	req dto.UpdateWishlistRequest,
) error {

	wishlist, err := s.wishlistRepo.GetByID(id)

	if err != nil {
		return err
	}

	if req.Name != "" {
		wishlist.Name = req.Name
	}

	if req.Description != nil {
		wishlist.Description = req.Description
	}

	if req.Color != nil {
		wishlist.Color = req.Color
	}

	return s.wishlistRepo.Update(wishlist)
}
