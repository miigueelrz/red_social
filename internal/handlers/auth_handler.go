package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"time"

	"red_social/internal/models"
	"red_social/internal/repositories"
	"red_social/internal/services"
	"red_social/internal/templates"
)

type AuthHandler struct {
	authService    *services.AuthService
	sessionRepo    *repositories.SessionRepository
	sessionExpiry  time.Duration
}

func NewAuthHandler(authService *services.AuthService, sessionRepo *repositories.SessionRepository) *AuthHandler {
	return &AuthHandler{
		authService:   authService,
		sessionRepo:   sessionRepo,
		sessionExpiry: 24 * time.Hour,
	}
}

func (h *AuthHandler) RegisterGET(w http.ResponseWriter, r *http.Request) {
	component := templates.RegisterForm()
	component.Render(r.Context(), w)
}

func (h *AuthHandler) RegisterPOST(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := h.authService.Register(username, email, password)
	if err != nil {
		component := templates.FormError(err.Error())
		component.Render(r.Context(), w)
		return
	}

	component := templates.FormSuccess("Registro exitoso")
	component.Render(r.Context(), w)
}

func (h *AuthHandler) LoginGET(w http.ResponseWriter, r *http.Request) {
	component := templates.LoginForm()
	component.Render(r.Context(), w)
}

func (h *AuthHandler) LoginPOST(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := h.authService.Login(email, password)
	if err != nil {
		component := templates.FormError(err.Error())
		component.Render(r.Context(), w)
		return
	}

	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		component := templates.FormError("internal error")
		component.Render(r.Context(), w)
		return
	}

	token := base64.RawURLEncoding.EncodeToString(tokenBytes)

	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])

	session := &models.Session{
		UserID:    user.ID,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(h.sessionExpiry),
	}

	if err := h.sessionRepo.CreateSession(session); err != nil {
		component := templates.FormError("internal error")
		component.Render(r.Context(), w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  session.ExpiresAt,
	})

	component := templates.FormSuccess("Inicio de sesión exitoso")
	component.Render(r.Context(), w)
}
