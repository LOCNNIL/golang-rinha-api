package putperson

type PostPersonRequest struct {
	Id        string   `param:"id"`
	Nickname  string   `json:"apelido" validate:"required"`
	Name      string   `json:"nome" validate:"required"`
	Birthdate string   `json:"nascimento" validate:"required"`
	Stack     []string `json:"stack"`
}
