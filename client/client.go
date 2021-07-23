package main

import (
	"context"
	"io"
	"log"

	pb "github.com/otabe555/grpc/proto"
)

const (
	address = "localhost:9099"
)

func Create(c pb.ContactRequestClient, user *pb.User, ctx context.Context) error {

	create_response, err := c.Create(ctx, &pb.CreateUserRequest{User: user})
	if err != nil {
		return err
	}
	log.Printf("Created account: %s", create_response.GetUser())
	return nil
}

func Update(c pb.ContactRequestClient, user *pb.User, ctx context.Context) error {
	update_response, err := c.Update(ctx, &pb.UpdateUserRequest{User: user})
	if err != nil {
		log.Fatalf("could not udpate the contact: %v", err)
	}
	log.Printf("Updated account: %s", update_response.GetUser())
	return nil
}

func Delete(c pb.ContactRequestClient, id int64, ctx context.Context) error {
	delete_response, err := c.Delete(ctx, &pb.DeleteUserRequest{Id: id})
	if err != nil {
		return err
	}
	if delete_response.Success {
		log.Printf("contact with id = %d successfully deleted:", id)
	}
	return nil
}

func Get(c pb.ContactRequestClient, id int64, ctx context.Context) error {
	get_response, err := c.Get(ctx, &pb.GetUserRequest{Id: id})
	if err != nil {
		return err
	}
	log.Printf("Got account: %s", get_response.GetUser())
	return nil
}

func GetAll(c pb.ContactRequestClient, ctx context.Context) error {
	get_responses, err := c.GetAll(ctx, &pb.GetAllRequest{})
	if err != nil {
		return err
	}
	var rpcStatus error
	for {
		response, err := get_responses.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		log.Printf("Got acc: %s \n", response.GetUser())
	}
	if rpcStatus != io.EOF {
		return err
	}
	return nil
}
