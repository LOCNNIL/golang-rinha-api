package getperson

import "github.com/LOCNNIL/golang-rinha-api/app/models"

type GetPersonResponse struct {
	Id        string   `json:"id"`
	Nickname  string   `json:"apelido"`
	Name      string   `json:"nome"`
	Birthdate string   `json:"nascimento"`
	Stack     []string `json:"stack"`
}

func NewPersonResponse(p models.People) *GetPersonResponse {
	return &GetPersonResponse{
		Id:        p.Id,
		Nickname:  p.Nickname,
		Name:      p.Name,
		Birthdate: p.Birthdate,
		Stack:     p.Stack,
	}
}
