package server

import (
	"context"
	"github.com/Heinirich/grpc/dbtools"
	"github.com/Heinirich/grpc/model"
	"github.com/Heinirich/grpc/protocol"
)

// GrpcController is a struct that implements the gRPC service interface
type GrpcController struct {
	connection *dbtools.DBInitializer
}

// GrpcServerInitializer initializes a new gRPC server controller.
//
// It takes the database driver name (dn) and data source name (dsn) as parameters,
// establishes a connection to the database using these credentials and returns
// a pointer to a GrpcController struct. If there is an error during the connection,
// it returns an error.
func GrpcServerInitializer(driverName string, dataSourceName string) (*GrpcController, error) {

	db, err := dbtools.Connect(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	return &GrpcController{
		connection: db,
	}, err
}

// GetStudentByID implements the GetStudentByID RPC method.
//
// It takes a SearchByID message as input, queries the database and returns
// a Student message.
//
// The function returns an error if there is an issue with the database query.
func (c *GrpcController) GetStudentByID(ctx context.Context, in *protocol.SearchByID) (*protocol.Student, error) {

	// Get a student from the database
	student, err := c.connection.SelectStudentBasedId(in.GetId())
	// Check for errors
	if err != nil {
		return nil, err
	}
	// Return the student
	return convertToGrpcModel(student), nil

}

// GetStudentsByName gets multiple students by their name.
//
// It takes a SearchByName message as input, queries the database and returns
// a stream of Student messages.
//
// The function returns an error if there is an issue with the database query.
func (c *GrpcController) GetStudentsByName(in *protocol.SearchByName, grpcStudents protocol.StudentService_GetStudentsByNameServer) error {
	// Get students from the database
	students, err := c.connection.SelectStudentsBasedName(in.GetName())
	// Check for errors
	if err != nil {
		return err
	}

	// Range over the students
	for _, student := range students {
		grpcStudent := convertToGrpcModel(student)
		if err := grpcStudents.Send(grpcStudent); err != nil {
			return err
		}
	}
	return nil
}

func convertToGrpcModel(student model.Student) *protocol.Student {
	return &protocol.Student{
		Id:   student.ID,
		Name: student.Name,
		Age:  student.Age,
	}
}
