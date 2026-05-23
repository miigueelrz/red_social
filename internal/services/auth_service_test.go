package services

import (
	"errors"
	"testing"
)

type MockUserRepository struct {
	CreateFn         func(user *User) error
	FindByEmailFn    func(email string) (*User, error)
	FindByUsernameFn func(username string) (*User, error)
}

func (m *MockUserRepository) Create(user *User) error {
	if m.CreateFn != nil { return m.CreateFn(user) }
	return nil
}

func (m *MockUserRepository) FindByEmail(email string) (*User, error) {
	if m.FindByEmailFn != nil { return m.FindByEmailFn(email) }
	return nil, nil
}

func (m *MockUserRepository) FindByUsername(username string) (*User, error) {
	if m.FindByUsernameFn != nil { return m.FindByUsernameFn(username) }
	return nil, nil
}

func TestAuthService_Register(t *testing.T) {
	t.Run("Registro exitoso", func(t *testing.T) {
		mockRepo := &MockUserRepository{ CreateFn: func(user *User) error { return nil } }
		service := NewAuthService(mockRepo)
		err := service.Register("john_doe", "john@example.com", "password123")
		if err != nil { t.Errorf("Se esperaba un registro exitoso, pero se obtuvo error: %v", err) }
	})

	t.Run("Error por campos vacíos", func(t *testing.T) {
		mockRepo := &MockUserRepository{}
		service := NewAuthService(mockRepo)
		err := service.Register("", "john@example.com", "password123")
		if err == nil { t.Error("Se esperaba error por nombre vacio") }
		err = service.Register("john_doe", "", "password123")
		if err == nil { t.Error("Se esperaba error por email vacio") }
		err = service.Register("john_doe", "john@example.com", "")
		if err == nil { t.Error("Se esperaba error por password vacia") }
	})

	t.Run("Error por usuario o email duplicado", func(t *testing.T) {
		mockRepo := &MockUserRepository{ CreateFn: func(user *User) error { return errors.New("unique constraint violation") } }
		service := NewAuthService(mockRepo)
		err := service.Register("duplicate_user", "duplicate@example.com", "password123")
		if err == nil { t.Error("Se esperaba error por datos duplicados") }
	})
}
