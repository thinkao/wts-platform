package validation

type ProblemAddValid struct {
	Content   string `validation:"required"`
	Option    string `validation:"required"`
	Answer    string `validation:"required"`
	Type      string `validation:"required"`
	Difficult string `validation:"required"`
}

type ProblemUpdateValid struct {
	Id        int `validation:"required"`
	Content   string
	Option    string
	Answer    string
	Type      string
	Difficult string
}

type ProblemGetValid struct {
	Type      string
	Difficult string
	Size      int
}

type FightGetValid struct {
	Type      string
	Difficult string
	Size      int
}

type FightPostValid struct {
	Id		int		`validation:"required"`
	Answer	string

}

type DeleteByIdValid struct {
	Id int `validate:"required"`
}
