package main

import (
	pb "GoRPC/helloworld"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:3333", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// 拨号，与服务端建立连接
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	// 连接服务端，打印response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Println(ctx)
	r, err := c.SayHello(ctx, &pb.Request{Name: *name})
	if err != nil {
		log.Fatalf("couldn't greet %v", err)
	}
	log.Printf("Greeting:%s", r.GetMessage())
}
