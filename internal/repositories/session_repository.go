package repositories

import (
	"database/sql"
	"fmt"

	"red_social/internal/models"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) CreateSession(session *models.Session) error {
	query := `INSERT INTO sessions (user_id, token_hash, expires_at) VALUES ($1, $2, $3) RETURNING id, created_at`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing create session statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(session.UserID, session.TokenHash, session.ExpiresAt).
		Scan(&session.ID, &session.CreatedAt)
	if err != nil {
		return fmt.Errorf("error creating session: %w", err)
	}

	return nil
}

func (r *SessionRepository) GetSessionByTokenHash(tokenHash string) (*models.Session, error) {
	query := `SELECT id, user_id, token_hash, expires_at, created_at FROM sessions WHERE token_hash = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing get session statement: %w", err)
	}
	defer stmt.Close()

	session := &models.Session{}
	err = stmt.QueryRow(tokenHash).Scan(
		&session.ID,
		&session.UserID,
		&session.TokenHash,
		&session.ExpiresAt,
		&session.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting session: %w", err)
	}

	return session, nil
}

func (r *SessionRepository) DeleteSession(id int64) error {
	query := `DELETE FROM sessions WHERE id = $1`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing delete session statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("error deleting session: %w", err)
	}

	return nil
}
