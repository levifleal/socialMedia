package schemas

type User struct {
	BaseSchema
	Id           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
}

type UserRespose struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	BaseSchema
}
