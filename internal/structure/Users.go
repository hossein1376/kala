package structure

import (
	"github.com/hossein1376/kala/pkg/Password"
)

type User struct {
	ID        int               `json:"id"`
	FirstName string            `json:"firstname"`
	LastName  string            `json:"lastname"`
	Username  string            `json:"username"`
	Password  Password.Password `json:"-"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	Role      string            `json:"role"`
	Profile   []Image           `json:"profile"`
	Addresses []Address         `json:"addresses"`
	Orders    []Order           `json:"orders"`
}

type UserRequest struct {
	UserName  string  `json:"username"`
	Password  string  `json:"password"`
	FirstName *string `json:"firstname"`
	LastName  *string `json:"lastname"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
}

type UserUpdateRequest struct {
	UserName  *string `json:"username"`
	Password  *string `json:"password"`
	FirstName *string `json:"firstname"`
	LastName  *string `json:"lastname"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type Address struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	Address     string  `json:"address"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Phone       string  `json:"phone"`
	ZipCode     string  `json:"zip_code"`
	Coordinates string  `json:"coordinates"`
	Seller      bool    `json:"seller"`
}
