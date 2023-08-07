package Errors

import (
	"fmt"
)

const UserNotFoundText = "user was not found"

type UserNotFound struct {
	ID       *int
	Username *string
}

func (UserNotFound) Error() string {
	return UserNotFoundText
}

func (u UserNotFound) GetID() *int {
	return u.ID
}

func (u UserNotFound) GetUsername() *string {
	return u.Username
}

func (u UserNotFound) WithID() string {
	if u.ID == nil {
		return UserNotFoundText
	}
	return fmt.Sprintf("user with id '%v' was not found", u.ID)
}

func (u UserNotFound) WithUsername() string {
	if u.Username == nil {
		return UserNotFoundText
	}
	return fmt.Sprintf("user with username '%v' was not found", u.Username)
}

func (u UserNotFound) WithIDWithUsername() string {
	if u.ID == nil || u.Username == nil {
		return UserNotFoundText
	}
	return fmt.Sprintf("user with id '%v' and username '%v' was not found", u.ID, u.Username)
}

const UserDisabledText = "user is disabled"

type UserDisabled struct {
	ID       *int
	Username *string
}

func (UserDisabled) Error() string {
	return UserDisabledText
}

func (u UserDisabled) GetID() *int {
	return u.ID
}

func (u UserDisabled) GetUsername() *string {
	return u.Username
}

func (u UserDisabled) WithID() string {
	if u.ID == nil {
		return UserDisabledText
	}
	return fmt.Sprintf("user with id '%v' is disabled", u.ID)
}

func (u UserDisabled) WithUsername() string {
	if u.Username == nil {
		return UserDisabledText
	}
	return fmt.Sprintf("user with username '%v' is disabled", u.Username)
}

func (u UserDisabled) WithIDWithUsername() string {
	if u.ID == nil || u.Username == nil {
		return UserDisabledText
	}
	return fmt.Sprintf("user with id '%v' and username '%v' is disabled", u.ID, u.Username)
}
