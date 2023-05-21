package main

import (
	"bufio"
	model "filesys/db/model"
	file "filesys/db/repo/file"
	"fmt"
	"os"
	"time"
)

// Helper functions for finishing the assingment

type customScanner bufio.Scanner

func (scanner *customScanner) scan(msg string) string {
	fmt.Println(msg)
	convScanner := bufio.Scanner(*scanner)
	convScanner.Scan()
	return convScanner.Text()
}

func createUser(ur *file.UserRepository) {
	scanner := customScanner(*bufio.NewScanner(os.Stdin))
	users, _ := ur.GetAll()
	fname := scanner.scan("Enter a first name:")
	lname := scanner.scan("Enter a last name:")
	email := scanner.scan("Enter an email:")
	password := scanner.scan("Enter a password:")

	u := model.User{
		ID:        int32(len(*users)),
		Firstname: fname,
		Lastname:  lname,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().Unix(),
		UpdateAt:  -1,
	}

	ur.Create(&u)
}

func getUserByEmail(ur *file.UserRepository, email string) {
	u, err := ur.GetByEmail(&email)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*u)
}

func updateUser(ur *file.UserRepository, email string) {
	user, _ := ur.GetByEmail(&email)
	ur.Update(user)
}

func main() {
	ur := &file.UserRepository{}
	scanner := customScanner(*bufio.NewScanner(os.Stdin))

	for i := 0; i < 3; i++ {
		fmt.Println("Creating a new user!")
		createUser(ur)
		fmt.Println()
	}

	fmt.Println("The list of all users:")
	users, _ := ur.GetAll()
	fmt.Println(*users)
	fmt.Println()

	fmt.Println("Getting user by email!")
	getUserByEmail(ur, scanner.scan("Enter an email address to get a user!"))
	fmt.Println()

	fmt.Println("Updating user!")
	updateUser(ur, scanner.scan("Enter an email address to update a user!"))
	fmt.Println()

	// Delete cannot be called immediately because the file is still open when delete is called
	/*fmt.Println("Deleting a user!")
	id, _ := strconv.Atoi(scanner.scan("Enter an ID to delete a user!"))
	ur.Delete(int32(id))*/
}
