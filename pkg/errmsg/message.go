package errmsg

const (
	ErrJwtEmptyUser                   = "for generating a JWT token, email is required"
	ErrWrongCredentials               = "username or password isn't correct"
	ErrSomeThingWentWrong             = "some thing went wrong"
	ErrAuthUserNotFound               = "user not found"
	ErrEmailIsNotValid                = "email is not valid"
	ErrAuthUserExisting               = "a user with this email is already registered"
	ErrPasswordIsNotValid             = "password is not valid"
	ErrorMsgInvalidInput              = "invalid input"
	ErrBadRequest                     = "Bad request"
	ErrUserNotFound                   = "user not found"
	ErrChannelNotFound                = "channel not found: %v"
	ErrFailedToOpenChannel            = "failed to open rabbitmq channel"
	ErrFailedToCloseChannel           = "failed to close rabbitmq channel"
	ErrFailedToUpdateWriteKeyMetaData = "failed to update write key metadata"
	ErrFailedToGetWriteKeyMetaData    = "get write key metadata"
	ErrWriteKeyNotFound               = "write key not found"
)
