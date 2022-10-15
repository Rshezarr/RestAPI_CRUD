package user

import "crud/db"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func (u *User) CreateUser() error {
	query := `INSERT INTO users (first_name, last_name) VALUES ($1, $2);`
	_, err := db.DB.Exec(query, u.FirstName, u.LastName)
	return err
}

func (u *User) GetUserByID() (User, error) {
	query := `SELECT * FROM users WHERE id = $1;`
	var usr User
	if err := db.DB.QueryRow(query, u.ID).Scan(&usr.ID, &usr.FirstName, &usr.LastName); err != nil {
		return User{}, err
	}
	return usr, nil
}

func (u *User) UpdateUserByID() error {
	query := `UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3;`
	_, err := db.DB.Exec(query, u.FirstName, u.LastName, u.ID)
	return err
}

func (u *User) DeleteUserByID() error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := db.DB.Exec(query, u.ID)
	return err
}
