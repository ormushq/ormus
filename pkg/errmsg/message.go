package errmsg

const (
	ErrJwtEmptyUser       = "for generating a JWT token, email is required"
	ErrWrongCredentials   = "username or password isn't correct"
	ErrSomeThingWentWrong = "some thing went wrong"
	ErrAuthUserNotFound   = "user not found"
	ErrEmailIsNotValid    = "email is not valid"
	ErrAuthUserExisting   = "a user with this email is already registered"
	ErrPasswordIsNotValid = "password is not valid"
	ErrorMsgInvalidInput  = "invalid input"
	ErrBadRequest         = "Bad request"
)
