package utils

import (
	"fmt"
	"github.com/bxcodec/faker/v4"
	"grpc/pb"
)

type SomeStructWithTags struct {
	PhoneNumber string `faker:"phone_number"`
	UserName    string `faker:"username"`
}

func GenerateUser() *pb.CreateUserRequest {
	user := SomeStructWithTags{}
	err := faker.FakeData(&user)
	if err != nil {
		fmt.Println(err)
	}
	return &pb.CreateUserRequest{
		Name:  user.UserName,
		Phone: user.PhoneNumber,
	}
}
