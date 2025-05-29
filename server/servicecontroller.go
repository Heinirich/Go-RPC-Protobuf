package server

import (
	"context"
	"github.com/Heinirich/grpc/dbtools"
	"github.com/Heinirich/grpc/model"
	"github.com/Heinirich/grpc/protocol"
	"google.golang.org/grpc"
)

type GrpcController struct {
	connection *dbtools.DBInitializer
}

func GrpcServerInitializer(dn, dsn string) (*GrpcController, error) {
	db, err := dbtools.Connect(dn, dsn)
	if err != nil {
		return nil, err
	}
	return &GrpcController{
		connection: db,
	}, nil
}

func (c *GrpcController) GetStudentByID(ctx context.Context, in *protocol.SearchByID, opts ...grpc.CallOption) (*protocol.Student, error) {

	student, err := c.connection.SelectStudentBasedId(in.GetId())
	if err != nil {
		return nil, err
	}
	return convertToGrpcModel(student), nil

}

func (c *GrpcController) GetStudentsByName(ctx context.Context, in *protocol.SearchByName, grpcStudents protocol.StudentService_GetStudentsByNameServer) error {

	students, err := c.connection.SelectStudentsBasedName(in.GetName())
	if err != nil {
		return err
	}

	for _, student := range students {
		grpcStudent := convertToGrpcModel(student)
		if err := grpcStudents.Send(grpcStudent); err != nil {
			return nil
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
