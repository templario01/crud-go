package repository

import (
	"crud-go/src/entity"
	"errors"

	"github.com/google/uuid"
)

/*
* "PostUseCase" va a trabajar con esta interfaz -> "PostRepository".
 */
type PostRepository interface {
	GetAll() ([]entity.Post, error)
	GetById(id uuid.UUID) (entity.Post, error)
	AddPost(post *entity.CreatePostRequest)
}

/*
  - Pero se va a crear una instancia de "PostRepositoryDB", por eso "PostRepositoryDB" debe tener todos los metodos de "PostRepository" implementados.
    postRepo := repository.NewPostRepositoryDB()
    postUsecase := usecase.NewPostUsecase(postRepo)

*
*/
type PostRepositoryDB struct {
	posts []entity.Post
}

func (pr *PostRepositoryDB) GetAll() ([]entity.Post, error) {
	return pr.posts, nil
}

func (pr *PostRepositoryDB) GetById(id uuid.UUID) (entity.Post, error) {
	for _, post := range pr.posts {
		if post.ID == id {
			return post, nil
		}
	}
	return entity.Post{}, errors.New("post not found")
}

func (pr *PostRepositoryDB) AddPost(post *entity.CreatePostRequest) {
	newPost := &entity.Post{ID: uuid.New(), Title: post.Title, Content: post.Content}
	pr.posts = append(pr.posts, *newPost)
}

func NewPostRepositoryDB() *PostRepositoryDB {
	return &PostRepositoryDB{
		posts: []entity.Post{
			{ID: uuid.New(), Title: "Post 1", Content: "Content 1"},
			{ID: uuid.New(), Title: "Post 2", Content: "Content 2"},
		},
	}
}
