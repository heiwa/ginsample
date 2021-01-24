package model

import "errors"

type Todo struct {
	Id int `json:id`
	Title string `json:title`
	IsFinish bool `json:isFinish`
}

type Todos struct {
	Todos []Todo
}

func NewTodos() *Todos {
	return &Todos{}
}

func (t *Todos) Add(todo Todo) {
	t.Todos = append(t.Todos, todo)
}

func (t *Todos) GetAll() []Todo {
	return t.Todos
}

func (t *Todos) Get(id int) (*Todo, error) {
	for _, todo := range t.Todos {
		if todo.Id == id {
			return &todo, nil
		}
	}
	return nil, errors.New("not found.")
}

func (t *Todos) Delete(id int) (*Todo, error){
	for i, todo := range t.Todos {
		if todo.Id == id {
			t.Todos = append(t.Todos[:i], t.Todos[i+1:]...)
			return &todo, nil
		}
	}
	return nil, errors.New("not found.")
}

func (t *Todos) Update(todo Todo) {
	for i, td := range t.Todos {
		if td.Id == todo.Id {
			t.Todos[i] = todo
			return
		}
	}
}