package services

import (
	"errors"
	"fmt"

	"red_social/internal/models"
)

type Post = models.Post

type PostRepository interface {
	CreatePost(post *Post) error
	GetRecentPosts(currentUserID int) ([]Post, error)
	ToggleLike(userID, postID int) (bool, error)
	GetLikesCount(postID int) (int, error)
}

type PostService struct {
	postRepo PostRepository
}

func NewPostService(postRepo PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) CreatePost(userID int, content, imageURL string) (*Post, error) {
	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(content) > 280 {
		return nil, errors.New("content must not exceed 280 characters")
	}

	post := &Post{
		UserID:   userID,
		Content:  content,
		ImageURL: imageURL,
	}

	if err := s.postRepo.CreatePost(post); err != nil {
		return nil, fmt.Errorf("error creating post: %w", err)
	}

	return post, nil
}

func (s *PostService) GetTimeline(currentUserID int) ([]Post, error) {
	posts, err := s.postRepo.GetRecentPosts(currentUserID)
	if err != nil {
		return nil, fmt.Errorf("error getting timeline: %w", err)
	}

	return posts, nil
}

func (s *PostService) ToggleLike(userID, postID int) (int, bool, error) {
	userLiked, err := s.postRepo.ToggleLike(userID, postID)
	if err != nil {
		return 0, false, err
	}
	
	likesCount, err := s.postRepo.GetLikesCount(postID)
	if err != nil {
		return 0, false, err
	}
	
	return likesCount, userLiked, nil
}
