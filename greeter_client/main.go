package main

import (
	"log"
	"os"
	pb "github.com/go-up/go-example/grpc-helloworld/helloworld"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

type Config struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:9876", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	for {
		r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {

			//code := grpc.Code(err)
			log.Println("could not greet: ", err)
			continue
		}
		log.Printf("Greeting: %s", r.Message)
		time.Sleep(time.Second)
	}

	//stream, err := c.SayHello2(context.Background(), &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//for {
	//	reply, err := stream.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		log.Printf("failed to recv: %v", err)
	//	}
	//	log.Printf("Greeting: %s", reply.Message)
	//
	//	time.Sleep(time.Second * 1)
	//}
}
