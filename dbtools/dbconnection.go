package dbtools

import (
	"database/sql"
	"github.com/Heinirich/grpc/model"
	"log"
)

type DBInitializer struct {
	db *sql.DB
}

func Connect(dn, dsn string) (*DBInitializer, error) {

	db, err := sql.Open(dn, dsn)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &DBInitializer{
		db: db,
	}, nil

}

func (initializer *DBInitializer) SelectStudentBasedId(id int64) (model.Student, error) {

	student := model.Student{}

	row := initializer.db.QueryRow("SELECT * FROM students WHERE id = ?", id)

	err := row.Scan(&student.ID, &student.Name, &student.Age)

	if err != nil {
		log.Println("No student found")
		return student, err
	}

	return student, nil
}

func (initializer *DBInitializer) SelectStudentsBasedName(name string) ([]model.Student, error) {

	var students []model.Student

	rows, err := initializer.db.Query("SELECT * FROM students WHERE name = ?", name)

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age)
		if err != nil {
			panic(err.Error())
		}
		students = append(students, student)
	}
	return students, nil
}
