package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"acy.com/gqlgendemo/graph/generated"
	"acy.com/gqlgendemo/graph/model"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
    book, err := r.BookService.CreateBook(&input)
    bookCreated := &model.Book{
        Author:    book.Author,
        Publisher: book.Publisher,
        Title:     book.Title,
        ID:        book.ID,
    }
    if err != nil {
        return nil, err
    }
    return bookCreated, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id int, input model.BookInput) (string, error) {
    err := r.BookService.UpdateBook(&input, id)
    if err != nil {
        return "nil", err
    }
    successMessage := "successfully updated"

    return successMessage, nil
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (string, error) {
    err := r.BookService.DeleteBook(id)
    if err != nil {
        return "", err
    }
    successMessage := "successfully deleted"
    return successMessage, nil
}

func (r *queryResolver) GetOneBook(ctx context.Context, id int) (*model.Book, error) {
    book, err := r.BookService.GetOneBook(id)
    selectedBook := &model.Book{
        ID:        book.ID,
        Author:    book.Author,
        Publisher: book.Publisher,
        Title:     book.Title,
    }
    if err != nil {
        return nil, err
    }
    return selectedBook, nil
}

func (r *queryResolver) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
    books, err := r.BookService.GetAllBooks()
    if err != nil {
        return nil, err
    }
    return books, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
