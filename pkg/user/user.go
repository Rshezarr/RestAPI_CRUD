package user

import (
	"crud/db"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func (u *User) CreateUser() error {
	_, err := db.DB.Exec(`INSERT INTO users (first_name, last_name) VALUES ($1, $2);`, u.FirstName, u.LastName)
	if err != nil {
		return fmt.Errorf("create user - %v", err)
	}

	return nil
}

func (u *User) GetUserByID() (User, error) {
	var usr User
	if err := db.DB.QueryRow(`SELECT * FROM users WHERE id = $1;`, u.ID).Scan(&usr.ID, &usr.FirstName, &usr.LastName); err != nil {
		return User{}, fmt.Errorf("get user - %v", err)
	}

	return usr, nil
}

func (u *User) UpdateUserByID() error {
	_, err := db.DB.Exec(`UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3;`, u.FirstName, u.LastName, u.ID)
	if err != nil {
		return fmt.Errorf("update user - %v", err)
	}

	return nil
}

func (u *User) DeleteUserByID() error {
	_, err := db.DB.Exec(`DELETE FROM users WHERE id = $1;`, u.ID)
	if err != nil {
		return fmt.Errorf("delete user - %v", err)
	}

	return nil
}
