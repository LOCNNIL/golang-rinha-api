package getpeople

type FindPeopleResponse struct {
	Id        string   `json:"id"`
	Nickname  string   `json:"apelido"`
	Name      string   `json:"nome"`
	Birthdate string   `json:"nascimento"`
	Stack     []string `json:"stack"`
}
