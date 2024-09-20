package errs

import "errors"

var (
	ErrPermissionDenied            = errors.New("ErrPermissionDenied")
	ErrValidationFailed            = errors.New("ErrValidationFailed")
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")
	ErrClassNotFound               = errors.New("ErrClassNotFound")
	ErrIncorrectUsernameOrPassword = errors.New("ErrIncorrectUsernameOrPassword")
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")
	ErrUserIsNotTeacher            = errors.New("ErrUserIsNotTeacher")
)
