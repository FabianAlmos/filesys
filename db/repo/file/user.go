package file

import (
	"bufio"
	"encoding/json"
	"filesys/db/model"
	"fmt"
	"io"
	"os"
	"strconv"
)

type UserRepository struct{}

func (ur *UserRepository) Create(user *model.User) (int32, error) {
	file, err := os.Create(fmt.Sprintf("data/users/%s.json", strconv.Itoa(int(user.ID))))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	json, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	writer.WriteString(string(json))
	writer.Flush()
	return user.ID, err
}

func (ur *UserRepository) GetByEmail(email *string) (*model.User, error) {
	users, err := ur.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, u := range *users {
		if u.Email == *email {
			return &u, err
		}
	}
	return &model.User{}, fmt.Errorf("couldn't find user by email : %s", *email)
}

func (ur *UserRepository) GetAll() (*[]model.User, error) {
	users := make([]model.User, 0)
	entries, err := os.ReadDir("data/users")
	if err != nil {
		fmt.Println(err)
	}
	for _, e := range entries {
		user := new(model.User)
		info, err := e.Info()
		if err != nil {
			fmt.Println(err)
		}
		content, err := os.Open(fmt.Sprintf("data/users/%s", info.Name()))
		if err != nil {
			fmt.Println(err)
		}
		bytes, err := io.ReadAll(content)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(bytes, &user)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, *user)
	}
	return &users, err
}

// Update helper functions
func ternaryStr(old, new string) string {
	if new == "\n" {
		return old
	}
	return new
}

func Scan(msg, oldVal string) string {
	scanner := *bufio.NewScanner(os.Stdin)
	fmt.Println(msg, "[", oldVal, "]")
	scanner.Scan()
	return ternaryStr(oldVal, scanner.Text())
}

func (ur *UserRepository) Update(user *model.User) (*model.User, error) {
	fmt.Println("If you don't want to update a filed just press Enter!")

	fname := Scan("Currently updating FirstName::", user.Firstname)
	lname := Scan("Currently updating LastName::", user.Lastname)
	password := Scan("Currently updating Password::", user.Password)

	u := model.User{
		Firstname: fname,
		Lastname:  lname,
		Password:  password,
	}

	return &u, fmt.Errorf("")
}

func (ur *UserRepository) Delete(id int32) error {
	err := os.Remove(fmt.Sprintf("data/users/%d.json", id))
	if err != nil {
		fmt.Println(err)
	}
	return err
}
