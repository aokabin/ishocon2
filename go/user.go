package main

import (
	"fmt"
)

// User Model
type User struct {
	ID       int
	Name     string
	Address  string
	MyNumber string
	Votes    int
}

func getUser(name string, address string, myNumber string) (user User, err error) {
	row := db.QueryRow("SELECT * FROM users WHERE mynumber = ?",
		name, address, myNumber)
	err = row.Scan(&user.ID, &user.Name, &user.Address, &user.MyNumber, &user.Votes)
	if err != nil {
		return
	}
	if user.Name != name || user.Address != address {
		err = fmt.Errorf("Error: %s", "Can not find user")
	}
	return
}
