package usecase

import (
	"crud-go/src/entity"
	"crud-go/src/repository"

	"github.com/google/uuid"
)

type PostUsecase interface {
	GetAllPosts() ([]entity.Post, error)
	GetPostById(id uuid.UUID) (entity.Post, error)
	CreatePost(post *entity.CreatePostRequest) error
	DeletePost(id uuid.UUID) error
	UpdatePost(id uuid.UUID, post *entity.CreatePostRequest) (entity.Post, error)
}

type postUsecase struct {
	postRepository repository.PostRepository
}

func NewPostUsecase(pr repository.PostRepository) PostUsecase {
	return &postUsecase{
		postRepository: pr,
	}
}

func (pu *postUsecase) GetAllPosts() ([]entity.Post, error) {
	return pu.postRepository.FindAll()
}

func (pu *postUsecase) GetPostById(id uuid.UUID) (entity.Post, error) {
	return pu.postRepository.Find(id)
}

func (pu *postUsecase) CreatePost(post *entity.CreatePostRequest) error {
	pu.postRepository.Create(post)
	return nil
}

func (pu *postUsecase) DeletePost(id uuid.UUID) error {
	pu.postRepository.Delete(id)
	return nil
}

func (pu *postUsecase) UpdatePost(id uuid.UUID, post *entity.CreatePostRequest) (entity.Post, error) {
	return pu.postRepository.Update(id, post)
}
