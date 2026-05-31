package handlers

import (
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"red_social/internal/middleware"
	"red_social/internal/services"
	"red_social/internal/templates"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	username, _ := r.Context().Value(middleware.UsernameKey).(string)

	user, err := h.userService.GetProfile(userID)
	if err != nil {
		log.Printf("ERROR getting profile for user %d: %v", userID, err)
		writeHTMXError(w, "Error al cargar el perfil.", http.StatusInternalServerError)
		return
	}

	component := templates.ProfilePage(user, username)
	component.Render(r.Context(), w)
}

func (h *UserHandler) HandleEditProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		writeHTMXError(w, "No autorizado.", http.StatusUnauthorized)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	mediaType, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	isMultipart := mediaType == "multipart/form-data"
	if isMultipart {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			writeHTMXError(w, "Error al procesar el formulario.", http.StatusBadRequest)
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			writeHTMXError(w, "Error al procesar el formulario.", http.StatusBadRequest)
			return
		}
	}

	name := r.FormValue("name")

	avatarURL := ""
	if isMultipart {
		file, header, err := r.FormFile("avatar")
		if err == nil {
			defer file.Close()

			allowedExts := map[string]bool{
				".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true,
			}
			ext := filepath.Ext(header.Filename)
			if !allowedExts[ext] {
				writeHTMXError(w, "Formato de imagen no soportado. Usa JPG, PNG, GIF o WebP.", http.StatusBadRequest)
				return
			}

			uploadDir := "./static/uploads"
			if err := os.MkdirAll(uploadDir, 0755); err != nil {
				log.Printf("ERROR creating upload dir %s: %v", uploadDir, err)
				writeHTMXError(w, "Error interno al preparar el directorio de imágenes.", http.StatusInternalServerError)
				return
			}

			filename := "avatar_" + time.Now().Format("20060102_150405") + ext
			dst, err := os.Create(filepath.Join(uploadDir, filename))
			if err != nil {
				log.Printf("ERROR creating avatar file %s: %v", filename, err)
				writeHTMXError(w, "Error interno al guardar la imagen.", http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			if _, err := io.Copy(dst, file); err != nil {
				log.Printf("ERROR copying avatar file %s: %v", filename, err)
				writeHTMXError(w, "Error interno al guardar la imagen.", http.StatusInternalServerError)
				return
			}

			avatarURL = "/static/uploads/" + filename
		} else if err != http.ErrMissingFile {
			writeHTMXError(w, "Error al leer el archivo subido.", http.StatusBadRequest)
			return
		}
	}

	if err := h.userService.UpdateProfile(userID, name, avatarURL); err != nil {
		writeHTMXError(w, err.Error(), http.StatusBadRequest)
		return
	}

	isHTMX := r.Header.Get("HX-Request") == "true"
	if isHTMX {
		w.Header().Set("HX-Redirect", "/profile")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}
