package main

import (
	"fmt"
	"log"
)

// 1. Single Responsibility Principle (SRP)
type User struct {
	Id   int
	Name string
}

type UserService struct {
}

func (userService *UserService) CreateUser(user User) {
	fmt.Println("User created: ", user)
	log.Println("User created: ", user)

}
func main() {

}
