package file

import (
	"bufio"
	"encoding/json"
	"filesys/db/model"
	"fmt"
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

}

func (ur *UserRepository) GetAll() (*[]model.User, error) {

}

func (ur *UserRepository) Update() (*model.User, error) {

}

func (ur *UserRepository) Delete() error {

}*/
