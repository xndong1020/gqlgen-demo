package graph

import (
	"sync"

	"acy.com/gqlgendemo/graph/model"
	"acy.com/gqlgendemo/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	BookService service.IBookService
	// All messages since launching the GraphQL endpoint
    ChatMessages  []*model.Message
    // All active subscriptions
    ChatObservers map[string]chan []*model.Message
	mu            sync.Mutex
}
