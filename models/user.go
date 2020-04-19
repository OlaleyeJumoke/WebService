package models

import (
	"fmt"
)

//User is the json format
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers this gets the user data
func GetUsers() []*User {
	return users
}

//AddUser adds users info
func AddUser(u User) (User, error) {
	//if u.ID == 0 {
	//	return User{}, errors.New("New user must not include Id")
	//}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

//GetUserByID gets user by ID
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with id '%v' not found", id)
}

//UpdateUser updates user
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with id '%v' not found", u.ID)
}

// RemoveUserByID ""
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with id '%v' not found", id)
}
