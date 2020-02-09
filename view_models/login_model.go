package view_models

import "errors"

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m LoginModel) Valid() error {
	if len(m.Username) < 1 {
		return errors.New("username is required")
	}
	if len(m.Password) < 1 {
		return errors.New("password is required")
	}
	return nil
}
