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

/*func (ur *UserRepository) GetByEmail(email *string) (*model.User, error) {

}*/

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

/*func (ur *UserRepository) Update() (*model.User, error) {

}

func (ur *UserRepository) Delete() error {

}*/
