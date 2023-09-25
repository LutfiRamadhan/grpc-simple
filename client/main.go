package main

import (
	"context"
	"log"
	"time"

	pb "grpc-simple/student"

	"google.golang.org/grpc"
)

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	student, err := client.FindStudentByEmail(ctx, &pb.Student{Email: email})
	if err != nil {
		log.Fatalln("Error find student", err.Error())
	}
	log.Println("Student:", student)
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1200", opts...)
	if err != nil {
		log.Fatalln("Error dial", err.Error())
	}
	defer conn.Close()

	client := pb.NewDataStudentClient(conn)
	getDataStudentByEmail(client, "lutfiramadan@gmail.com")
	getDataStudentByEmail(client, "lutfi.ramadhan@gmail.com")
}
