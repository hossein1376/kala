package structure

import (
	"kala/pkg/Password"
)

type User struct {
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
}
