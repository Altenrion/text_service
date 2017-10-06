package main

import (
	"flag"
	"log"
	pb "github.com/altenrion/text_service/api"

	"google.golang.org/grpc"
	"context"
	"os"
	"fmt"
)

func main()  {

	backHost := flag.String("b", "localhost:8080","Address of the service backend")
	flag.Parse()

	if flag.NArg() <1 {
		fmt.Printf("usage:\n\t%s \"text to speak\"\n", os.Args[0])
		os.Exit(1)
	}

	conn, err := grpc.Dial(*backHost, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to %s : %v", *backHost, err )
	}

	defer conn.Close()


	client := pb.NewTextServiceClient(conn)

	message := &pb.Message{Sender: flag.Arg(0) , Text:  flag.Arg(1) }
	res, err := client.TransformMessage(context.Background(), message)
	if err != nil {
		log.Fatalf("could not say %s: %v", message.Text, err)
	}

	log.Printf("Получен ответ: %s : %s", res.Sender, res.Text)

}