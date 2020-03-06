package main

import (
	"flag"
	"fmt"
	"gitlab.com/jhthenerd/openDo/file"
	"gitlab.com/jhthenerd/openDo/todo"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

var add = flag.Bool("add", false, "add todo")
var remove = flag.Bool("remove", false, "remove todo")

var list = flag.String("list", "Inbox", "list for new todo")

var interactive = flag.Bool("i", false, "interactive mode")
var reset = flag.Bool("r", false, "reset database")
var verbose = flag.Bool("v", false, "verbose mode")
var resetCounter = flag.Bool("rc", false, "reset counter")

var user todo.User
var err error

func init() {
	// parse flags
	flag.Parse()

	// check for database reset
	if *reset {
		user, err = todo.InitUser()
		err = file.CreateFile(user)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// read database
	user, err = file.ReadFile()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if *verbose {
		fmt.Println("Database Loaded...")
		fmt.Println("User Name: " + user.Name)
		fmt.Println("User Uuid: " + user.Uuid.String())
	}

	if *add && *remove {
		log.Fatal("add and remove cannot be used in the same command")
	}
	// add
	if *add {
		addToDo()
	}
	// remove
	if *remove {
		removeToDo()
	}
	// check for counter reset
	if *resetCounter {
		user.ResetCounter()
		fmt.Println("Counter reset!")
		fmt.Printf("New value: %d\n", user.IdCounter)
		err = file.CreateFile(user)
		if err != nil {
			log.Fatal(err)
		}
	}


	// Display all todos
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	_, _ = fmt.Fprintln(w, "List \tID \tName \tPriority \tContexts")
	for l, list := range user.Lists {
		for _, dos := range list.ToDos {
			_, _ = fmt.Fprintf(w, "%s\t%d\t%v\t%d\t%v\n",
				l, dos.Id, dos.Name, dos.Priority(), dos.Contexts)
		}
	}
	_ = w.Flush()
}

func addToDo() {
	//fmt.Println("ADD STUFF")
	if flag.Args()[0] == "" {
		log.Fatal("please add a name for the todo!")
	}
	t := todo.NewToDo(flag.Args()[0], user.IdCounter)
	user.AddId()
	if err != nil {
		log.Fatal(err)
	}
	user.Lists[*list].ToDos = append(user.Lists[*list].ToDos, *t)
	err = file.CreateFile(user)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func removeToDo() {
	if flag.Args()[0] == "" {
		log.Fatal("please enter an id to remove")
	}
	id, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range user.Lists {
		for i, dos := range list.ToDos {
			if dos.Id == id {
				list.ToDos = append(list.ToDos[:i], list.ToDos[i+1:]...)
				user.RemoveId(id)
				err = file.CreateFile(user)
				if err != nil {
					log.Fatal(err)
				}
				return
			}
		}
	}
}