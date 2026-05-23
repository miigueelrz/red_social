package handlers

import (
	"net/http"
	"strconv"

	"red_social/internal/middleware"
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

	posts, err := h.postService.GetTimeline(int(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := templates.Timeline(posts)
	component.Render(r.Context(), w)
}

func (h *PostHandler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("content")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	post, err := h.postService.CreatePost(int(userID), content)
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

	component := templates.LikeButton(postID, likesCount, userLiked)
	component.Render(r.Context(), w)
}
