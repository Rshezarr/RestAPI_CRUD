package user

import (
	DataBase "crud/db"
	"encoding/json"
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

	data, err := json.Marshal(u1.Data)
	if err != nil {
		log.Fatal(err)
	}

	mock.ExpectExec(`INSERT INTO users`).WithArgs([]byte(data)).WillReturnResult(sqlmock.NewResult(1, 1))

	if err = u1.CreateUser(); err != nil {
		log.Println(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	//READ USER
	mock.ExpectQuery(`SELECT`).WithArgs(u1.ID).WillReturnRows(sqlmock.NewRows([]string{"data"}).AddRow(data))

	if _, err = u1.GetUserByID(); err != nil {
		log.Println(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	//UPDATE USER
	// u2 := User{
	// 	Data: Data{
	// 		FirstName: "test_first_name_02",
	// 		LastName:  "test_last_name_02",
	// 		Interests: "test_interest_02,test_interest_02",
	// 	},
	// }

	// data, err = json.Marshal(u2.Data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// mock.ExpectExec(`UPDATE users SET`).WithArgs(data, u1.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	// if err := u2.UpdateUserByID(); err != nil {
	// 	log.Println(err)
	// }

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }

	//READ UPDATED USER
	// mock.ExpectQuery(`SELECT`).WithArgs(u1.ID).WillReturnRows(sqlmock.NewRows([]string{"data"}).AddRow(data))

	// if _, err = u2.GetUserByID(); err != nil {
	// 	log.Println(err)
	// }

	// if err = mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }

	// //DELETE USER
	// mock.ExpectExec(`DELETE FROM users`).WithArgs(u1.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	// if err := u2.DeleteUserByID(); err != nil {
	// 	log.Println(err)
	// }

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }

	defer DataBase.DB.Close()
}
