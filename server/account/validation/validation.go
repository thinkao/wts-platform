package validation

type RegistValid struct {
	Username string `validate:"max=20"`
	Password string `validate:"required"`
	Phone    string `validate:"max=11,min=11"`
}

type LoginValid struct {
	UsernamePhone string `valid:"max=20"`
	Password      string `valid:"required"`
}
