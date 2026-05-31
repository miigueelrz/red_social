package services

import (
	"fmt"

	"red_social/internal/models"
)

type UserProfileRepository interface {
	FindByID(id int64) (*models.User, error)
	UpdateProfile(userID int64, name, avatarURL string) error
}

type UserService struct {
	userRepo UserProfileRepository
}

func NewUserService(userRepo UserProfileRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetProfile(userID int64) (*models.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, fmt.Errorf("error getting profile: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *UserService) UpdateProfile(userID int64, name, avatarURL string) error {
	if name == "" {
		name = ""
	}

	if len(name) > 100 {
		return fmt.Errorf("el nombre no puede superar los 100 caracteres")
	}

	if err := s.userRepo.UpdateProfile(userID, name, avatarURL); err != nil {
		return fmt.Errorf("error updating profile: %w", err)
	}

	return nil
}
