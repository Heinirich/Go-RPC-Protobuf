package model

// Student represents a student with their ID, name and age.
//
// The fields are:
//
// - ID: a unique identifier for the student, represented as an int64.
// - Name: the name of the student, represented as a string.
// - Age: the age of the student, represented as an int32.
type Student struct {
	ID   int64
	Name string
	Age  int32
}
