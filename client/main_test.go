package main

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	pb "github.com/otabe555/grpc/proto"
	"google.golang.org/grpc"
)

var (
	ctx    context.Context
	conn   *grpc.ClientConn
	cancel context.CancelFunc
	c      pb.ContactRequestClient
)

func TestMain(t *testing.M) {
	var err error
	conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Couldnt connect: %v", err)
	}
	c = pb.NewContactRequestClient(conn)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	os.Exit(t.Run())
}
