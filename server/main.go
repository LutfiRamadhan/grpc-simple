package main

import (
	"context"
	"encoding/json"
	pb "grpc-simple/student"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type DataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*pb.Student
}

func (d *DataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	log.Println("Incoming Request. Email:", student.Email)
	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}
	return nil, nil
}

func (d *DataStudentServer) loadData() {
	data, err := ioutil.ReadFile("data/student.json")
	if err != nil {
		log.Fatalln("Error load data", err.Error())
	}

	if err = json.Unmarshal(data, &d.students); err != nil {
		log.Fatalln("Error unmarshal data", err.Error())
	}
}

func newServer() *DataStudentServer {
	s := DataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("Error listen", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())
	log.Println("Server running on port :1200")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("Error serve", err.Error())
	}
}
