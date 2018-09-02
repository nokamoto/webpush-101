package main

import (
	"flag"
	"fmt"
	pb "github.com/nokamoto/webpush-101/webpush-go/protobuf"
	"github.com/nokamoto/webpush-101/webpush-go/webpush-lib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

var (
	port = flag.Int("p", 8000, "gRPC server port")
	priv = flag.String("priv", "", "application server private key base64 encoded")
	pub  = flag.String("pub", "", "application server public key base64 encoded")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("listen %v port", *port)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	pair, err := webpushlib.NewApplicationServerKeyPairFromBase64StdEncodingKeyPair(*priv, *pub)
	if err != nil {
		log.Fatalf("application server key error %s", err)
	}

	client := &webpushlib.PushServiceClient{KeyPair: pair, Client: &http.Client{}}

	pb.RegisterWebpushServiceServer(s, &server{client: client})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
