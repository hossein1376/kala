package Errors

const UserNotFoundText = "user was not found"

type UserNotFound struct {
	ID       *int
	Username *string
}

func (UserNotFound) Error() string {
	return UserNotFoundText
}

const UserDisabledText = "user is disabled"

type UserDisabled struct {
	ID       *int
	Username *string
}

func (UserDisabled) Error() string {
	return UserDisabledText
}
