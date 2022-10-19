package user

import "testing"

// import (
// 	DataBase "crud/db"
// 	"log"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/jmoiron/sqlx"
// )

func TestCRUD(t *testing.T) {}

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	sqlxDB := sqlx.NewDb(db, "sqlmock")
// 	DataBase.DB = sqlxDB

// 	//CREATE USER
// 	u1 := User{FirstName: "test_first_name_01", LastName: "test_last_name_01"}
// 	mock.ExpectExec(`INSERT INTO users`).WithArgs(u1.FirstName, u1.LastName).WillReturnResult(sqlmock.NewResult(1, 1))

// 	err = u1.CreateUser()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	if err = mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}

// 	//READ USER
// 	mock.ExpectQuery(`SELECT`).WithArgs(u1.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name"}).AddRow(u1.ID, u1.FirstName, u1.LastName))

// 	_, err = u1.GetUserByID()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	if err = mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}

// 	//UPDATE USER
// 	u2 := User{FirstName: "test_first_name_02", LastName: "test_last_name_02"}
// 	mock.ExpectExec(`UPDATE users SET`).WithArgs(u2.FirstName, u2.LastName, u2.ID).WillReturnResult(sqlmock.NewResult(1, 1))

// 	err = u2.UpdateUserByID()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	if err = mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}

// 	//READ UPDATED USER
// 	mock.ExpectQuery(`SELECT`).WithArgs(u1.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name"}).AddRow(u2.ID, u2.FirstName, u2.LastName))

// 	_, err = u2.GetUserByID()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	if err = mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}

// 	//DELETE USER
// 	mock.ExpectExec(`DELETE FROM users`).WithArgs(u1.ID).WillReturnResult(sqlmock.NewResult(1, 1))

// 	err = u2.DeleteUserByID()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	if err = mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}

// 	defer DataBase.DB.Close()
// }
