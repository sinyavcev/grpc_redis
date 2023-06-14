package utils

import (
	"fmt"
	"github.com/bxcodec/faker/v4"
	"grpc/internal/models"
)

type SomeStructWithTags struct {
	PhoneNumber string `faker:"phone_number"`
	UserName    string `faker:"username"`
}

func GenerateUser() *models.User {
	user := SomeStructWithTags{}
	err := faker.FakeData(&user)
	if err != nil {
		fmt.Println(err)
	}
	return &models.User{
		Name:  user.UserName,
		Phone: user.PhoneNumber,
	}
}
