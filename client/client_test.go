package main

import (
	"testing"

	pb "github.com/otabe555/grpc/proto"
)

var id int64

func TestCreate(t *testing.T) {
	user := &pb.User{
		Name:   "Otabek",
		Phone:  "998999048088",
		Gender: "Male",
		Email:  "mad@gmail.com",
	}
	if err := Create(c, user, ctx); err != nil {
		t.Errorf("Failed to create user: %v", err)
	}
	id = user.Id
}

func TestUpdate(t *testing.T) {
	if err := Update(c, &pb.User{
		Id:     3,
		Name:   "Stan Lee",
		Phone:  "998999048088",
		Gender: "Male",
		Email:  "stan@gmail.com",
	}, ctx); err != nil {
		t.Errorf("Failed to Update user: %v", err)
	}
}

func TestGet(t *testing.T) {
	if err := Get(c, 3, ctx); err != nil {
		t.Errorf("Failed to Update user: %v", err)
	}
}

func TestDelete(t *testing.T) {
	if err := Delete(c, id, ctx); err != nil {
		t.Errorf("failed to delete contact: %v", err)
	}
}

func TestGetAll(t *testing.T) {
	if err := GetAll(c, ctx); err != nil {
		t.Errorf("failed to get all contacts: %v", err)
	}
	conn.Close()
}
