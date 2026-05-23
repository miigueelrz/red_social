package handlers

import (
	"fmt"
	"io"
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

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) HandleGetTimeline(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	username, _ := r.Context().Value(middleware.UsernameKey).(string)

	posts, err := h.postService.GetTimeline(int(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := templates.Timeline(username, posts)
	component.Render(r.Context(), w)
}

func (h *PostHandler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "error parsing form", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	imageURL := ""
	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		uploadDir := "./static/uploads"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			http.Error(w, "error creating upload directory", http.StatusInternalServerError)
			return
		}

		ext := filepath.Ext(header.Filename)
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		dst, err := os.Create(filepath.Join(uploadDir, filename))
		if err != nil {
			http.Error(w, "error saving file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "error saving file", http.StatusInternalServerError)
			return
		}

		imageURL = "/static/uploads/" + filename
	}

	post, err := h.postService.CreatePost(int(userID), content, imageURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
