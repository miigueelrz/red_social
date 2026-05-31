package handlers

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"red_social/internal/middleware"
	"red_social/internal/models"
	"red_social/internal/services"
	"red_social/internal/templates"
)

func writeHTMXError(w http.ResponseWriter, msg string, status int) {
	log.Printf("HTMX error (status=%d): %s", status, msg)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	w.Write([]byte(msg))
}

type PostHandler struct {
	postService *services.PostService
	userService *services.UserService
}

func NewPostHandler(postService *services.PostService, userService *services.UserService) *PostHandler {
	return &PostHandler{postService: postService, userService: userService}
}

const maxPostUploadSize = 20 << 20

func (h *PostHandler) HandleGetTimeline(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		component := templates.Landing()
		component.Render(r.Context(), w)
		return
	}

	user, err := h.userService.GetProfile(userID)
	if err != nil {
		log.Printf("ERROR getting user profile in timeline: %v", err)
		writeHTMXError(w, "Error al cargar el usuario.", http.StatusInternalServerError)
		return
	}

	posts, err := h.postService.GetTimeline(int(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := templates.Timeline(user, posts)
	component.Render(r.Context(), w)
}

func (h *PostHandler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxPostUploadSize)

	mediaType, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	isMultipart := mediaType == "multipart/form-data"
	if isMultipart {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			writeHTMXError(w, "No se pudo procesar el formulario o la imagen supera los 20 MB.", http.StatusBadRequest)
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			writeHTMXError(w, "Error al procesar el formulario.", http.StatusBadRequest)
			return
		}
	}

	content := r.FormValue("content")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		writeHTMXError(w, "No autorizado.", http.StatusUnauthorized)
		return
	}

	parentIDStr := r.FormValue("parent_id")
	var parentID *int
	if parentIDStr != "" {
		pid, err := strconv.Atoi(parentIDStr)
		if err != nil {
			writeHTMXError(w, "ID de publicación inválido.", http.StatusBadRequest)
			return
		}
		parentID = &pid
	}

	imageURL := ""
	if isMultipart {
		file, header, err := r.FormFile("image")
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
				writeHTMXError(w, "Error interno al preparar el directorio de imágenes. Intenta de nuevo.", http.StatusInternalServerError)
				return
			}

			filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
			dst, err := os.Create(filepath.Join(uploadDir, filename))
			if err != nil {
				log.Printf("ERROR creating file %s: %v", filename, err)
				writeHTMXError(w, "Error interno al guardar la imagen. Intenta de nuevo.", http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			if _, err := io.Copy(dst, file); err != nil {
				log.Printf("ERROR copying file %s: %v", filename, err)
				writeHTMXError(w, "Error interno al guardar la imagen. Intenta de nuevo.", http.StatusInternalServerError)
				return
			}

			imageURL = "/static/uploads/" + filename
		} else if err != http.ErrMissingFile {
			writeHTMXError(w, "Error al leer el archivo subido.", http.StatusBadRequest)
			return
		}
	}

	post, err := h.postService.CreatePost(int(userID), content, imageURL, parentID)
	if err != nil {
		writeHTMXError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := r.Context().Value(middleware.UserKey).(*models.User)
	if user != nil {
		post.Author = user.Username
		post.AuthorName = user.Name
		post.AuthorAvatarURL = user.AvatarURL
	}

	if parentID != nil {
		component := templates.ReplyItem(*post)
		component.Render(r.Context(), w)
		return
	}

	component := templates.PostItem(*post)
	component.Render(r.Context(), w)
}

func (h *PostHandler) HandleToggleLike(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	postIDStr := r.PathValue("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "invalid post id", http.StatusBadRequest)
		return
	}

	likesCount, userLiked, err := h.postService.ToggleLike(int(userID), postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := templates.LikeButton(models.Post{
		ID:         postID,
		LikesCount: likesCount,
		UserLiked:  userLiked,
	})
	component.Render(r.Context(), w)
}
