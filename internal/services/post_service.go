package services

import (
	"errors"
	"fmt"

	"red_social/internal/models"
)

type Post = models.Post

type PostRepository interface {
	CreatePost(post *Post) error
	GetRecentPosts() ([]Post, error)
}

type PostService struct {
	postRepo PostRepository
}

func NewPostService(postRepo PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) CreatePost(userID int, content string) (*Post, error) {
	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(content) > 280 {
		return nil, errors.New("content must not exceed 280 characters")
	}

	post := &Post{
		UserID:  userID,
		Content: content,
	}

	if err := s.postRepo.CreatePost(post); err != nil {
		return nil, fmt.Errorf("error creating post: %w", err)
	}

	return post, nil
}

func (s *PostService) GetTimeline() ([]Post, error) {
	posts, err := s.postRepo.GetRecentPosts()
	if err != nil {
		return nil, fmt.Errorf("error getting timeline: %w", err)
	}

	return posts, nil
}
