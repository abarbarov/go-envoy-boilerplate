package todo

import (
	"context"
	"github.com/satori/go.uuid"
	"log"
)

type Server struct {
	Todos []*TodoObject
}

func (s *Server) GetTodos(ctx context.Context, _ *GetTodoParams) (*TodoResponse, error) {
	log.Printf("get tasks")
	return &TodoResponse{Todos: s.Todos}, nil
}

func (s *Server) AddTodo(ctx context.Context, newTodo *AddTodoParams) (*TodoObject, error) {
	log.Printf("Received new task %s", newTodo.Task)
	todoObject := &TodoObject{
		Id:   uuid.NewV1().String(),
		Task: newTodo.Task,
	}
	s.Todos = append([]*TodoObject{todoObject}, s.Todos...)

	return todoObject, nil
}

func (s *Server) CompleteTodo(ctx context.Context, delTodo *CompleteTodoParams) (*CompleteResponse, error) {
	log.Printf("Complete task")
	for _, todo := range s.Todos {
		if todo.Id == delTodo.Id {
			todo.Completed = true
			break
		}
	}

	return &CompleteResponse{Message: "success"}, nil
}

func (s *Server) DeleteTodo(ctx context.Context, delTodo *DeleteTodoParams) (*DeleteResponse, error) {
	log.Printf("Delete task")
	var updatedTodos []*TodoObject
	for index, todo := range s.Todos {
		if todo.Id == delTodo.Id {
			updatedTodos = append(s.Todos[:index], s.Todos[index+1:]...)
			break
		}
	}
	s.Todos = updatedTodos

	return &DeleteResponse{Message: "success"}, nil
}
