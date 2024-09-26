package errs

import "errors"

var (
	ErrPermissionDenied            = errors.New("ErrPermissionDenied")
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")
	ErrClassNotFound               = errors.New("ErrClassNotFound")
	ErrIncorrectUsernameOrPassword = errors.New("ErrIncorrectUsernameOrPassword")
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")
	ErrUserIsNotTeacher            = errors.New("ErrUserIsNotTeacher")
	ErrUserDeactivatedOrDeleted    = errors.New("ErrUserDeactivatedOrDeleted")
	ErrFailedSetTeacherToClass     = errors.New("ErrFailedSetTeacherToClass")
	ErrDateIsPast                  = errors.New("ErrDateIsPast")
	ErrIncorrectOldPassword        = errors.New("ErrIncorrectOldPassword")
	ErrFailedValidation            = errors.New("ErrFailedValidation")
)
