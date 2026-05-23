package services

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"red_social/internal/models"
)

type User = models.User

type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
}

type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(email, password string) (*User, error) {
	email = strings.TrimSpace(email)

	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *AuthService) Register(username, email, password string) error {
	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)

	if username == "" || email == "" || password == "" {
		return errors.New("username, email, and password are required")
	}

	existing, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return fmt.Errorf("error checking existing email: %w", err)
	}
	if existing != nil {
		return errors.New("email already registered")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	user := &User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedBytes),
	}

	if err := s.userRepo.Create(user); err != nil {
		return fmt.Errorf("error saving user: %w", err)
	}

	return nil
}
