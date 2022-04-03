package main

import (
    "context"
    "fmt"
    "log"

		"user/pb"
    "google.golang.org/grpc"
)
func main() {
    conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatal("client connection error:", err)
    }
    defer conn.Close()
    client := pb.NewRockPaperScissorsServiceClient(conn)
    message := &pb.Empty{}

		user, err := client.GetUser(context.TODO(), message)
		if err != nil {
			fmt.Printf("ERROR: %#v \n", err)
      return
		}

    fmt.Printf("User is %v \n", user.Name)
  }