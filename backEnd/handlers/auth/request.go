package auth

import "fmt"

func errParamIsRequired(n, t string) error {
	return fmt.Errorf("param: %s (type %s) is required", n, t)
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {

	if r.Name == "" && r.Email == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}


	return nil
}

func (r *LoginUserRequest) Validate() error {
	if r.Email == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}

	return nil
}
