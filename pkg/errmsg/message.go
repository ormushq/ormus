package errmsg

const (
	ErrJwtEmptyUser       = "for generating a JWT token, email is required"
	ErrWrongCredentials   = "username or password isn't correct"
	ErrSomeThingWentWrong = "some thing went wrong"
	ErrAuthUserNotFound   = "user not found"
	ErrAuthUserExisting   = "a user with this email is already registered"
)
