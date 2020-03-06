package todo

import (
	"bufio"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"sort"
)

type User struct {
	Name      string
	Uuid      uuid.UUID
	IdCounter int
	Ids       []int
	Lists     map[string]*ToDoList
}

func NewUser(name string) (*User, error) {
	var id, e = uuid.NewV4()
	if e != nil {
		return &User{}, e
	}
	u := &User{Name: name, Uuid: id, IdCounter: 0, Lists: map[string]*ToDoList{}}
	return u, nil
}

func (u *User) InitDatabase() {
	u.Lists["Inbox"] = &ToDoList{ToDos: []ToDo{}}
	u.Lists["Inbox"].ToDos = append(u.Lists["Inbox"].ToDos, *NewToDo("Welcome to OpenDo!", u.IdCounter))
	u.AddId()
}

func InitUser() (User, error) {
	// init scanner
	sc := bufio.NewScanner(os.Stdin)

	// scan stuff
	fmt.Println("Input Name: ")
	sc.Scan()
	if sc.Err() != nil {
		return User{}, sc.Err()
	}

	// initialise user
	u, err := NewUser(sc.Text())
	if err != nil {
		return User{}, err
	}

	fmt.Println("Uuid " + u.Uuid.String() + " is now named \"" + u.Name + "\"")

	// init database
	u.InitDatabase()

	return *u, nil
}

func (u *User) AddId() {
	// add and sort value
	u.Ids = append(u.Ids, u.IdCounter)
	sort.Ints(u.Ids)
	u.IncrementCounter()
}

func (u *User) IncrementCounter() {
	// increment counter
	u.IdCounter++

	// reset counter if is ludicrously large
	if u.IdCounter > 9999 {
		u.IdCounter = 0
	}

	// make sure new value does not collide
	for _, id := range u.Ids {
		if u.IdCounter == id {
			u.IdCounter++
		}
	}
}

func (u *User) RemoveId(id int) {
	for i, x := range u.Ids {
		if x == id {
			u.Ids = append(u.Ids[:i], u.Ids[i+1:]...)
			return
		}
	}
}

func (u *User) ResetCounter() {
	u.IdCounter = -1
	u.IncrementCounter()
}
