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
	return pu.postRepository.GetAll()
}

func (pu *postUsecase) GetPostById(id uuid.UUID) (entity.Post, error) {
	return pu.postRepository.GetById(id)
}

func (pu *postUsecase) CreatePost(post *entity.CreatePostRequest) error {
	pu.postRepository.AddPost(post)
	return nil
}
