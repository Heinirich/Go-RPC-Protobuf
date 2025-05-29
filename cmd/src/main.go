package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Heinirich/grpc/protocol"
	"github.com/Heinirich/grpc/server"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type Configuration struct {
	Dn  string `json:"driverName"`
	Dsn string `json:"dsn"`
}

var conf Configuration = Configuration{}

func runGrpcClient() {

	// Listen for incoming connections on port 8085
	fmt.Println("Server is listening on port 8085")

	newClient, err := grpc.Dial("127.0.0.1:8085", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer func(newClient *grpc.ClientConn) {
		err := newClient.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(newClient)

	grpcClient := protocol.NewStudentServiceClient(newClient)

	fmt.Println("\n\nWelcome,Please enter number to select your requests\n")

	fmt.Println("1. Get all by name")
	fmt.Println("2. Get student by id")

	fmt.Println("\nPlease enter your choice")

	var input string
	fmt.Scanln(&input)

	if strings.EqualFold(input, "1") {
		value := ""

		fmt.Print("Enter your name:")

		_, err := fmt.Scanln(&value)
		if err != nil {
			return
		}
		fmt.Println(value)

		students, err := grpcClient.GetStudentsByName(context.Background(), &protocol.SearchByName{Name: value})

		if err != nil {
			log.Fatal(err.Error())
		}

		for {
			// Recv receives the next response message from the server
			student, err := students.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println("Name:", student.Name)
			fmt.Println("Age:", student.Age)
			fmt.Println("Id:", student.Id)
		}
	} else if strings.EqualFold(input, "2") {

		value := ""

		fmt.Print("Enter your id:")

		_, err := fmt.Scanln(&value)
		if err != nil {
			return
		}
		fmt.Printf("Your id is:", value)

		id, err := strconv.Atoi(value)

		if err != nil {
			log.Fatal(err.Error())
		}

		student, err := grpcClient.GetStudentByID(context.Background(), &protocol.SearchByID{Id: int64(id)})

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Name:", student.Name)
		fmt.Println("Age:", student.Age)
		fmt.Println("Id:", student.Id)
	} else {
		fmt.Println("Invalid input")
	}

}

func runGrpcServer() {

	fmt.Println("Starting server .............. ")

	// Listen for incoming connections on port 8085
	listener, err := net.Listen("tcp", "127.0.0.1:8085")

	// Check for errors
	if err != nil {
		panic(err)
	}

	fmt.Println("Server is listening on port 8085")

	var options []grpc.ServerOption

	newServer := grpc.NewServer(options...)

	studentServer, err := server.GrpcServerInitializer(conf.Dn, conf.Dsn)

	if err != nil {
		panic(err)
	}

	// Register the server
	protocol.RegisterStudentServiceServer(newServer, studentServer)

	err = newServer.Serve(listener)

	if err != nil {
		log.Fatal(err.Error())
	}

}

func main() {
	file, err := os.Open("configuration/file.json")

	if err != nil {
		panic(err)
	}

	var configData map[string]string

	// Allow us to decode the JSON file
	err = json.NewDecoder(file).Decode(&configData)

	if err != nil {
		panic(err)
	}

	option := flag.String("admin", "server", "Communication between client and server")

	flag.Parse()

	switch *option {
	case "client":
		runGrpcClient()
	case "server":
		runGrpcServer()
	}

}
