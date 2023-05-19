package repo

import "filesys/db/model"

type UserRepositoryI interface {
	Create(user *model.User) (int32, error)
	GetByEmail(email *string) (*model.User, error)
	GetAll() (*[]model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id int32) error
}
