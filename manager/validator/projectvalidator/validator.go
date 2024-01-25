package projectvalidator

type UserExistenceChecker interface {
	IsUserIDValid(email string) (bool, error)
}

type Validator struct {
	userExistenceChecker UserExistenceChecker
}

func New(userExistenceChecker UserExistenceChecker) Validator {
	return Validator{
		userExistenceChecker: userExistenceChecker,
	}
}
