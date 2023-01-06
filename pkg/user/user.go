package user

import (
	"crud/db"
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int
	Data `json:"data" db:"data"`
}

type Data struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Interests string `json:"interests"`
}

func (u *User) CreateUser() error {
	data, err := json.Marshal(u.Data)
	if err != nil {
		return fmt.Errorf("pkg: create user - %w", err)
	}

	_, err = db.DB.Exec(`INSERT INTO users (data) VALUES ($1);`, data)
	if err != nil {
		return fmt.Errorf("pkg: create user - %w", err)
	}
	return nil
}

func (u *User) GetUserByID() (User, error) {
	var user User
	var data string
	fmt.Println(u.ID)
	if err := db.DB.Get(&data, `SELECT data FROM users WHERE id = $1;`, u.ID); err != nil {
		return User{}, fmt.Errorf("pkg: get user get - %w", err)
	}

	if err := json.Unmarshal([]byte(data), &user.Data); err != nil {
		return User{}, fmt.Errorf("pkg: get user - %w", err)
	}

	user.ID = u.ID

	return user, nil
}

func (u *User) UpdateUserByID() error {
	data, err := json.Marshal(u.Data)
	if err != nil {
		return fmt.Errorf("pkg: update user - %w", err)
	}

	_, err = db.DB.Exec(`UPDATE users SET data = $1 WHERE id = $2;`, data, u.ID)
	if err != nil {
		return fmt.Errorf("pkg: update user - %w", err)
	}

	return nil
}

func (u *User) DeleteUserByID() error {
	_, err := db.DB.Exec(`DELETE FROM users WHERE id = $1;`, u.ID)
	if err != nil {
		return fmt.Errorf("pkg: delete user - %w", err)
	}
	return nil
}
