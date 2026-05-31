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
	GetReplies(postID, currentUserID int) ([]Post, error)
	GetPostByID(postID int) (*Post, error)
	UpdatePost(postID, userID int, content, imageURL string) error
	ToggleLike(userID, postID int) (bool, error)
	GetLikesCount(postID int) (int, error)
	DeletePost(postID, userID int) error
}

type PostService struct {
	postRepo PostRepository
}

func NewPostService(postRepo PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) CreatePost(userID int, content, imageURL string, parentID *int) (*Post, error) {
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
		ParentID: parentID,
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

	for i := range posts {
		replies, err := s.getRepliesTree(posts[i].ID, currentUserID)
		if err != nil {
			return nil, fmt.Errorf("error getting replies for post %d: %w", posts[i].ID, err)
		}
		posts[i].Replies = replies
	}

	return posts, nil
}

func (s *PostService) getRepliesTree(postID, currentUserID int) ([]Post, error) {
	replies, err := s.postRepo.GetReplies(postID, currentUserID)
	if err != nil {
		return nil, err
	}

	for i := range replies {
		childReplies, err := s.getRepliesTree(replies[i].ID, currentUserID)
		if err != nil {
			return nil, err
		}
		replies[i].Replies = childReplies
	}

	return replies, nil
}

func (s *PostService) GetPostByID(postID int) (*Post, error) {
	return s.postRepo.GetPostByID(postID)
}

func (s *PostService) UpdatePost(postID, userID int, content, imageURL string) error {
	return s.postRepo.UpdatePost(postID, userID, content, imageURL)
}

func (s *PostService) DeletePost(postID, userID int) error {
	return s.postRepo.DeletePost(postID, userID)
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
