package Todo

import (
	"github.com/google/uuid"
)

// TODO aka-somix : Try to create a shared library among modules
type Todo struct {
	Id string
	Name string
	Description string
	Completed bool
}
func New(name string, descr string) Todo {
	return Todo{
		Id: uuid.New().String(),
		Name: name,
		Description: descr,
		Completed: false,
	}
}
