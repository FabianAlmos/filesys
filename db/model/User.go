package model

type User struct {
	ID        int32  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
	UpdateAt  int64  `json:"updated_at"`
}
