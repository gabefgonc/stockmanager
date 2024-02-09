package admins

import "errors"

var (
	ErrInternal           = errors.New("internal Error")
	ErrNameTooShort       = errors.New("name too short")
	ErrNameTooLong        = errors.New("name too long")
	ErrEmailInvalid       = errors.New("email invalid")
	ErrPhoneNumberInvalid = errors.New("phone Number invalid")
	ErrPasswordTooShort   = errors.New("password too short")
	ErrPasswordTooLong    = errors.New("password too long")
	ErrWrongCredentials   = errors.New("wrong credentials")
)
