package middleware

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"red_social/internal/repositories"
)

type contextKey string

const UserIDKey contextKey = "user_id"
const UsernameKey contextKey = "username"
const UserKey contextKey = "user"

type AuthMiddleware struct {
	sessionRepo *repositories.SessionRepository
	userRepo    *repositories.UserRepository
}

func NewAuthMiddleware(sessionRepo *repositories.SessionRepository, userRepo *repositories.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{sessionRepo: sessionRepo, userRepo: userRepo}
}

func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := m.resolveSession(r)
		if ctx == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *AuthMiddleware) OptionalAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := m.resolveSession(r)
		if ctx != nil {
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func (m *AuthMiddleware) resolveSession(r *http.Request) context.Context {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return nil
	}

	hash := sha256.Sum256([]byte(cookie.Value))
	tokenHash := hex.EncodeToString(hash[:])

	session, err := m.sessionRepo.GetSessionByTokenHash(tokenHash)
	if err != nil || session == nil || time.Now().After(session.ExpiresAt) {
		return nil
	}

	user, err := m.userRepo.FindByID(session.UserID)
	if err != nil || user == nil {
		return nil
	}

	ctx := context.WithValue(r.Context(), UserIDKey, session.UserID)
	ctx = context.WithValue(ctx, UsernameKey, user.Username)
	ctx = context.WithValue(ctx, UserKey, user)
	return ctx
}
