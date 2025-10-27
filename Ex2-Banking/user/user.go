package main

import (
	"banking/card"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address string
}

var users []User

func main() {
	AddUser(1, "John Doe", "john@example.com", "1234567890", "123 Main St")
	fmt.Println(users)
	card.GetCards()
}

func GetAllUsers() {
	fmt.Println("Getting all users:")
	for _, user := range users {
		fmt.Printf("User ID: %d, Name: %s\n", user.ID, user.Name)
	}
}

func AddUser(id int, name string, email string, phone string, address string) {
	existUser := FindUserByPhone(phone)
	expiryDate := time.Now().AddDate(5, 0, 0)
	expiryDateStr := expiryDate.Format("01/06") // Format MM/YY


	if existUser == nil {
		users = append(users, User{
			ID:      id,
			Name:    name,
			Email:   email,
			Phone:   phone,
			Address: address,
		})
		card.AddCard(rand.Intn(90000)+10000, strconv.Itoa((rand.Intn(90000) + 10000)), id, expiryDateStr, rand.Intn(900)+100,  )
	} else {
		fmt.Println("User already exists.")
	}

}

func FindUserByPhone(phone string) *User {
	for i := range users {
		if users[i].Phone == phone {
			return &users[i]
		}
	}
	return nil
}

func FindUserByID(id int) *User {
	for i := range users {
		if users[i].ID == id {
			return &users[i]
		}
	}
	return nil
}

func UpdateUser(id int, user User) {
	existUser := FindUserByID(id)
	if existUser != nil {
		*existUser = user
	}
}

func DeleteUser(id int) {
	existUser := FindUserByID(id)
	if existUser != nil {
		users = append(users[:id], users[id+1:]...)
	}
}

