package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/3dw1nM0535/sealer/graph/model"
)

type Resolver struct {
	users      []*model.User
	todos      []*model.Todo
	lastTodoId int
}

func NewResolver() *Resolver {
	users := make([]*model.User, 0)
	users = append(users, &model.User{ID: "1", Name: "John Doe"})
	users = append(users, &model.User{ID: "2", Name: "Jane Doe"})

	return &Resolver{
		users: users,
		todos: make([]*model.Todo, 0),
	}
}
