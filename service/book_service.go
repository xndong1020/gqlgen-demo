package service

import (
	"acy.com/gqlgendemo/graph/model"
	"acy.com/gqlgendemo/models"
	"acy.com/gqlgendemo/repository"
)

type IBookService interface {
	CreateBook(bookInput *model.BookInput) (*models.Book, error)
	UpdateBook(bookInput *model.BookInput, id int) error
	DeleteBook(id int) error
	GetOneBook(id int) (*models.Book, error)
	GetAllBooks() ([]*model.Book, error)
} 

type BookService struct {
	BookRepository *repository.BookRepository
}

func NewBookService(bookRepository *repository.BookRepository) *BookService {
	return &BookService{
		BookRepository: bookRepository,
	}
}

// CreateBook implements IBookService
func (b *BookService) CreateBook(bookInput *model.BookInput) (*models.Book, error) {
	return b.BookRepository.CreateBook(bookInput)
}

// DeleteBook implements IBookService
func (b *BookService) DeleteBook(id int) error {
	return b.BookRepository.DeleteBook(id)
}

// GetAllBooks implements IBookService
func (b *BookService) GetAllBooks() ([]*model.Book, error) {
	return b.BookRepository.GetAllBooks()
}

// GetOneBook implements IBookService
func (b *BookService) GetOneBook(id int) (*models.Book, error) {
	return b.BookRepository.GetOneBook(id)
}

// UpdateBook implements IBookService
func (b *BookService) UpdateBook(bookInput *model.BookInput, id int) error {
	return b.BookRepository.UpdateBook(bookInput, id)
}
