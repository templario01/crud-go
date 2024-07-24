package repository

import (
	"crud-go/src/entity"
	"errors"

	"github.com/google/uuid"
)

type PostRepository interface {
	FindAll() ([]entity.Post, error)
	Find(id uuid.UUID) (entity.Post, error)
	Create(post *entity.CreatePostRequest)
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, post *entity.CreatePostRequest) (entity.Post, error)
}

type PostRepositoryDB struct {
	posts []entity.Post
}

func (pr *PostRepositoryDB) FindAll() ([]entity.Post, error) {
	return pr.posts, nil
}

func (pr *PostRepositoryDB) Find(id uuid.UUID) (entity.Post, error) {
	for _, post := range pr.posts {
		if post.ID == id {
			return post, nil
		}
	}
	return entity.Post{}, errors.New("post not found")
}

func (pr *PostRepositoryDB) Create(post *entity.CreatePostRequest) {
	newPost := &entity.Post{ID: uuid.New(), Title: post.Title, Content: post.Content}
	pr.posts = append(pr.posts, *newPost)
}

func (pr *PostRepositoryDB) Delete(id uuid.UUID) error {
	currentIndex := -1
	for i, post := range pr.posts {
		if post.ID == id {
			currentIndex = i
			break
		}
	}
	if currentIndex == -1 {
		return errors.New("post not found")
	}
	pr.posts = append(pr.posts[:currentIndex], pr.posts[currentIndex+1:]...)
	return nil
}

func (pr *PostRepositoryDB) Update(id uuid.UUID, post *entity.CreatePostRequest) (entity.Post, error) {
	currentIndex := -1
	for i, post := range pr.posts {
		if post.ID == id {
			currentIndex = i
			break
		}
	}
	if currentIndex == -1 {
		return entity.Post{}, errors.New("post not found")
	}
	pr.posts[currentIndex].Title = post.Title
	pr.posts[currentIndex].Content = post.Content

	return pr.posts[currentIndex], nil
}

func NewPostRepositoryDB() *PostRepositoryDB {
	return &PostRepositoryDB{
		posts: []entity.Post{
			{ID: uuid.New(), Title: "Post 1", Content: "Content 1"},
			{ID: uuid.New(), Title: "Post 2", Content: "Content 2"},
		},
	}
}
