package graph

import (
	"acy.com/gqlgendemo/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	BookService service.IBookService
}
