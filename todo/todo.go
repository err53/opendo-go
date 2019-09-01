package todo

import "errors"

type ToDo struct {
	Name     string
	Id       int
	priority int
	Contexts []string
}

func (t *ToDo) Priority() int {
	return t.priority
}

func (t *ToDo) SetPriority(priority int) error {
	if priority >= 0 && priority <= 4 {
		t.priority = priority
		return nil
	}
	return errors.New("invalid priority")
}

func NewToDo(name string, id int) *ToDo {
	return &ToDo{Name: name, Id: id}
}

func NewToDoPri(name string, id int, pri int) (*ToDo, error) {
	u := NewToDo(name, id)
	err := u.SetPriority(pri)
	if err != nil {
		return &ToDo{}, err
	}
	return u, nil
}

type ToDoList struct {
	ToDos []ToDo
	//Lists map[string]ToDoList
}
