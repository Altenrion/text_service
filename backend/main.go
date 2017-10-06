package main

import (
	"github.com/sirupsen/logrus"
	"flag"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "github.com/altenrion/text_service/api"
	"os"
	"log"
)

func main(){

	port := flag.Int("p", 8080, "Some passed port")
	flag.Parse()

	//SomeArg := os.Args[1]






	logrus.Infof("Listening to port %d", *port)


	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("Could not listen port %d: %v", *port, err)

	}


	serv := grpc.NewServer()
	pb.RegisterTextServiceServer(serv, server{})

	err = serv.Serve(lis)
	if err != nil {
		logrus.Fatalf("could not serve: %v", err)
	}

}

type server struct{}

func (server) TransformMessage (ctx context.Context, message *pb.Message) ( *pb.Message, error){

	if message.Sender == "Admin"{
		return nil, fmt.Errorf("you shoul use secured channel for connection")
	}

	replyMessage := message
	replyMessage.Text = fmt.Sprintf("Hi %s! Your message '%s' was registered! ", message.Sender, message.Text)
	replyMessage.Sender = "Admin"
	return replyMessage, nil

}

	//SomeArg := os.Args[1]
	//
	//if SomeArg == ""{
	//	log.Fatal("You forgot the param")
	//}
	//
	////log.SetFormatter(&log.JSONFormatter{})
	//
	//log.SetOutput(os.Stdout)
	//
	//log.WithFields(log.Fields{
	//	"param": SomeArg,
	//	"comment":   "what's that???",
	//}).Info("Some variable was passed...")
