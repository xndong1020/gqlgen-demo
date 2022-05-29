package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"acy.com/gqlgendemo/graph/generated"
	"acy.com/gqlgendemo/graph/model"
	"acy.com/gqlgendemo/utilities"
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

func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (string, error) {
	err := r.BookService.DeleteBook(id)
	if err != nil {
		return "", err
	}
	successMessage := "successfully deleted"
	return successMessage, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id int, input model.BookInput) (string, error) {
	err := r.BookService.UpdateBook(&input, id)
	if err != nil {
		return "nil", err
	}
	successMessage := "successfully updated"

	return successMessage, nil
}

func (r *mutationResolver) PostMessage(ctx context.Context, user string, content string) (string, error) {
	// Construct the newly sent message and append it to the existing messages
	msg := model.Message{
		ID:      strconv.Itoa(len(r.ChatMessages)),
		User:    user,
		Content: content,
	}
	r.ChatMessages = append(r.ChatMessages, &msg)
	r.mu.Lock()
	// Notify all active subscriptions that a new message has been posted by posted. In this case we push the now
	// updated ChatMessages array to all clients that care about it.
	for _, observer := range r.ChatObservers {
		observer <- r.ChatMessages
	}
	r.mu.Unlock()
	return msg.ID, nil
}

func (r *queryResolver) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	books, err := r.BookService.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
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

func (r *subscriptionResolver) Messages(ctx context.Context) (<-chan []*model.Message, error) {
	// Create an ID and channel for each active subscription. We will push changes into this channel.
	// When a new subscription is created by the client, this resolver will fire first.
	id := utilities.RandString(8)
	msgs := make(chan []*model.Message, 1)

	// Start a goroutine to allow for cleaning up subscriptions that are disconnected.
	// This go routine will only get past Done() when a client terminates the subscription. This allows us
	// to only then remove the reference from the list of ChatObservers since it is no longer needed.
	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(r.ChatObservers, id)
		r.mu.Unlock()
	}()

	r.mu.Lock()
	// Keep a reference of the channel so that we can push changes into it when new messages are posted.
	r.ChatObservers[id] = msgs
	r.mu.Unlock()
	// This is optional, and this allows newly subscribed clients to get a list of all the messages that have been
	// posted so far. Upon subscribing the client will be pushed the messages once, further changes are handled
	// in the PostMessage mutation.
	r.ChatObservers[id] <- r.ChatMessages
	return msgs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }


