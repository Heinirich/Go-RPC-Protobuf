package dbtools

import (
	"database/sql"
	"github.com/Heinirich/grpc/model"
	_ "github.com/go-sql-driver/mysql" // 👈 required for MySQL driver registration
	"log"
)

// DBInitializer is a struct that wraps a database connection. It provides methods to initialize the database and to execute queries.
type DBInitializer struct {
	db *sql.DB
}

// Connect creates a new database connection and returns a pointer to the DBInitializer struct.
func Connect(dn, dsn string) (*DBInitializer, error) {
	// Open a database connection
	db, err := sql.Open(dn, dsn)
	// Check for errors
	if err != nil {
		log.Fatal(err.Error())
	}
	// Return a pointer to the DBInitializer struct containing the database connection
	return &DBInitializer{
		db: db,
	}, nil

}

// SelectStudentBasedId returns a student object based on the provided ID.
// The function receiver is a pointer to the DBInitializer struct.
func (initializer *DBInitializer) SelectStudentBasedId(id int64) (model.Student, error) {

	// Create a new student object
	//student := model.Student{}
	var student model.Student

	// Query the database for a student based on the provided ID
	row := initializer.db.QueryRow("SELECT id, name, age FROM students WHERE id = ?", id)

	// Scan the query result into the student object
	err := row.Scan(&student.ID, &student.Name, &student.Age)

	// Check for errors
	if err != nil {
		return student, err
	}

	// Return the student object
	return student, nil
}

func (initializer *DBInitializer) SelectStudentsBasedName(name string) ([]model.Student, error) {

	// Create a slice to hold the student objects
	var students []model.Student

	// Query the database for students based on the provided name
	rows, err := initializer.db.Query("SELECT id, name, age FROM students WHERE name = ?", name)

	// Check for errors
	if err != nil {
		log.Fatal(err.Error())
	}

	// Scan the query result into the student objects
	for rows.Next() {
		// Create a new student object
		var student model.Student
		// Scan the query result into the student object
		err := rows.Scan(&student.ID, &student.Name, &student.Age)
		// Check for errors
		if err != nil {
			panic(err.Error())
		}
		// Append the student object to the slice
		students = append(students, student)
	}
	// Return the slice of student objects
	return students, nil
}
