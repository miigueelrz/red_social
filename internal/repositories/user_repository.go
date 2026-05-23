package repositories

import (
	"database/sql"
	"fmt"

	"red_social/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (username, email, password_hash, bio, avatar_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing create user statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		user.Username,
		user.Email,
		user.PasswordHash,
		user.Bio,
		user.AvatarURL,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	query := `SELECT id, username, email, password_hash, bio, avatar_url, created_at, updated_at
		FROM users WHERE email = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing get user by email statement: %w", err)
	}
	defer stmt.Close()

	user := &models.User{}
	err = stmt.QueryRow(email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Bio,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting user by email: %w", err)
	}

	return user, nil
}

func (r *UserRepository) FindByID(id int64) (*models.User, error) {
	query := `SELECT id, username, email, password_hash, bio, avatar_url, created_at, updated_at
		FROM users WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing get user by id statement: %w", err)
	}
	defer stmt.Close()

	user := &models.User{}
	err = stmt.QueryRow(id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Bio,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting user by id: %w", err)
	}

	return user, nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, email, password_hash, bio, avatar_url, created_at, updated_at
		FROM users WHERE username = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing get user by username statement: %w", err)
	}
	defer stmt.Close()

	user := &models.User{}
	err = stmt.QueryRow(username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Bio,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting user by username: %w", err)
	}

	return user, nil
}
