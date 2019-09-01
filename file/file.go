package file

import (
	"encoding/json"
	"gitlab.com/jhthenerd/openDo/todo"
	"io/ioutil"
	"os"
)

func CreateFile(user todo.User) error {

	// convert user into json
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// get home directory
	dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Create Directory
	err = os.MkdirAll(dir+"/.openDo", os.ModePerm)
	if err != nil {
		return err
	}

	// write json
	err = ioutil.WriteFile(dir+"/.openDo/data.json", data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile() (todo.User, error) {

	// get home directory
	dir, err := os.UserHomeDir()
	if err != nil {
		return todo.User{}, err
	}

	if _, err := os.Stat(dir + "/.openDo/data.json"); err == nil {
		// file exists

		// read file
		bytes, err := ioutil.ReadFile(dir + "/.openDo/data.json")
		if err != nil {
			return todo.User{}, err
		}

		// convert json into user data
		var data todo.User
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			return todo.User{}, err
		}

		return data, err

	} else if os.IsNotExist(err) {
		// file does not exist

		u, err := todo.InitUser()
		if err != nil {
			return todo.User{}, err
		}

		// create file
		err = CreateFile(u)
		if err != nil {
			return todo.User{}, err
		}
		return u, nil

	} else {
		// whomst??? throw all the errors!
		return todo.User{}, err
	}
}
