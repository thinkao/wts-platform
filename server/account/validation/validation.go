package validation

type RegistValid struct {
	Email    string
	Password string `validate:"required"`
	Phone    string `validate:"max=11,min=11"`
}

type LoginValid struct {
	Phone    string `validate:"max=11,min=11"`
	Password string `valid:"required"`
}
