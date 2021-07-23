package main

import (
	"context"
	"log"
	"net"

	"github.com/otabe555/grpc/contactlist"
	"github.com/otabe555/grpc/model"
	pb "github.com/otabe555/grpc/proto"

	"google.golang.org/grpc"
)

const (
	port = ":9099"
)

type server struct {
	pb.UnimplementedContactRequestServer
}

func (s *server) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := req.GetUser()
	cm, err := contactlist.NewContactManager()
	if err != nil {
		return nil, err
	}

	contact := &model.Contact{
		Name:      user.GetName(),
		Phone:     user.GetPhone(),
		Gender:    user.GetGender(),
		Email:     user.GetEmail(),
		CreatedAt: user.GetCreatedat(),
	}

	user.Id = int64(contact.ID)
	err = cm.AddContact(contact)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{User: user}, nil
}

func (s *server) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := req.GetUser()
	cm, err := contactlist.NewContactManager()
	if err != nil {
		return nil, err
	}
	contact := &model.Contact{
		Name:      user.GetName(),
		Phone:     user.GetPhone(),
		Gender:    user.GetGender(),
		Email:     user.GetEmail(),
		CreatedAt: user.GetCreatedat(),
	}

	err = cm.UpdateContact(int(user.GetId()), contact)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{User: user}, nil
}

func (s *server) Get(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user := req.GetId()
	cm, err := contactlist.NewContactManager()
	if err != nil {
		return nil, err
	}
	contact, err := cm.GetContact(int(user))
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{User: &pb.User{
		Id:        int64(contact.ID),
		Name:      contact.Name,
		Phone:     contact.Phone,
		Gender:    contact.Gender,
		Email:     contact.Email,
		Createdat: contact.CreatedAt,
	}}, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	user := req.GetId()
	cm, err := contactlist.NewContactManager()
	if err != nil {
		return &pb.DeleteUserResponse{Success: false}, err
	}
	err = cm.DeleteContact(int(user))
	if err != nil {
		return &pb.DeleteUserResponse{Success: false}, err
	}

	return &pb.DeleteUserResponse{Success: true}, nil
}

func (s *server) GetAll(req *pb.GetAllRequest, stream pb.ContactRequest_GetAllServer) error {
	var contacts []model.Contact
	cm, err := contactlist.NewContactManager()
	if err != nil {
		return err
	}
	contacts, err = cm.GetAllContacts()
	if err != nil {
		return err
	}
	for index := range contacts {
		user := &pb.User{
			Id:        int64(contacts[index].ID),
			Name:      contacts[index].Name,
			Phone:     contacts[index].Phone,
			Gender:    contacts[index].Gender,
			Email:     contacts[index].Email,
			Createdat: contacts[index].CreatedAt,
		}

		if err := stream.Send(&pb.GetAllResponse{User: user}); err != nil {
			return err
		}

	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterContactRequestServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
