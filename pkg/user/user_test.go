package user

import (
	DataBase "crud/db"
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestCRUD(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	DataBase.DB = sqlxDB

	//CREATE USER
	u1 := User{
		Data: Data{
			FirstName: "test_first_name_01",
			LastName:  "test_last_name_01",
			Interests: "test_interest1_01,test_interest2_01",
		},
	}

	data := fmt.Sprintf("%s %s %s", u1.Data.FirstName, u1.Data.LastName, u1.Data.Interests)

	mock.ExpectExec(`INSERT INTO users`).WithArgs(data).WillReturnResult(sqlmock.NewResult(1, 1))

	err = u1.CreateUser()
	if err != nil {
		log.Println(err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	//READ USER
	mock.ExpectQuery(`SELECT`).WithArgs(u1.ID).WillReturnRows(sqlmock.NewRows([]string{"data"}).AddRow(data))

	_, err = u1.GetUserByID()
	if err != nil {
		log.Println(err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	//UPDATE USER
	u2 := User{
		Data: Data{
			FirstName: "test_first_name_02",
			LastName:  "test_last_name_02",
			Interests: "test_interest_02,test_interest_02",
		},
	}

	data = fmt.Sprintf("%s %s %s", u2.Data.FirstName, u2.Data.LastName, u2.Data.Interests)

	mock.ExpectExec(`UPDATE users SET`).WithArgs(data, u1.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	err = u2.UpdateUserByID()
	if err != nil {
		log.Println(err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	//READ UPDATED USER
	mock.ExpectQuery(`SELECT`).WithArgs(u1.ID).WillReturnRows(sqlmock.NewRows([]string{"data"}).AddRow(data))

	_, err = u2.GetUserByID()
	if err != nil {
		log.Println(err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	//DELETE USER
	mock.ExpectExec(`DELETE FROM users`).WithArgs(u1.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	err = u2.DeleteUserByID()
	if err != nil {
		log.Println(err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	defer DataBase.DB.Close()
}
